package handler

import (
	"BsmgRefactoring/define"
	"BsmgRefactoring/middleware"
	"BsmgRefactoring/server"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 로그인 중인지 확인
func (h *BsmgHandler) GetChkLoginRequest(c echo.Context) error {
	log.Println("getChkLogin Req")
	// JWT 검증  ------------------------------------------
	apiResponse, resultCode := h.uc.CheckLoginIng(c)
	if resultCode != define.Success {
		apiResponse.Result.ResultCode = int32(resultCode)
		return c.JSON(http.StatusOK, apiResponse)
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (h *BsmgHandler) PostLoginRequest(c echo.Context) error {
	log.Println("postLoginRequest")

	server, _ := c.Get("Server").(*server.ServerProcessor)

	apiRequest := define.BsmgMemberLoginRequest{}
	apiResponse := define.BsmgMemberResponse{}

	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	err := c.Bind(&apiRequest)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerLogin", "err": err})
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse = h.uc.CheckUserLogin(c, apiRequest)

	return c.JSON(http.StatusOK, echo.Map{
		// "token":         token, // 쿠키에 저장하면 안보내도 되긴 함
		"dm_memberInfo": apiResponse.MemberInfo,
		"Result":        apiResponse.Result,
	})
}

func (h *BsmgHandler) PostLogoutRequest(c echo.Context) error {
	log.Println("postLogoutRequest")
	result := h.uc.UserLogout(c)

	return c.JSON(http.StatusOK, result)
}
