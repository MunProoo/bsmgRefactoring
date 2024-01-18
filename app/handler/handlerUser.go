package handler

import (
	"BsmgRefactoring/define"
	"BsmgRefactoring/middleware"
	"BsmgRefactoring/server"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *BsmgHandler) GetUserListRequest(c echo.Context) error {
	log.Println("GetUserList Req")

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	// Handler -> Usecase -> Repository
	result := h.uc.GetUserListRequest()
	return c.JSON(http.StatusOK, result)
}

// 아이디 중복체크 확인
func (h *BsmgHandler) GetIdCheckRequest(c echo.Context) (err error) {
	log.Println("getIdCheckRequest")

	var apiResponse define.OnlyResult

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	// dm에 넣어서 전송중이므로 이렇게 받아야함.
	// TODO : parameter로 추가 (offset처럼)
	memID := c.Request().FormValue("@d1#mem_id")

	apiResponse = h.uc.GetIdCheckRequest(memID)
	return c.JSON(http.StatusOK, apiResponse)
}

// 사용자 검색
func (h *BsmgHandler) GetUserSearchRequest(c echo.Context) error {
	log.Println("getUserSearchRequest")

	var apiResponse define.BsmgMemberListResponse

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	var searchData define.SearchData
	searchCombo := c.Request().FormValue("@d1#search_combo")
	combo, _ := strconv.Atoi(searchCombo)
	searchData.SearchCombo = int32(combo)
	searchData.SearchInput = c.Request().FormValue("@d1#search_input")

	memberList, err := h.uc.SelectMemberListSearch(searchData)
	// memberList, err := server.DBManager.DBGorm.SelectMemberListSearch(searchData)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerUser", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.MemberList = memberList
	apiResponse.TotalCount.Count = int32(len(memberList))
	apiResponse.Result.ResultCode = define.Success

	fmt.Println(memberList, err)
	return c.JSON(http.StatusOK, apiResponse)
}

// 사용자 등록 -> 통신 json으로 변경필요
func (h *BsmgHandler) PostUserReq(c echo.Context) error {
	log.Println("postUserReq")

	var apiRequest define.BsmgMemberRequest
	var apiResponse define.OnlyResult

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	err := c.Bind(&apiRequest)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerUser", "err": err})
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse = h.uc.PostUserReq(c, apiRequest)
	return c.JSON(http.StatusOK, apiResponse)
}

// 사용자 수정
func (h *BsmgHandler) PutUserReq(c echo.Context) error {
	log.Println("putUserReq")

	var apiRequest define.BsmgPutMemberRequest
	var apiResponse define.OnlyResult

	server := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	err := c.Bind(&apiRequest)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerUser", "err": err})
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse = h.uc.PutUserReq(c, apiRequest)

	return c.JSON(http.StatusOK, apiResponse)
}

// 사용자 삭제
func (h *BsmgHandler) DeleteUserReq(c echo.Context) (err error) {
	log.Println("deleteUserReq")

	var apiResponse define.OnlyResult

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	// 인가 체크
	if _, resultCode := h.uc.AuthorizationCheck(c); resultCode != define.Success {
		apiResponse.Result.ResultCode = int32(resultCode)
		return c.JSON(http.StatusOK, apiResponse)
	}

	memID := c.Param("memID")

	// 사용자는 지워도 그 사람의 업무보고는 남겨야지. 기록이니까
	err = h.uc.DeleteMember(memID)
	// err = server.DBManager.DBGorm.DeleteMember(memID)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerUser", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}
