package main

import (
	"BsmgRefactoring/define"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func getUserListRequest(c echo.Context) error {
	log.Println("getUserList Req")

	result := &define.BsmgMemberResponse{}

	// DB에서 가져오는거로 변경
	userList, err := server.dbManager.DBGorm.SelectUserList()
	if err != nil {
		result.Result.ResultCode = define.DataBaseError
		return c.JSON(http.StatusOK, result)
	}
	count := len(userList)
	result.TotalCount.Count = int32(count)
	if count > 0 {
		result.MemberList = userList
	}

	result.Result.ResultCode = define.Success

	// 테스트용으로 무조건 통과되게
	return c.JSON(http.StatusOK, result)
}

// 아이디 중복체크 확인
func getIdCheckRequest(c echo.Context) error {
	log.Println("getIdCheckRequest")

	var result define.BsmgMemberRequest

	value, err := c.FormParams()
	if err != nil {
		log.Printf("%v \n", err)
		result.Data.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}
	parser := initFormParser(value)
	mem_id, err := parser.getStringValue(0, "mem_id", 0)

	// db에 있는지 확인

	fmt.Println(mem_id)

	result.Data.Result.ResultCode = define.Success

	// 테스트용으로 무조건 통과되게
	return c.JSON(http.StatusOK, result)
}

// 사용자 검색
func getUserSearchRequest(c echo.Context) error {
	log.Println("getUserSearchRequest")

	var result define.BsmgMemberResponse
	var search *define.SearchData

	value, err := c.FormParams()
	if err != nil {
		log.Printf("%v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}
	parser := initFormParser(value)
	search = parseSearchRequest(parser)

	fmt.Printf("%v ", search)

	return c.JSON(http.StatusOK, result)
}

// 사용자 등록 -> 통신 json으로 변경필요
func postUserReq(c echo.Context) error {
	log.Println("postUserReq")

	// var apiRequest define.BsmgMemberRequest
	var apiResponse define.BsmgMemberResponse
	var member *define.BsmgMemberInfo
	server := c.Get("Server").(*ServerProcessor)

	value, err := c.FormParams()
	if err != nil {
		log.Printf("%v \n", err)
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}
	parser := initFormParser(value)
	member, err = parseUserRegistRequest(parser)

	if err != nil {
		log.Printf("%v \n", err)
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}

	err = server.dbManager.DBGorm.InsertMember(*member)
	if err != nil {
		log.Printf("%v \n", err)
		apiResponse.Result.ResultCode = define.DataBaseError
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}

// 사용자 수정
func putUserReq(c echo.Context) error {
	log.Println("putUserReq")

	var apiRequest define.BsmgPutMemberRequest
	var apiResponse define.OnlyResult

	err := c.Bind(&apiRequest)
	if err != nil {
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}
	var member define.BsmgMemberInfo
	server := c.Get("Server").(*ServerProcessor)

	for _, reqMember := range apiRequest.Data.MemberList {
		member.Mem_ID = reqMember.Mem_ID
		// member.Mem_Name = reqMember.Mem_Name
		rank, _ := strconv.Atoi(reqMember.Mem_Rank)
		member.Mem_Rank = int32(rank)
		part, _ := strconv.Atoi(reqMember.Mem_Part)
		member.Mem_Part = int32(part)

		setVal := make(map[string]interface{})
		setVal["mem_name"] = reqMember.Mem_Name
		setVal["mem_rank"] = rank
		setVal["mem_part"] = part

		server.dbManager.DBGorm.UpdateUser(setVal, member.Mem_ID)
	}

	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}

// 사용자 삭제
func deleteUserReq(c echo.Context) error {
	log.Println("deleteUserReq")

	var result define.BsmgMemberResponse
	// server := c.Get("server").(*ServerProcessor)
	url := c.Request().URL.Path[1:]
	reqSlice := strings.Split(url, "/")
	memID := reqSlice[3]

	// 내용 DB에 Delete 작업해야함
	// TODO : user를 서버 memory에 넣고, userID를 mem_Index로 변환하여 빠르게 처리작업

	fmt.Printf("%v ", memID)
	// server.dbManager.DBGorm.InsertMember(memID)

	return c.JSON(http.StatusOK, result)
}
