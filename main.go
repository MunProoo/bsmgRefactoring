package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// db
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
	e.Use(sessionMiddleware)                  // 세션관리 추가

	// 시작 페이지 핸들러 함수
	e.GET("/", func(c echo.Context) error {
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
