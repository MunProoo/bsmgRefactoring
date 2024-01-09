package handler

import (
	"BsmgRefactoring/app/define"
	"BsmgRefactoring/app/middleware"
	"BsmgRefactoring/app/server"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func getUserListRequest(c echo.Context) error {
	log.Println("getUserList Req")

	result := &define.BsmgMemberListResponse{}

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	// DB에서 가져오는거로 변경
	userList, err := server.DBManager.DBGorm.SelectUserList()
	if err != nil {
		result.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, result)
	}
	count := len(userList)
	result.TotalCount.Count = int32(count)
	if count > 0 {
		result.MemberList = userList
	}

	result.Result.ResultCode = define.Success
	return c.JSON(http.StatusOK, result)
}

// 아이디 중복체크 확인
func getIdCheckRequest(c echo.Context) (err error) {
	log.Println("getIdCheckRequest")

	var apiResponse define.OnlyResult

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	// dm에 넣어서 전송중이므로 이렇게 받아야함.
	// TODO : parameter로 추가 (offset처럼)
	memID := c.Request().FormValue("@d1#mem_id")

	isExist, err := server.DBManager.DBGorm.CheckMemberIDDuplicate(memID)
	if err != nil {
		// record not found는 nil로 오도록 처리함. 그외 문제만 DB에러로
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	if isExist {
		apiResponse.Result.ResultCode = define.ErrorDuplicatedID
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.Result.ResultCode = define.Success
	return c.JSON(http.StatusOK, apiResponse)
}

// 사용자 검색
func getUserSearchRequest(c echo.Context) error {
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

	memberList, err := server.DBManager.DBGorm.SelectMemberListSearch(searchData)
	if err != nil {
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
func postUserReq(c echo.Context) error {
	log.Println("postUserReq")

	var apiRequest define.BsmgMemberRequest
	var apiResponse define.BsmgMemberResponse

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	// value, err := c.FormParams()
	// if err != nil {
	// 	log.Printf("%v \n", err)
	// 	apiResponse.Result.ResultCode = define.ErrorInvalidParameter
	// 	return c.JSON(http.StatusOK, apiResponse)
	// }
	// parser := InitFormParser(value)
	// member, err = parseUserRegistRequest(parser)
	err := c.Bind(&apiRequest)
	if err != nil {
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}

	member := apiRequest.Data.MemberInfo
	// argon2 사용하여 salting, hashing
	// Pass the plaintext password and parameters to our generateFromPassword
	encodedHash, err := middleware.GenerateFromPassword(member.Mem_Password)
	if err != nil {
		log.Printf("%v \n", err)
		// TODO : 암호화 전용 에러코드 생성 필요
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}

	member.Mem_Password = encodedHash

	err = server.DBManager.DBGorm.InsertMember(member)
	if err != nil {
		log.Printf("%v \n", err)
		apiResponse.Result.ResultCode = define.ErrorDataBase
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

	server := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	err := c.Bind(&apiRequest)
	if err != nil {
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}
	var member define.BsmgMemberInfo

	for _, reqMember := range apiRequest.Data.MemberList {
		member = reqMember.ParseMember()
		server.DBManager.DBGorm.UpdateUser(member)
	}

	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}

// 사용자 삭제
func deleteUserReq(c echo.Context) (err error) {
	log.Println("deleteUserReq")

	var apiResponse define.OnlyResult

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	memID := c.Param("memID")

	// 사용자는 지워도 그 사람의 업무보고는 남겨야지. 기록이니까
	err = server.DBManager.DBGorm.DeleteMember(memID)
	if err != nil {
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}
