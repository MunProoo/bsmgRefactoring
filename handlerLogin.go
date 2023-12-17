package main

import (
	"BsmgRefactoring/define"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 로그인 중인지 확인
func getChkLoginRequest(c echo.Context) error {
	log.Println("getChkLogin Req")
	var result define.BsmgMemberResult
	isAuthenticated := chkSession(c)
	if !isAuthenticated {
		result.Result.ResultCode = define.ErrorInvalidParameter
	} else {

		result = getSessionData(c)
		result.Result.ResultCode = define.Success
	}
	return c.JSON(http.StatusOK, result)
}

func postLoginRequest(c echo.Context) error {
	log.Println("postLoginRequest")
	var result *define.BsmgMemberResult
	result = &define.BsmgMemberResult{}

	value, err := c.FormParams()
	if err != nil {
		log.Printf("%v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	parser := initFormParser(value)
	if parser == nil {
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	result = parseLoginRequest(parser)

	result.Result.ResultCode = define.Success
	result.MemberInfo.Mem_Name = "뀨뀨"
	result.MemberInfo.Mem_Rank = "관리자"

	err = initSession(c)
	if err != nil {
		log.Printf("%v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}
	// 세션 생성
	createSession(c, result.MemberInfo)

	// 테스트용으로 무조건 통과되게
	return c.JSON(http.StatusOK, result)
}

func postLogoutRequest(c echo.Context) error {
	log.Println("postLogoutRequest")
	result := define.OnlyResult{}
	// 에러코드 정리 필요
	deleteSession(c)
	result.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, result)
}
