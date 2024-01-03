package main

import (
	"BsmgRefactoring/define"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// 사용자 업무보고 속성트리
func getAttrTreeReq(c echo.Context) (err error) {
	log.Println("getAttrTreeReq")

	var result define.BsmgTreeResult

	server.mutex.Lock()
	defer server.mutex.Unlock()
	result.AttrTreeList, err = server.dbManager.DBGorm.MakeAttrTree()
	if err != nil {
		log.Printf("%v \n", err)
		return err
	}
	result.PartTreeList = make([]define.PartTree, 0)

	return c.JSON(http.StatusOK, result)
}

// 직급, 부서 정보
func getRankPartReq(c echo.Context) error {
	log.Println("getRankPartReq")

	var result define.BsmgRankPartResult
	server := c.Get("Server").(*ServerProcessor) // 서버가 전역 변수인데 context에서 꺼내서 쓸까말까@@@@@@@@@@@@@@@@@@@@@

	server.mutex.Lock()
	defer server.mutex.Unlock()
	rankList, err := server.dbManager.DBGorm.SelectRankList()
	if err != nil {
		log.Printf("%v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}
	result.RankList = rankList

	partList, err := server.dbManager.DBGorm.SelectPartist()
	if err != nil {
		log.Printf("%v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}
	result.PartList = partList

	result.Result.ResultCode = define.Success
	return c.JSON(http.StatusOK, result)
}

// 주간 업무보고 카테고리 정보
func getPartTree(c echo.Context) (err error) {
	log.Println("getPartTree")

	var apiResponse define.BsmgTreeResult

	server.mutex.Lock()
	defer server.mutex.Unlock()

	apiResponse.PartTreeList, err = server.dbManager.DBGorm.MakePartTree()
	if err != nil {
		log.Printf("%v \n", err)
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.Result.ResultCode = define.Success
	return c.JSON(http.StatusOK, apiResponse)
}

// 주간 업무보고를 수정한다면, 해당 팀의 팀장으로 바로 보고자 변경하기 위해
// 팀장을 response
func getToRptReq(c echo.Context) (err error) {
	log.Println("getToRptReq")

	var apiResponse define.BsmgTeamLeaderResponse

	server.mutex.Lock()
	defer server.mutex.Unlock()

	partIdx, _ := strconv.Atoi(c.Request().FormValue("@d1#part_idx"))

	apiResponse.Part.TeamLeader, err = server.dbManager.DBGorm.SelectPartLeader(int32(partIdx))
	if err != nil {
		log.Printf("getToRptReq : %v \n", err)
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}
	apiResponse.Part.PartIdx = int32(partIdx)
	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}

func getAttr1Req(c echo.Context) (err error) {
	log.Println("getAttr1Req")

	var apiResponse define.BsmgAttr1Response

	server.mutex.Lock()
	defer server.mutex.Unlock()

	apiResponse.Attr1List, err = server.dbManager.DBGorm.SelectAttr1List()
	if err != nil {
		log.Printf("getAttr1Req : %v \n", err)
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}
