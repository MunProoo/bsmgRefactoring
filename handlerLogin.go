package main

import (
	"BsmgRefactoring/define"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// 로그인 중인지 확인
func getChkLoginRequest(c echo.Context) error {
	log.Println("getChkLogin Req")
	var apiResponse define.BsmgMemberResponse

	// JWT 검증  ------------------------------------------

	claims, err := checkToken(c)
	if err != nil {
		log.Printf("%v\n", err)
		apiResponse.Result.ResultCode = define.ErrorInvalidToken
		return c.JSON(http.StatusOK, apiResponse)
	}

	err = apiResponse.MemberInfo.ParsingClaim(claims)
	if err != nil {
		log.Printf("%v\n", err)
		apiResponse.Result.ResultCode = define.ErrorInvalidToken
		return c.JSON(http.StatusOK, apiResponse)
	}

	/*
		세션을 통한 사용자 정보 Get은 필요없음
		isAuthenticated := checkSession(c)
		if !isAuthenticated {
			result.Result.ResultCode = define.ErrorInvalidParameter
		} else {

			result = getSessionData(c)
			result.Result.ResultCode = define.Success
		}
	*/
	return c.JSON(http.StatusOK, apiResponse)
}

func postLoginRequest(c echo.Context) error {

	log.Println("postLoginRequest")
	apiRequest := &define.BsmgMemberLoginRequest{}
	apiResponse := &define.BsmgMemberResponse{}

	server.mutex.Lock()
	defer server.mutex.Unlock()

	err := c.Bind(apiRequest)
	if err != nil {
		log.Printf("%v \n", err)
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}

	// member := request.Data.MemberInfo.ParseMember()
	member := apiRequest.Data.MemberInfo

	err = server.dbManager.DBGorm.SelectMemberInfo(&member)
	if err != nil {
		if err.Error() == "record not found" {
			// 웹에서 에러코드 통해서 아이디 혹은 비밀번호가 틀립니다로
			apiResponse.Result.ResultCode = define.ErrorLoginFailed
			return c.JSON(http.StatusOK, apiResponse)
		}
	}

	// 비밀번호 매칭 확인
	match, err := comparePasswordAndHash(apiRequest.Data.MemberInfo.Mem_Password, member.Mem_Password)
	if err != nil || !match {
		apiResponse.Result.ResultCode = define.ErrorLoginFailed
		return c.JSON(http.StatusOK, apiResponse)
	}
	// 인증 성공 ----------------------

	// 중복 로그인 확인
	if !isNotDuplicateLogin(c, member.Mem_ID) {
		log.Printf("%v \n", "중복 로그인입니다.")
		apiResponse.Result.ResultCode = define.ErrorLoginDuplication
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.Result.ResultCode = define.Success
	apiResponse.MemberInfo = member

	// JWT 토큰 생성
	// Set custom claims
	claims := &MemberClaims{
		member.Mem_ID,
		member.Mem_Name,
		member.Mem_Rank,
		member.Mem_Part,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	// AccessToken, RefreshToken , 쿠키 생성
	makeJwtToken(c, claims)
	if err != nil {
		apiResponse.Result.ResultCode = define.ErrorTokenCreationFailed
		return c.JSON(http.StatusOK, apiResponse)
	}

	// 세션 생성
	createSession(c, &apiResponse.MemberInfo)

	return c.JSON(http.StatusOK, echo.Map{
		// "token":         token, // 쿠키에 저장하면 안보내도 되긴 함
		"dm_memberInfo": apiResponse.MemberInfo,
		"Result":        apiResponse.Result,
	})
}

func postLogoutRequest(c echo.Context) error {
	log.Println("postLogoutRequest")
	result := define.OnlyResult{}

	deleteCookie(c, AccessCookieName)
	deleteCookie(c, RefreshCookieName)

	deleteSession(c)
	result.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, result)
}
