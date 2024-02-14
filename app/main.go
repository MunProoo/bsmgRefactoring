package main

// SPA (클라이언트 측 페이지 전환) 방식이므로 세션체크는 클라이언트에서 한다.

import (
	"BsmgRefactoring/define"
	"BsmgRefactoring/handler"
	bsmgMd "BsmgRefactoring/middleware"
	"BsmgRefactoring/repository"

	"BsmgRefactoring/router"
	"BsmgRefactoring/server"
	"BsmgRefactoring/usecase"

	"encoding/gob"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "BsmgRefactoring/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// var server ServerProcessor // 핸들러에서도 접근 가능하도록 전역으로 할당

func init() {
	// Gob에 BsmgMemberInfo 타입 등록
	gob.Register(define.BsmgMemberInfo{})
}

const (
	swaggerIndexURL = "/swagger/index.html"
)

// @title BSMG Swagger API
// @version 1.0
// @host localhost:3000
// @BasePath /bsmg

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	// 로그 시작
	bsmgMd.InitLog()

	server := server.InitServer()

	e := echo.New()

	// DBManager 생성
	server.ConnectDataBase() // 서브 고루틴에서 연결 시 비동기이므로.. 첫 연결은 절차적으로
	// dbManager := server.DBManager

	// Repository 생성 DB 의존성 주입
	// repo := repository.NewBsmgRepository(dbManager)
	repo := repository.NewBsmgRepository()

	// UseCase 생성 및 Repository 의존성 주입
	useCase := usecase.NewBsmgUsecase(repo)
	server.CreateCron(useCase) // 주간보고 스케쥴러 생성

	// handler 생성 및 Usecase 의존성 주입
	bsmgHandler := handler.NewBsmgHandler(useCase)

	// Swagger Docs 추가
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 서버변수 컨텍스트에 할당
			c.Set("Server", server)

			// Swagger UI를 불러오는 URL (/swagger/index.html)이 정적파일의 index.html과 충돌하여 동작하지 않으므로 후킹 추가
			if c.Request().URL.Path == swaggerIndexURL {
				return echoSwagger.WrapHandler(c)
			}
			return next(c)
		}
	})

	// middleware ------------------------------------
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("views/webRoot")) // eXBuilder6 의존성 파일 추가
	// e.Use(middleware.Static("views/bsmgApp/webRoot")) // 빠른 디버깅용. 배포위치를 변경하여 front 수정 시 바로 반영되도록

	/*
		// JWT로 변경하면서 사용 X. 브라우저간 세션 쿠키 공유 안되는 문제있음
		e.Use(session.Middleware(bsmgMd.Store)) // 세션 미들웨어 추가
		e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				// 세션 초기화
				bsmgMd.InitSessionMiddleware(c)
				// 스레드 안정성을 고려하여 서버 변수를 컨텍스트에 할당
				c.Set("Server", server)
				return next(c)
			}
		})
	*/

	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})

	// login API만 JWT 인증 제외하기 위해 JWT 미들웨어보다 먼저 선언
	e.POST("/bsmg/login/login", bsmgHandler.PostLoginRequest)
	e.GET("/bsmg/login/chkLogin", bsmgHandler.GetChkLoginRequest)

	// // URL 그룹화 + JWT 미들웨어 적용
	bsmgGroup := e.Group("/bsmg", echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(bsmgMd.MemberClaims)
		},
		// ContextKey: "member", // 클레임의 이름 . default : user
		// SigningMethod: "RS256",      // 토큰 서명 방식 . defaelt : HMAC SHA-256 (HS256)
		SigningKey: []byte(bsmgMd.AccessTokenKey),
		// TokenLookup: "header:Auth", // 헤더이름 Auth, Value에 Bearer 안써도 되게
		TokenLookup: "cookie:AC_bsmgCookie",
	}),
		echojwt.WithConfig(echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(bsmgMd.MemberClaims)
			},
			SigningKey: []byte(bsmgMd.RefreshTokenKey),
			// TokenLookup: "header:Auth", // 헤더이름 Auth, Value에 Bearer 안써도 되게
			TokenLookup: "cookie:RS_bsmgCookie",
		}),
	)
	// Route
	router.InitRouteGroup(bsmgGroup, *bsmgHandler)

	go server.StartServer(repo) // DB 상태 체크 및 재연결 등..

	e.Logger.Fatal(e.Start(":3000"))
}
