package main

// SPA (클라이언트 측 페이지 전환) 방식이므로 세션체크는 클라이언트에서 한다.

import (
	"BsmgRefactoring/define"
	"encoding/gob"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-contrib/session"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var server ServerProcessor // 핸들러에서도 접근 가능하도록 전역으로 할당

func init() {
	// Gob에 BsmgMemberInfo 타입 등록
	gob.Register(define.BsmgMemberInfo{})
}

func main() {
	// TODO : server를 goroutine으로 돌리기
	// DB 채널통신 테스트중
	server.reqCh = make(chan interface{}, 6000)
	server.resCh = make(chan interface{}, 6000)

	go server.StartServer()

	e := echo.New()

	// middleware ------------------------------------
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("views/webRoot")) // eXBuilder6 의존성 파일 추가
	// e.Use(middleware.Static("views/bsmgApp/webRoot")) // 빠른 디버깅용. 배포위치를 변경하여 front 수정 시 바로 반영되도록

	e.Use(session.Middleware(store)) // 세션 미들웨어 추가
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 세션 초기화
			initSessionMiddleware(c)
			// 스레드 안정성을 고려하여 서버 변수를 컨텍스트에 할당
			c.Set("Server", &server)
			return next(c)
		}
	})

	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})

	// login API만 JWT 인증 제외하기 위해 JWT 미들웨어보다 먼저 선언
	e.POST("/bsmg/login/login", postLoginRequest)
	e.GET("/bsmg/login/chkLogin", getChkLoginRequest)

	// URL 그룹화 + JWT 미들웨어 적용
	bsmgGroup := e.Group("/bsmg", echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(MemberClaims)
		},
		// ContextKey: "member", // 클레임의 이름 . default : user
		// SigningMethod: "RS256",      // 토큰 서명 방식 . defaelt : HMAC SHA-256 (HS256)
		SigningKey: []byte(AccessTokenKey),
		// TokenLookup: "header:Auth", // 헤더이름 Auth, Value에 Bearer 안써도 되게
		TokenLookup: "cookie:AC_bsmgCookie",
	}),
		echojwt.WithConfig(echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(MemberClaims)
			},
			SigningKey: []byte(RefreshTokenKey),
			// TokenLookup: "header:Auth", // 헤더이름 Auth, Value에 Bearer 안써도 되게
			TokenLookup: "cookie:RS_bsmgCookie",
		}),
	)
	// Route
	initRouteGroup(bsmgGroup)

	e.Logger.Fatal(e.Start(":3000"))
}
