package main

import (
	"BsmgRefactoring/define"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 로그인 중인지 확인
func getChkLoginRequest(c echo.Context) error {
	log.Println("getChkLogin Req")
	var result define.BsmgMemberResponse
	isAuthenticated := checkSession(c)
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
	request := &define.BsmgMemberRequest{}
	response := &define.BsmgMemberResponse{}

	err := c.Bind(request)
	if err != nil {
		log.Printf("%v \n", err)
		response.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, response)
	}
	// value, err := c.FormParams()
	// if err != nil {
	// 	log.Printf("%v \n", err)
	// 	result.Data.Result.ResultCode = define.ErrorInvalidParameter
	// 	return c.JSON(http.StatusOK, result)
	// }

	// parser := initFormParser(value)
	// if parser == nil {
	// 	result.Data.Result.ResultCode = define.ErrorInvalidParameter
	// 	return c.JSON(http.StatusOK, result)
	// }

	response.Result.ResultCode = define.Success
	response.MemberInfo.Mem_ID = request.Data.MemberInfo.Mem_ID
	response.MemberInfo.Mem_Name = "뀨뀨"
	response.MemberInfo.Mem_Rank = 0
	response.MemberInfo.Mem_Part = 0

	// 세션 생성
	createSession(c, &response.MemberInfo)

	// 테스트용으로 무조건 통과되게
	return c.JSON(http.StatusOK, response)
}

func postLogoutRequest(c echo.Context) error {
	log.Println("postLogoutRequest")
	result := define.OnlyResult{}
	// 에러코드 정리 필요
	deleteSession(c)
	result.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, result)
}
