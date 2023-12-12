package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("views/webRoot"))
	e.Use(sessionMiddleware)

	// 시작 페이지 핸들러 함수
	e.GET("/", func(c echo.Context) error {
		return c.File("index.html") // 예시로 index.html 파일을 반환하도록 설정
	})

	// login URL 그룹화
	bsmgLoginGroup := e.Group("/bsmg/login")

	// 그룹 내에서 처리할 핸들러 함수 등록
	bsmgLoginGroup.GET("/chkLogin", getChkLoginRequest)

	// bsmgUserGroup.GET("/profile", handleUserProfile)
	// bsmgUserGroup.GET("/settings", handleUserSettings)
	// bsmgUserGroup.POST("/update", handleUserUpdate)

	// Route
	// fileServer := http.FileServer(http.Dir("./webRoot"))

	e.Logger.Fatal(e.Start(":3000"))
}
