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

// server 변수를 현재 context에서 꺼내쓰는데 타입에 대해서 불명확하기때문에 이건 지양해야 하는 방향.

func (h *BsmgHandler) GetReportSearchReq(c echo.Context) (err error) {
	log.Println("getReportSearchReq")

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	var searchData define.SearchData

	searchCombo := c.Request().FormValue("@d1#search_combo")
	combo, _ := strconv.Atoi(searchCombo)
	searchData.SearchCombo = int32(combo)
	searchData.SearchInput = c.Request().FormValue("@d1#search_input")

	offset, _ := strconv.Atoi(c.Request().FormValue("offset"))
	limit, _ := strconv.Atoi(c.Request().FormValue("limit"))

	pageInfo := define.PageInfo{
		Offset: int32(offset),
		Limit:  int32(limit),
	}

	apiResponse := h.uc.SelectReportListReq(searchData, pageInfo)

	return c.JSON(http.StatusOK, apiResponse)
}

// 업무 속성에 따른 보고 정보
func (h *BsmgHandler) GetReportAttrSearchReq(c echo.Context) (err error) {
	log.Println("getReportAttrSearchReq")

	var apiResponse define.BsmgReportListResponse

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	attrData := define.AttrSearchData{}
	attrValue, _ := strconv.Atoi(c.Request().FormValue("@d1#attrValue"))
	attrData.AttrValue = int32(attrValue)
	attrCategory, _ := strconv.Atoi(c.Request().FormValue("@d1#attrCategory"))
	attrData.AttrCategory = int32(attrCategory)

	pageInfo := define.PageInfo{}
	offset, _ := strconv.Atoi(c.Request().FormValue("offset"))
	pageInfo.Offset = int32(offset)
	limit, _ := strconv.Atoi(c.Request().FormValue("limit"))
	pageInfo.Limit = int32(limit)

	var totalCount int32
	apiResponse.ReportList, totalCount, err = h.uc.SelecAttrSearchReportList(pageInfo, attrData)
	// apiResponse.ReportList, totalCount, err = server.DBManager.DBGorm.SelecAttrSearchReportList(pageInfo, attrData)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.TotalCount.Count = totalCount
	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}

// 업무 보고 Detail 정보
func (h *BsmgHandler) GetReportInfoReq(c echo.Context) (err error) {
	log.Println("getReportInfoReq")

	var apiResponse define.BsmgReportInfoResponse

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	// 임시코드. 깔끔하게 수정하려면 get요청을 setParameter로 변경
	idxData := c.Request().FormValue("@d1#rpt_idx")
	rpt_idx, _ := strconv.Atoi(idxData)
	fmt.Println("rpt_idx : ", rpt_idx)

	apiResponse.ReportInfo, err = h.uc.SelectReportInfo(rpt_idx)
	// apiResponse.ReportInfo, err = server.DBManager.DBGorm.SelectReportInfo(rpt_idx)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}

// 업무보고에 들어있는 일정 정보
func (h *BsmgHandler) GetScheduleReq(c echo.Context) (err error) {
	log.Println("getScheduleReq")

	var apiResponse define.BsmgScheduleListResponse

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	stringIdx := c.Request().FormValue("@d1#rpt_idx")
	rptIdx, _ := strconv.Atoi(stringIdx)
	scheduleList, err := h.uc.SelectScheduleList(int32(rptIdx))
	// scheduleList, err := server.DBManager.DBGorm.SelectScheduleList(int32(rptIdx))
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}
	apiResponse.ScheduleList = scheduleList
	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}

// 주간보고 리스트 정보
func (h *BsmgHandler) GetWeekRptSearchReq(c echo.Context) (err error) {
	log.Println("getWeekRptSearchReq")

	var apiResponse define.BsmgWeekReportListResponse

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	pageInfo := define.PageInfo{}
	offset, _ := strconv.Atoi(c.Request().FormValue("offset"))
	limit, _ := strconv.Atoi(c.Request().FormValue("limit"))
	pageInfo.Offset, pageInfo.Limit = int32(offset), int32(limit)

	searchData := define.SearchData{}
	combo, _ := strconv.Atoi(c.Request().FormValue("@d1#search_combo"))
	input := c.Request().FormValue("@d1#search_input")
	searchData.SearchCombo, searchData.SearchInput = int32(combo), input

	var totalCount int32
	apiResponse.WeekReportList, totalCount, err = h.uc.SelectWeekReportList(pageInfo, searchData)
	// apiResponse.WeekReportList, totalCount, err = server.DBManager.DBGorm.SelectWeekReportList(pageInfo, searchData)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.TotalCount.Count = totalCount
	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}

