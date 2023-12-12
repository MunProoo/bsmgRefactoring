package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

type SessionData struct {
	Authenticated bool
	ID            string
	RememberedID  string
}

// 세션 초기화 작업
func sessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		// w := c.Response().Writer

		session, err := store.Get(r, "Member")
		if err != nil {
			return err
		}

		// 세셩에서 사용자 ID 저장되있으면 로그인상태
		c.Set("Member", session)
		return next(c)
	}
}

func chkSession(c echo.Context) (data SessionData) {
	// 세션체크
	session := c.Get("Member").(*sessions.Session)
	if session.Values["mem_id"] == nil {
		// 로그인 페이지로 리다이렉트 or 에러처리
		// c.Redirect(http.StatusFound, "/bsmg/login")
		return
	}

	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		data.Authenticated = true
		data.ID = session.Values["mem_id"].(string)
	} else {
		data.Authenticated = false
	}

	data.RememberedID = session.Values["rememberedID"].(string)
	return

}
