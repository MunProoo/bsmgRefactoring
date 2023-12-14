package main

import (
	"BsmgRefactoring/define"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

// 로그인 중인지 세션을 통해 확인
func getChkLoginRequest(c echo.Context) error {

	var result define.BsmgMemberResult
	chkSess := chkSession(c)
	if !chkSess.Authenticated {
		result.Result.ResultCode = 1

	} else {
		result.MemberInfo = &define.BsmgMemberInfo{}
		result.MemberInfo.Mem_ID = "mem_id"
		result.MemberInfo.Mem_Name = "name"
		result.MemberInfo.Mem_Rank = "rank"
		result.MemberInfo.Mem_Part = "part"
		result.Result.ResultCode = 0
	}

	return c.JSON(http.StatusOK, result)
}

func postLoginRequest(c echo.Context) error {
	// Get URL DATA는 queryParam
	ID := c.QueryParam("id")
	fmt.Println(ID)

	var result define.BsmgMemberResult

	memID := c.FormValue("mem_id")
	memPW := c.FormValue("mem_pw")
	value, err := c.FormParams()
	fmt.Println(err)
	memID = value.Get("mem_id")
	// 이렇게 받을수밖에없을까
	memID = value.Get("@d1#mem_id")

	result.MemberInfo.Mem_ID = memID
	result.MemberInfo.Mem_Name = "Test"
	result.MemberInfo.Mem_Password = memPW
	result.Result.ResultCode = 0

	var sessionValue define.SessionValue
	sessionValue.ID = memID

	createSession(c, sessionValue)

	// 세션으로부터 값 받아오기 테스트
	session := c.Get("Member").(*sessions.Session)
	sessionID := session.Values["mem_id"]

	fmt.Printf("%v", sessionID)
	return nil
}