// 주간 업무보고 카테고리 검색 ( = 부서로 검색)
func (h *BsmgHandler) GetWeekRptCategorySearch(c echo.Context) (err error) {
	log.Println("getWeekRptCategory")

	var apiResponse define.BsmgWeekReportListResponse

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	pageInfo := define.PageInfo{}
	offset, _ := strconv.Atoi(c.Request().FormValue("offset"))
	limit, _ := strconv.Atoi(c.Request().FormValue("limit"))
	pageInfo.Offset, pageInfo.Limit = int32(offset), int32(limit)

	partIdx, _ := strconv.Atoi(c.Request().FormValue("@d1#part_value"))

	var totalCount int32
	// DB 처리
	apiResponse.WeekReportList, totalCount, err = h.uc.SelectWeekReportCategorySearch(pageInfo, partIdx)
	// apiResponse.WeekReportList, totalCount, err = server.DBManager.DBGorm.SelectWeekReportCategorySearch(pageInfo, partIdx)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.TotalCount.Count = int32(totalCount)
	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}

// 주간 업무보고 Detail
func (h *BsmgHandler) GetWeekRptInfoReq(c echo.Context) (err error) {
	log.Println("getWeekRptInfoReq")

	var apiResponse define.BsmgWeekReportInfoResponse

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	wRptIdx, _ := strconv.Atoi(c.Request().FormValue("@d1#wRpt_idx"))
	// DB 처리
	apiResponse.WeekReportInfo, err = h.uc.SelectWeekReportInfo(wRptIdx)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}
	// apiResponse.WeekReportInfo, err = server.DBManager.DBGorm.SelectWeekReportInfo(wRptIdx)

	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}

// 업무보고 확인 (상급자 기능)
func (h *BsmgHandler) GetConfirmRptReq(c echo.Context) (err error) {
	log.Println("getConfirmRptReq")

	var apiResponse define.OnlyResult

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	rptIdx, _ := strconv.Atoi(c.Request().FormValue("@d1#rpt_idx"))

	// 권한 체크
	if _, resultCode := h.uc.AuthorizationCheck(c); resultCode != define.Success {
		apiResponse.Result.ResultCode = int32(resultCode)
		return c.JSON(http.StatusOK, apiResponse)
	}

	err = h.uc.ConfirmRpt(int32(rptIdx))
	// err = server.DBManager.DBGorm.ConfirmRpt(int32(rptIdx))
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}

func (h *BsmgHandler) PostReportReq(c echo.Context) error {
	log.Println("postReportReq")

	apiRequest := define.BsmgReportInfoRequest{}
	apiResponse := define.BsmgReportInfoResponse{}

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	// request data Get
	err := c.Bind(&apiRequest)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse = h.uc.PostReportReq(c, apiRequest)
	return c.JSON(http.StatusOK, apiResponse)

}

func (h *BsmgHandler) PostRegistScheduleReq(c echo.Context) (err error) {
	log.Println("postRegistScheduleReq")

	apiRequest := define.BsmgPostScheduleRequest{}
	apiResponse := define.OnlyResult{}

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	err = c.Bind(&apiRequest)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse = h.uc.PostRegistScheduleReq(c, apiRequest)
	return c.JSON(http.StatusOK, apiResponse)
}

func (h *BsmgHandler) PutReportReq(c echo.Context) (err error) {
	log.Println("putReportReq ")

	apiRequest := define.BsmgReportInfoRequest{}
	apiResponse := define.BsmgReportInfoResponse{}

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	err = c.Bind(&apiRequest)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}
	// parsing
	report := apiRequest.Data.BsmgReportInfo.ParseReport()

	apiResponse = h.uc.PutReportReq(c, report)
	return c.JSON(http.StatusOK, apiResponse)
}

func (h *BsmgHandler) PutScheduleReq(c echo.Context) (err error) {
	log.Println("putScheduleReq ")

	apiRequest := define.BsmgPutScheduleRequest{}
	apiResponse := define.OnlyResult{}

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	err = c.Bind(&apiRequest)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}

	idxData, _ := strconv.Atoi(apiRequest.Data.RptIdx.RptIdx)
	idx := int32(idxData)

	apiResponse = h.uc.PutScheduleReq(apiRequest, idx)
	return c.JSON(http.StatusOK, apiResponse)
}

func (h *BsmgHandler) DeleteReportReq(c echo.Context) (err error) {
	log.Println("deleteReportReq ")
	apiResponse := define.OnlyResult{}

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	rptIdxParam, _ := strconv.Atoi(c.Param("rptIdx"))
	rptIdx := int32(rptIdxParam)

	apiResponse = h.uc.DeleteReportReq(c, rptIdx)
	return c.JSON(http.StatusOK, apiResponse)
}

func (h *BsmgHandler) PutWeekRptReq(c echo.Context) (err error) {
	apiRequest := define.BsmgPutWeekReportRequest{}
	apiResponse := define.OnlyResult{}

	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	err = c.Bind(&apiRequest)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}

	// parsing
	report := apiRequest.Data.WeekReportInfo.ParseReport()

	apiResponse = h.uc.PutWeekRptReq(c, report)
	return c.JSON(http.StatusOK, apiResponse)
}

func (h *BsmgHandler) DeleteWeekRptReq(c echo.Context) (err error) {
	server, _ := c.Get("Server").(*server.ServerProcessor)
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	wRptIdx, _ := strconv.Atoi(c.Param("wRptIdx"))

	apiResponse := h.uc.DeleteWeekRptReq(c, wRptIdx)

	return c.JSON(http.StatusOK, apiResponse)
}
