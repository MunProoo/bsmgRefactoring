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
	apiRequest := &define.BsmgMemberLoginRequest{}
	apiResponse := &define.BsmgMemberResponse{}

	err := c.Bind(apiRequest)
	if err != nil {
		log.Printf("%v \n", err)
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}

	// member := request.Data.MemberInfo.ParseMember()
	member := apiRequest.Data.MemberInfo

	// member 바뀌는지 확인 필요
	err = server.dbManager.DBGorm.Login(&member)
	if err != nil {
		if err.Error() == "record not found" {
			// 웹에서 에러코드 통해서 아이디 혹은 비밀번호가 틀립니다로
			apiResponse.Result.ResultCode = define.ErrorLoginFailed
			return c.JSON(http.StatusOK, apiResponse)
		}
	}

	apiResponse.Result.ResultCode = define.Success
	apiResponse.MemberInfo = member

	// 세션 생성
	createSession(c, &apiResponse.MemberInfo)

	// 테스트용으로 무조건 통과되게
	return c.JSON(http.StatusOK, apiResponse)
}

func postLogoutRequest(c echo.Context) error {
	log.Println("postLogoutRequest")
	result := define.OnlyResult{}
	// 에러코드 정리 필요
	deleteSession(c)
	result.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, result)
}
