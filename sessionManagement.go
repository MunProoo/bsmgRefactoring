package main

import (
	"BsmgRefactoring/define"
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// echo framework로 구조 변경후엔 토큰 기반 세션으로 바꿔볼까

var (
	key   = []byte("suuuuper-secret-key")
	store = sessions.NewCookieStore(key)
)

// 세션 생성 및 체크
func initSessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if sess == nil {
			return err
		}

		// 세션 Init
		if _, ok := sess.Values["initialized"]; !ok {
			sess.Options = &sessions.Options{
				Path:     "/",   // 모든 경로에서 세션 확인
				MaxAge:   86400, // 하루
				HttpOnly: true,
			}
			sess.Values["initialized"] = true
		}

		chkSession(c)

		// 다음 미들웨어 호출 or 핸들러에 제어 전달 next(c)
		return next(c)
	}
}

func chkSession(c echo.Context) (authentication bool) {
	// 세션체크
	// session := c.Get(sessions.).(*sessions.Session) // 여기서 BSMG가 nil이라 타입변환 시 죽음!!!!!!!!!!!!!!!!!!!!!!!!!
	session, err := session.Get(sessionKey, c)
	if err != nil {
		log.Printf("%v \n", err)
		c.Redirect(http.StatusOK, "/bsmg/login")
		return
	}

	if session.Values["Member"] == nil || !session.Values["authenticated"].(bool) {
		c.Redirect(http.StatusOK, "/bsmg/login")
		return
	}

	member := session.Values["Member"].(define.BsmgMemberInfo)
	if member.Mem_ID == "" {
		// 로그인 페이지로 리다이렉트 or 에러처리
		c.Redirect(http.StatusOK, "/bsmg/login")
		return
	}

	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		authentication = true
	}

	return

}

func createSession(c echo.Context, member *define.BsmgMemberInfo) {
	// session, _ := c.Get("BSMG").(*sessions.Session)
	session, err := session.Get(sessionKey, c)
	if err != nil {
		log.Printf("%v \n", err)
	}

	// Set session
	session.Values["authenticated"] = true
	session.Values["Member"] = member

	err = session.Save(c.Request(), c.Response())
	if err != nil {
		// 로그
		log.Printf("%v \n", err)
	}
}

// 세션ID 생성용
func generateSessionID() string {
	// 16 바이트 난수 생성
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Printf("%v", err)
		return ""
	}

	// 16바이트를 32자의 16진수 문자열로 변환
	uniqueID := hex.EncodeToString(randomBytes)
	return uniqueID
}

func deleteSession(c echo.Context) {
	session, err := session.Get("BSMG", c)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	// delete session
	session.Values["authenticated"] = false
	session.Values["Member"] = nil
	session.Save(c.Request(), c.Response())
}

func getSessionData(c echo.Context) (result define.BsmgMemberResult) {
	session, err := session.Get(sessionKey, c)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	member := session.Values["Member"].(define.BsmgMemberInfo)
	result.MemberInfo = &define.BsmgMemberInfo{}
	result.MemberInfo.Mem_ID = member.Mem_ID
	result.MemberInfo.Mem_Name = member.Mem_Name
	result.MemberInfo.Mem_Rank = "관리자"
	// result.MemberInfo.Mem_Rank = member.Mem_Rank
	result.MemberInfo.Mem_Part = member.Mem_Part
	return
}

func initSession(c echo.Context) (err error) {
	sess, err := session.Get("session", c)
	if sess == nil {
		return err
	}

	// 세션 Init
	if _, ok := sess.Values["initialized"]; !ok {
		sess.Options = &sessions.Options{
			Path:     "/",   // 모든 경로에서 세션 확인
			MaxAge:   86400, // 하루
			HttpOnly: true,
		}
		sess.Values["initialized"] = true
	}

	return
}
