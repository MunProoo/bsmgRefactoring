package main

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// db 연결, 미들웨어 연결 등은 server state에 맞춰서 go routine으로 기동하도록 변경할까
	var server ServerProcessor
	err := server.ConnectDataBase()
	if err != nil {
		// 로그
		log.Printf("ConnectDataBase Failed . err = %v\n", err)
		return
	}

	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("views/webRoot")) // eXBuilder6 의존성 파일 추가
	e.Use(session.Middleware(store))          // 세션 미들웨어 추가
	e.Use(initSessionMiddleware)

	// 세션 확인
	e.GET("/", func(c echo.Context) error {
		// initSessionMiddleware(c)

		return c.File("index.html")
	})

	// login URL 그룹화
	bsmgLoginGroup := e.Group("/bsmg/login")

	// 그룹 내에서 처리할 핸들러 함수 등록
	bsmgLoginGroup.GET("/chkLogin", getChkLoginRequest)
	bsmgLoginGroup.POST("/login", postLoginRequest)

	// bsmgUserGroup.GET("/profile", handleUserProfile)
	// bsmgUserGroup.GET("/settings", handleUserSettings)
	// bsmgUserGroup.POST("/update", handleUserUpdate)

	// Route
	// fileServer := http.FileServer(http.Dir("./webRoot"))

	e.Logger.Fatal(e.Start(":3000"))
}
