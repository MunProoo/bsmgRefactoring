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
// JWT 토큰 키에 대해서 SigningKey를 고정해서 사용하는 방법과
// 외부에서 Key를 동적으로 가져오는 방법이 있다.
// 두 가지 방식의 장단점이 있고, 난 간편성을 위해 고정해서 사용하도록 한다.
// https://echo.labstack.com/docs/cookbook/jwt

const (
	AccessTokenKey    = "AC_Suuuper-Secret-Key"
	AccessCookieName  = "AC_bsmgCookie"
	RefreshTokenKey   = "RS_Suuuper-Secret-Key"
	RefreshCookieName = "RS_bsmgCookie"
)

type MemberClaims struct {
	Mem_ID   string `json:"mem_id" gorm:"type:varchar(20);unique_key"`
	Mem_Name string `json:"mem_name" gorm:"type:nvarchar(50)"`
	Mem_Rank int32  `json:"mem_rank" gorm:"type:int"`
	Mem_Part int32  `json:"mem_part" gorm:"type:int"`
	jwt.RegisteredClaims
}

func makeJwtToken(c echo.Context, claims *MemberClaims) error {

	// Access Token 생성 (유효기간 20분)
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 20))
	err := createToken(c, claims, AccessTokenKey, AccessCookieName)
	if err != nil {
		return err
	}

	// Refresh Token 생성 (48시간)
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 48))
	err = createToken(c, claims, RefreshTokenKey, RefreshCookieName)
	if err != nil {
		return err
	}

	return nil
}

// 단일 토큰 생성
func createToken(c echo.Context, claims *MemberClaims, tokenKey, cookieName string) error {
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(tokenKey))
	if err != nil {
		return err
	}

	// 쿠키로 저장 (클라이언트단에서 따로 저장하고 요청 시 헤더에 넣어주는 것보다 편하니까)
	createCookie(c, claims, tokenString, cookieName)
	return nil
}

// JWT 토큰으로 쿠키 생성
func createCookie(c echo.Context, claims *MemberClaims, tokenString, cookieName string) {

	cookie := new(http.Cookie)
	cookie.Name = cookieName
	cookie.Value = tokenString

	loc, _ := time.LoadLocation("Asia/Seoul")
	cookie.Expires = claims.ExpiresAt.Time.In(loc)
	cookie.HttpOnly = true // XSS 대비

	// 설정안하면 /bsmg/login으로 path가 잡혀버려서 도메인은 같아도 공유가 안되버림
	cookie.Path = "/bsmg"
	c.SetCookie(cookie)
}

// JWT 토큰 쿠키 삭제
func deleteCookie(c echo.Context, cookieName string) {
	// 만료기간을 이전날짜로 하여 쿠키 삭제
	expire := time.Now().AddDate(0, 0, -1)

	cookie := new(http.Cookie)
	cookie.Name = cookieName
	cookie.Value = ""
	cookie.Expires = expire
	cookie.HttpOnly = true // XSS 대비

	// 설정안하면 /bsmg/login으로 path가 잡혀버려서 도메인은 같아도 공유가 안되버림
	cookie.Path = "/bsmg"
	c.SetCookie(cookie)
}

// 쿠키 -> JWT 토큰 추출
func extractJwtFromCookie(c echo.Context, cookieName string) (string, error) {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return "", err
	}

	jwtTokenString := cookie.Value
	return jwtTokenString, nil
}

// JWT 토큰 -> 클레임 추출
func extractClaimsFromToken(tokenString, tokenKey string) (jwt.MapClaims, error) {
	// 토큰 파싱
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// 토큰의 서명방식 확인  default :  HMAC SHA256 == HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// 이곳에서 토큰을 검증하는데 사용되는 key를 반환
		return []byte(tokenKey), nil
	})

	if err != nil {
		return nil, err
	}

	// 클레임 추출
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// 토큰 만료시간 체크
// true : 살아있음 , false : 만료됨
// 내 프로젝트는 쿠키에 토큰을 저장하고, 쿠키는 만료시간이 다되면 자동삭제니까 내 프로젝트에선 필요없음
func isExpired(claims jwt.MapClaims) (expired bool) {
	exp, exist := claims["exp"].(float64)
	if !exist {
		return
	}

	// JWT에서는 초 단위의 타임스탬프로 표시되므로 time.Unix 함수를 사용하여 시간을 생성
	expirationTime := time.Unix(int64(exp), 0)
	// 현재 시간과 만료 시간 비교
	if time.Now().After(expirationTime) {
		return
	}
	expired = true
	return
}

// Token을 쿠키로부터 꺼내서 검증 후 반환
// AccessToken 만료시 재발급
func checkToken(c echo.Context) (jwt.MapClaims, error) {
	// accessToken 검증
	accessClaims, err := ExtractClaims(c, "AC")
	if err == nil || err.Error() == "http: named cookie not present" {
		// accessToken 쿠키가 만료되서 사라짐
		refreshClaims, err := ExtractClaims(c, "RS")
		if err != nil {
			return nil, err
		}

		// claims 재생성
		memberClaims, err := ClaimsMappingMember(refreshClaims)
		if err != nil {
			return nil, err
		}

		// 토큰 재발급
		makeJwtToken(c, &memberClaims)
		return accessClaims, nil
		/*
			if isExpired(refreshClaims) {
				return accessClaims, nil
			}
		*/
	}
	return nil, errors.New("token is expired")

}

func ExtractClaims(c echo.Context, category string) (jwt.MapClaims, error) {
	cookieName := category + "_bsmgCookie"
	tokenKey := category + "_Suuuper-Secret-Key"

	tokenString, err := extractJwtFromCookie(c, cookieName)
	if err != nil {
		return nil, err
	}

	claims, err := extractClaimsFromToken(tokenString, tokenKey)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func ClaimsMappingMember(claims jwt.MapClaims) (MemberClaims, error) {
	var exist bool
	var memberClaim MemberClaims

	memberClaim.Mem_ID, exist = claims["mem_id"].(string)
	if !exist {
		return MemberClaims{}, errors.New("token doesn't have a member")
	}
	memberClaim.Mem_Name, exist = claims["mem_name"].(string)
	if !exist {
		return MemberClaims{}, errors.New("token doesn't have a member")
	}

	// 무조건 숫자는 float으로 옴
	memPartFloat, exist := claims["mem_part"].(float64)
	if !exist {
		return MemberClaims{}, errors.New("token doesn't have a member")
	}
	memberClaim.Mem_Part = int32(memPartFloat)

	// 무조건 숫자는 float으로 옴
	memRankFloat, exist := claims["mem_rank"].(float64)
	if !exist {
		return MemberClaims{}, errors.New("token doesn't have a member")
	}
	memberClaim.Mem_Rank = int32(memRankFloat)

	return memberClaim, nil

}
