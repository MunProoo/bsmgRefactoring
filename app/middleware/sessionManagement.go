package middleware

import (
	"BsmgRefactoring/define"
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// echo framework로 구조 변경후엔 토큰 기반 세션으로 바꿔볼까
const (
	SessionKey = "BSMG"
)

var (
	Key   = []byte("suuuuper-secret-key")
	Store = sessions.NewCookieStore(Key)
)

/*
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
*/
// 세션 생성 및 체크
func InitSessionMiddleware(c echo.Context) error {
	session, err := session.Get(SessionKey, c)
	if session == nil {
		return err
	}

	// 세션 Init
	if _, ok := session.Values["initialized"]; ok {
		return nil
	}

	url := c.Request().URL.RawQuery
	fmt.Println("url = ", url)
	session.Options = &sessions.Options{
		Path:     "/",   // 모든 경로에서 세션 확인
		MaxAge:   86400, // 하루
		HttpOnly: true,
		// SameSite: http.SameSiteNoneMode,
	}
	session.Values["initialized"] = true
	session.Save(c.Request(), c.Response())
	return err
}

// func checkSessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		session, _ := session.Get(SessionKey, c)

// 		if c.Request().URL.RawQuery == "/" {
// 			return next(c)
// 		}

// 		if session.Values["Member"] == nil || !session.Values["authenticated"].(bool) {
// 			return c.Redirect(http.StatusOK, "/")
// 		}

//			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
//				result := define.OnlyResult{}
//				result.Result.ResultCode = define.NotAuthorizedUser
//				return c.JSON(http.StatusOK, result)
//			}
//			return next(c)
//		}
//	}
func CheckSession(c echo.Context) bool {
	session, err := session.Get(SessionKey, c)
	if err != nil {
		return false
	}

	if session.Values["Member"] == nil || !session.Values["authenticated"].(bool) {
		return false
	}

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		result := define.OnlyResult{}
		result.Result.ResultCode = define.ErrorNotAuthorizedUser
		return false
	}
	return true
}

func CreateSession(c echo.Context, member *define.BsmgMemberInfo) *sessions.Session {
	log.Println("Session 생성!!")
	session, err := session.Get(SessionKey, c)
	if err != nil {
		log.Printf("%v \n", err)
	}

	if _, ok := session.Values["initialized"]; ok {
		// Set session
		session.Values["authenticated"] = true
		session.Values["Member"] = member

		err = session.Save(c.Request(), c.Response())
		if err != nil {
			// 로그
			log.Printf("%v \n", err)
		}
	}
	return session
}

func DeleteSession(c echo.Context) {
	session, err := session.Get(SessionKey, c)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	// delete session
	session.Values["authenticated"] = false
	session.Values["Member"] = nil
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response())
}

func GetSessionData(c echo.Context) (result define.BsmgMemberResponse) {
	session, err := session.Get(SessionKey, c)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	member := session.Values["Member"].(define.BsmgMemberInfo)
	result.MemberInfo = define.BsmgMemberInfo{}
	result.MemberInfo.Mem_ID = member.Mem_ID
	result.MemberInfo.Mem_Name = member.Mem_Name
	result.MemberInfo.Mem_Rank = member.Mem_Rank
	result.MemberInfo.Mem_Part = member.Mem_Part
	return
}

func IsDuplicatedLogin(c echo.Context, loginID string, sessionMap map[string]*sessions.Session) bool {
	presentSession, err := session.Get(SessionKey, c)
	if err != nil {
		log.Printf("%v", err)
		return false
	}

	// 로그인 기록이 있음
	if existingSession, ok := sessionMap[loginID]; ok {

		// 브라우저가 동일
		if strings.EqualFold(presentSession.ID, existingSession.ID) {
			log.Printf("Same Browser Duplicate Login : %s", loginID)
			return false
		}

		// 다른 브라우저에서의 요청 -> 기존 브라우저 연결 해제
		existingSession.Values["authenticated"] = false
		existingSession.Values["Member"] = nil
		existingSession.Options.MaxAge = -1
		existingSession.Save(c.Request(), c.Response())

	}

	member, exist := presentSession.Values["Member"]
	if !exist {
		return true
	}

	sessionID := member.(define.BsmgMemberInfo).Mem_ID

	return sessionID != loginID
}
