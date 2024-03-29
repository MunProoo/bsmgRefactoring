package handler

import (
	"BsmgRefactoring/define"
	"BsmgRefactoring/middleware"
	"BsmgRefactoring/server"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// 사용자 업무보고 속성트리
func (h *BsmgHandler) GetAttrTreeReq(c echo.Context) (err error) {
	log.Println("getAttrTreeReq")

	var result define.BsmgTreeResult

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()
	result.AttrTreeList, err = h.uc.MakeAttrTree()
	// result.AttrTreeList, err = server.DBManager.DBGorm.MakeAttrTree()
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerSetting", "err": err})
		return err
	}
	result.PartTreeList = make([]define.PartTree, 0)

	return c.JSON(http.StatusOK, result)
}

// 직급, 부서 정보
func (h *BsmgHandler) GetRankPartReq(c echo.Context) error {
	log.Println("getRankPartReq")

	var result define.BsmgRankPartResult
	server := c.Get("Server").(*server.ServerProcessor)

	server.Mutex.Lock()
	defer server.Mutex.Unlock()
	rankList, err := h.uc.SelectRankList()
	// rankList, err := server.DBManager.DBGorm.SelectRankList()
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerSetting", "err": err})
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}
	result.RankList = rankList

	partList, err := h.uc.SelectPartist()
	// partList, err := server.DBManager.DBGorm.SelectPartist()
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerSetting", "err": err})
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}
	result.PartList = partList

	result.Result.ResultCode = define.Success
	return c.JSON(http.StatusOK, result)
}

// 주간 업무보고 카테고리 정보
func (h *BsmgHandler) GetPartTree(c echo.Context) (err error) {
	log.Println("getPartTree")

	var apiResponse define.BsmgTreeResult

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	apiResponse.PartTreeList, err = h.uc.MakePartTree()
	// apiResponse.PartTreeList, err = server.DBManager.DBGorm.MakePartTree()
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerSetting", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.Result.ResultCode = define.Success
	return c.JSON(http.StatusOK, apiResponse)
}

// 주간 업무보고를 수정한다면, 해당 팀의 팀장으로 바로 보고자 변경하기 위해
// 팀장을 response
func (h *BsmgHandler) GetToRptReq(c echo.Context) (err error) {
	log.Println("getToRptReq")

	var apiResponse define.BsmgTeamLeaderResponse

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	partIdx, _ := strconv.Atoi(c.Request().FormValue("@d1#part_idx"))

	apiResponse.Part.TeamLeader, err = h.uc.SelectPartLeader(int32(partIdx))
	// apiResponse.Part.TeamLeader, err = server.DBManager.DBGorm.SelectPartLeader(int32(partIdx))
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerSetting", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}
	apiResponse.Part.PartIdx = int32(partIdx)
	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}

func (h *BsmgHandler) GetAttr1Req(c echo.Context) (err error) {
	log.Println("getAttr1Req")

	var apiResponse define.BsmgAttr1Response

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	apiResponse.Attr1List, err = h.uc.SelectAttr1List()
	// apiResponse.Attr1List, err = server.DBManager.DBGorm.SelectAttr1List()
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerSetting", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}
