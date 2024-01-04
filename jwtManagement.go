package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

/*
JWT를 사용하기 위해선, 서버에서 토큰을 발급해서 클라이언트에 전달해야한다.
세션같은 경우에는 Echo 프레임워크에서 session.Save(c.Request(), c.Response()) 과 같이
메소드가 제공되서 클라이언트에 쿠키로 저장을 바로 할 수 있는데, JWT는 그런 기능이 제공되지 않는다.
따라서, HTTP 응답에 token을 같이 담아서 보내준다.

보안을 위해서는 그렇게 하고, 클라이언트가 소중히 보관한 토큰값을 일일이 API요청시마다 헤더에 넣어서 주던가 해야하지만,
편의성을 위해서 서버에서 쿠키에 담고, 쿠키의 토큰을 알아서 검사하도록 하자...

*/
// JWT 토큰 키에 대해서 myJWTKey처럼 SigningKey를 고정해서 사용하는 방법과
// 외부에서 Key를 동적으로 가져오는 방법이 있다.
// 두 가지 방식의 장단점이 있고, 난 간편성을 위해 고정해서 사용하도록 한다.
// https://echo.labstack.com/docs/cookbook/jwt

const (
	myJWTKey     = "Suuuper-Secret-BSMG-Key"
	myCookieName = "bsmgToken"
)

type MemberClaims struct {
	Mem_ID   string `json:"mem_id" gorm:"type:varchar(20);unique_key"`
	Mem_Name string `json:"mem_name" gorm:"type:nvarchar(50)"`
	Mem_Rank int32  `json:"mem_rank" gorm:"type:int"`
	Mem_Part int32  `json:"mem_part" gorm:"type:int"`
	jwt.RegisteredClaims
}

func makeJwtToken(claims *MemberClaims) (string, error) {

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(myJWTKey))
	if err != nil {
		return "", err
	}

	return t, err
}

// JWT 토큰 쿠키 생성
func createCookie(c echo.Context, claims *MemberClaims, token string) {
	// 쿠키 (클라이언트가 따로 저장하는 수고를 덜기 위해서...)
	cookie := new(http.Cookie)
	cookie.Name = myCookieName
	cookie.Value = token
	cookie.Expires = claims.ExpiresAt.Local()
	c.SetCookie(cookie)
}

// JWT 토큰 쿠키 삭제
func deleteCookie(c echo.Context) {
	// 만료기간을 이전날짜로 하여 쿠키 삭제
	expire := time.Now().AddDate(0, 0, -1)

	cookie := new(http.Cookie)
	cookie.Name = myCookieName
	cookie.Value = ""
	cookie.Expires = expire
	c.SetCookie(cookie)
}

// 쿠키에서 JWT 토큰 추출
func extractJwtFromCookie(c echo.Context) (string, error) {
	cookie, err := c.Cookie("bsmgToken")
	if err != nil {
		return "", err
	}

	jwtTokenString := cookie.Value
	return jwtTokenString, nil
}

// JWT 토큰에서 클레임 추출
func extractClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	// 토큰 파싱
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// 토큰의 서명방식 확인  default :  HMAC SHA256 == HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// 이곳에서 토큰을 검증하는데 사용되는 key를 반환
		return []byte(myJWTKey), nil
	})

	if err != nil {
		return nil, err
	}

	// 클레임 추출
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return claims, nil

}
