package handler

import (
	"BsmgRefactoring/define"
	"BsmgRefactoring/middleware"
	"BsmgRefactoring/server"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// 로그인 중인지 확인
func (h *BsmgHandler) GetChkLoginRequest(c echo.Context) error {
	log.Println("getChkLogin Req")
	var apiResponse define.BsmgMemberResponse

	// JWT 검증  ------------------------------------------

	claims, err := middleware.CheckToken(c)
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

func (h *BsmgHandler) PostLoginRequest(c echo.Context) error {
	log.Println("postLoginRequest")

	server, _ := c.Get("Server").(*server.ServerProcessor)

	apiRequest := &define.BsmgMemberLoginRequest{}
	apiResponse := &define.BsmgMemberResponse{}

	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	err := c.Bind(apiRequest)
	if err != nil {
		log.Printf("%v \n", err)
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}

	// member := request.Data.MemberInfo.ParseMember()
	member := apiRequest.Data.MemberInfo

	err = server.DBManager.DBGorm.SelectMemberInfo(&member)
	if err != nil {
		if err.Error() == "record not found" {
			// 웹에서 에러코드 통해서 아이디 혹은 비밀번호가 틀립니다로
			apiResponse.Result.ResultCode = define.ErrorLoginFailed
			return c.JSON(http.StatusOK, apiResponse)
		}
	}

	// 비밀번호 매칭 확인
	match, err := middleware.ComparePasswordAndHash(apiRequest.Data.MemberInfo.Mem_Password, member.Mem_Password)
	if err != nil || !match {
		apiResponse.Result.ResultCode = define.ErrorLoginFailed
		return c.JSON(http.StatusOK, apiResponse)
	}
	// 인증 성공 ----------------------

	// 중복 로그인 확인
	if !middleware.IsNotDuplicateLogin(c, member.Mem_ID) {
		log.Printf("%v \n", "중복 로그인입니다.")
		apiResponse.Result.ResultCode = define.ErrorLoginDuplication
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.Result.ResultCode = define.Success
	apiResponse.MemberInfo = member

	// JWT 토큰 생성
	// Set custom claims
	claims := &middleware.MemberClaims{
		Mem_ID:   member.Mem_ID,
		Mem_Name: member.Mem_Name,
		Mem_Rank: member.Mem_Rank,
		Mem_Part: member.Mem_Part,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	// AccessToken, RefreshToken , 쿠키 생성
	middleware.MakeJwtToken(c, claims)
	if err != nil {
		apiResponse.Result.ResultCode = define.ErrorTokenCreationFailed
		return c.JSON(http.StatusOK, apiResponse)
	}

	// 세션 생성
	middleware.CreateSession(c, &apiResponse.MemberInfo)

	return c.JSON(http.StatusOK, echo.Map{
		// "token":         token, // 쿠키에 저장하면 안보내도 되긴 함
		"dm_memberInfo": apiResponse.MemberInfo,
		"Result":        apiResponse.Result,
	})
}

func (h *BsmgHandler) PostLogoutRequest(c echo.Context) error {
	log.Println("postLogoutRequest")
	result := define.OnlyResult{}

	middleware.DeleteCookie(c, middleware.AccessCookieName)
	middleware.DeleteCookie(c, middleware.RefreshCookieName)

	middleware.DeleteSession(c)
	result.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, result)
}
