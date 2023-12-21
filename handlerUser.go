package main

import (
	"BsmgRefactoring/define"
	"fmt"
	"net/http"
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

// 사용자 등록
func postUserReq(c echo.Context) error {
	log.Println("postUserReq")

	var result define.BsmgMemberRequest
	var member *define.BsmgMemberInfo
	server := c.Get("Server").(*ServerProcessor)

	value, err := c.FormParams()
	if err != nil {
		log.Printf("%v \n", err)
		result.Data.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}
	parser := initFormParser(value)
	member, err = parseUserRegistRequest(parser)

	if err != nil {
		log.Printf("%v \n", err)
		result.Data.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	// 내용 DB에 INSERT 작업해야함

	fmt.Printf("%v ", member)
	err = server.dbManager.DBGorm.InsertMember(*member)
	if err != nil {
		log.Printf("%v \n", err)
		result.Data.Result.ResultCode = define.DataBaseError
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusOK, result)
}

// 사용자 수정
func putUserReq(c echo.Context) error {
	log.Println("putUserReq")

	var req define.BsmgMemberRequest
	c.Bind(&req)
	fmt.Printf("%v ", req)
	var res define.BsmgMemberResponse
	var member *define.BsmgMemberInfo
	server := c.Get("server").(*ServerProcessor)

	value, err := c.FormParams()
	if err != nil {
		log.Printf("%v \n", err)
		res.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, res)
	}
	parser := initFormParser(value)
	member, err = parseUserRegistRequest(parser)

	if err != nil {
		log.Printf("%v \n", err)
		res.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, res)
	}

	// 내용 DB에 Update 작업해야함

	fmt.Printf("%v ", member)
	server.dbManager.DBGorm.InsertMember(*member)

	return c.JSON(http.StatusOK, res)
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
