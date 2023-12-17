package main

import (
	"BsmgRefactoring/define"
	"encoding/gob"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	sessionKey = "BSMG"
)

var server ServerProcessor // 핸들러에서도 접근 가능하도록 전역으로 할당

func init() {
	// Gob에 BsmgMemberInfo 타입 등록
	gob.Register(define.BsmgMemberInfo{})
}

func main() {
	// TODO : server를 goroutine으로 돌려야함
	// DB 채널통신 테스트중
	server.reqCh = make(chan interface{}, 6000)
	server.resCh = make(chan interface{}, 6000)

	go server.StartServer()

	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("views/webRoot")) // eXBuilder6 의존성 파일 추가

	e.Use(session.Middleware(store)) // 세션 미들웨어 추가
	// 스레드 안정성을 고려하여 서버 변수를 컨텍스트에 할당
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("server", &server)
			return next(c)
		}
	})

	e.GET("/", func(c echo.Context) error {
		c.Set("server", &server)
		return c.File("index.html")
	})

	// URL 그룹화
	bsmgGroup := e.Group("/bsmg/")

	// Route
	initRouteGroup(bsmgGroup)

	e.Logger.Fatal(e.Start(":3000"))
}
