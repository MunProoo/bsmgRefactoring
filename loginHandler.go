package main

import (
	"BsmgRefactoring/define"
	"net/http"

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
