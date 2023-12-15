package main

import (
	"BsmgRefactoring/define"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 로그인 중인지 확인
func getChkLoginRequest(c echo.Context) error {

	var result define.BsmgMemberResult
	isAuthenticated := chkSession(c)
	if !isAuthenticated {
		result.Result.ResultCode = define.ErrorInvalidParameter
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
	var result *define.BsmgMemberResult
	result = &define.BsmgMemberResult{}

	value, err := c.FormParams()
	if err != nil {
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	parser := initFormParser(value)
	if parser == nil {
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	result = parseLoginRequest(parser)

	result.Result.ResultCode = 0
	result.MemberInfo.Mem_Name = "뀨뀨"

	// 세션 생성
	createSession(c, result.MemberInfo)

	// 테스트용으로 무조건 통과되게
	return c.JSON(http.StatusOK, result)
}
