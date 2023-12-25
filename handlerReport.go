package main

import (
	"BsmgRefactoring/define"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// server를 지금 Main에서 전역으로 선언하고 그걸 갖다 쓰는데 이게 맞나?
// 아니면 echo의 context에서 꺼내서 써야하나??????

// 업무보고 리스트 정보
func getReportListReq(c echo.Context) (err error) {
	log.Println("getReportListReq")

	var apiResponse define.BsmgReportListResponse

	pageInfo := define.PageInfo{}
	offset, _ := strconv.Atoi(c.Request().FormValue("offset"))
	pageInfo.Offset = int32(offset)
	limit, _ := strconv.Atoi(c.Request().FormValue("limit"))
	pageInfo.Limit = int32(limit)

	searchData := define.SearchData{}

	var totalCount int32 // 페이징 처리
	apiResponse.ReportList, totalCount, err = server.dbManager.DBGorm.SelectReportList(pageInfo, searchData)
	if err != nil {
		log.Printf("%v \n", err)
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.TotalCount.Count = totalCount
	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}

func getReportSearchReq(c echo.Context) (err error) {
	log.Println("getReportSearchReq")

	apiResponse := define.BsmgReportListResponse{}

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

	var totalCount int32 // 페이징처리
	apiResponse.ReportList, totalCount, err = server.dbManager.DBGorm.SelectReportList(pageInfo, searchData)

	if err != nil {
		log.Printf("%v \n", err)
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.TotalCount.Count = totalCount
	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}

// 업무 속성에 따른 보고 정보
func getReportAttrSearchReq(c echo.Context) (err error) {
	log.Println("getReportAttrSearchReq")

	var apiResponse define.BsmgReportListResponse

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
	apiResponse.ReportList, totalCount, err = server.dbManager.DBGorm.SelecAttrSearchReportList(pageInfo, attrData)
	if err != nil {
		log.Printf("%v \n", err)
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.TotalCount.Count = totalCount
	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}

// 업무 보고 Detail 정보
func getReportInfoReq(c echo.Context) (err error) {
	log.Println("getReportInfoReq")

	var apiResponse define.BsmgReportInfoResponse

	// 임시코드. 깔끔하게 수정하려면 get요청을 setParameter로 변경
	idxData := c.Request().FormValue("@d1#rpt_idx")
	rpt_idx, _ := strconv.Atoi(idxData)
	fmt.Println("rpt_idx : ", rpt_idx)

	apiResponse.ReportInfo, err = server.dbManager.DBGorm.SelectReportInfo(rpt_idx)
	if err != nil {
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}

// 업무보고에 들어있는 일정 정보
func getScheduleReq(c echo.Context) (err error) {
	log.Println("getScheduleReq")

	var apiRespone define.BsmgScheduleListResponse

	stringIdx := c.Request().FormValue("@d1#rpt_idx")
	rptIdx, _ := strconv.Atoi(stringIdx)
	scheduleList, err := server.dbManager.DBGorm.SelectScheduleList(int32(rptIdx))
	if err != nil {
		log.Printf("%v \n", err)
		apiRespone.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiRespone)
	}
	apiRespone.ScheduleList = scheduleList
	apiRespone.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiRespone)
}

// 주간보고 리스트 정보
func getWeekRptListReq(c echo.Context) error {
	log.Println("getWeekRptListReq")

	var result define.BsmgWeekRptResult

	value, err := c.FormParams()
	if err != nil {
		log.Printf("%v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	parser := initFormParser(value)
	if parser == nil {
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	weekRptList, err := getPageInfo(c.Request().URL.RawQuery)
	if err != nil {
		log.Printf("페이징처리 오류 %v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	fmt.Printf("%v \n", weekRptList)

	var totalCount int = 0
	// DB 처리

	result.TotalCount.Count = int32(totalCount)
	result.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, result)
}

// 주간보고 리스트 정보
func getWeekRptSearchReq(c echo.Context) (err error) {
	log.Println("getWeekRptSearchReq")

	var apiResponse define.BsmgWeekReportListResponse

	pageInfo := define.PageInfo{}
	offset, _ := strconv.Atoi(c.Request().FormValue("offset"))
	limit, _ := strconv.Atoi(c.Request().FormValue("limit"))
	pageInfo.Offset, pageInfo.Limit = int32(offset), int32(limit)

	searchData := &define.SearchData{}
	combo, _ := strconv.Atoi(c.Request().FormValue("@d1#search_combo"))
	input := c.Request().FormValue("@d1#search_input")
	searchData.SearchCombo, searchData.SearchInput = int32(combo), input

	var totalCount int32
	// apiResponse.WeekReportList, totalCount, err = server.dbManager.DBGorm.SelectWeekReportList(pageInfo, searchData)
	if err != nil {
		log.Printf("getWeekRptSearchReq: %v \n", err)
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.TotalCount.Count = totalCount
	apiResponse.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiResponse)
}

// 주간 업무보고 카테고리 정보
func getWeekRptCategory(c echo.Context) error {
	log.Println("getWeekRptCategory")

	var result define.BsmgWeekRptResult

	value, err := c.FormParams()
	if err != nil {
		log.Printf("%v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	parser := initFormParser(value)
	if parser == nil {
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	RptList, err := getPageInfo(c.Request().URL.RawQuery)
	if err != nil {
		log.Printf("페이징처리 오류 %v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	partValue, err := parser.getInt32Value(0, "part_value", 0)
	if err != nil {
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	fmt.Printf("%v \n", RptList)
	fmt.Printf("%v \n", partValue)

	var totalCount int = 0
	// DB 처리

	result.TotalCount.Count = int32(totalCount)
	result.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, result)
}

// 주간 업무보고 Detail
func getWeekRptInfoReq(c echo.Context) error {
	log.Println("getWeekRptInfoReq")

	var result define.BsmgWeekRptResult

	value, err := c.FormParams()
	if err != nil {
		log.Printf("%v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	parser := initFormParser(value)
	if parser == nil {
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	wRpt_idx, err := parser.getInt32Value(0, "wRpt_idx", 0)
	if err != nil {
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}
	fmt.Printf("%v \n", wRpt_idx)

	var totalCount int = 0
	// DB 처리

	result.TotalCount.Count = int32(totalCount)
	result.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, result)
}

// 업무보고 확인 (상급자 기능)
func getConfirmRptReq(c echo.Context) (err error) {
	log.Println("getConfirmRptReq")

	var apiRespone define.OnlyResult

	idxParam := c.Request().FormValue("@d1#rpt_idx")
	rptIdx, _ := strconv.Atoi(idxParam)

	// 서버에서도 확인작업 하면 좋겠지만.. 일단 웹에서 권한에 대해 확인했으니 패스

	err = server.dbManager.DBGorm.ConfirmRpt(int32(rptIdx))
	if err != nil {
		apiRespone.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiRespone)
	}

	apiRespone.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, apiRespone)
}

func postReportReq(c echo.Context) error {
	log.Println("postReportReq")

	apiRequest := define.BsmgReportInfoRequest{}
	apiResponse := define.BsmgReportInfoResponse{}

	// 세션으로 클라이언트 정보 Get
	session, err := session.Get(sessionKey, c)
	if err != nil {
		apiResponse.Result.ResultCode = define.ErrorSession
		return c.JSON(http.StatusOK, apiResponse)
	}
	client := session.Values["Member"].(define.BsmgMemberInfo)

	// request data Get
	err = c.Bind(&apiRequest)
	if err != nil {
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}

	// parsing
	report := apiRequest.Data.BsmgReportInfo.ParseReport()
	report.Rpt_Reporter = client.Mem_ID

	// DB 처리
	server := c.Get("Server").(*ServerProcessor)
	err = server.dbManager.DBGorm.InsertDailyReport(report)
	if err != nil {
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	// 스케쥴 등록을 위한 idx 반환
	idx, err := server.dbManager.DBGorm.SelectLatestRptIdx(report.Rpt_Reporter)
	if err != nil {
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}
	report.Rpt_Idx = idx

	apiResponse.ReportInfo = report
	apiResponse.Result.ResultCode = define.Success
	return c.JSON(http.StatusOK, apiResponse)

}

func postRegistScheduleReq(c echo.Context) (err error) {
	log.Println("postRegistScheduleReq")

	apiRequest := define.BsmgPostScheduleRequest{}
	apiRespone := define.OnlyResult{}

	err = c.Bind(&apiRequest)
	if err != nil {
		apiRespone.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiRespone)
	}

	idx, _ := strconv.Atoi(apiRequest.Data.BsmgReportInfo.Rpt_idx)
	for _, scheduleString := range apiRequest.Data.BsmgScheduleInfo {
		schedule := scheduleString.ParseSchedule()

		schedule.Rpt_Idx = int32(idx)
		err = server.dbManager.DBGorm.InsertSchedule(schedule)
		if err != nil {
			apiRespone.Result.ResultCode = define.ErrorDataBase
			return c.JSON(http.StatusOK, apiRespone)
		}
	}

	apiRespone.Result.ResultCode = define.Success
	return c.JSON(http.StatusOK, apiRespone)
}

func putReportReq(c echo.Context) (err error) {
	log.Println("putReportReq ")

	apiRequest := define.BsmgReportInfoRequest{}
	apiResponse := define.BsmgReportInfoResponse{}
	err = c.Bind(&apiRequest)
	if err != nil {
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}
	// parsing
	report := apiRequest.Data.BsmgReportInfo.ParseReport()

	// 세션으로 클라이언트 정보 Get
	session, err := session.Get(sessionKey, c)
	if err != nil {
		apiResponse.Result.ResultCode = define.ErrorSession
		return c.JSON(http.StatusOK, apiResponse)
	}
	client := session.Values["Member"].(define.BsmgMemberInfo)

	// 본인만 수정 가능
	if client.Mem_ID != report.Rpt_Reporter {
		apiResponse.Result.ResultCode = define.ErrorNotAuthorizedUser
		return c.JSON(http.StatusOK, apiResponse)
	}

	err = server.dbManager.DBGorm.UpdateReportInfo(report)
	if err != nil {
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.Result.ResultCode = define.Success
	return c.JSON(http.StatusOK, apiResponse)
}

func putScheduleReq(c echo.Context) (err error) {
	log.Println("putScheduleReq ")

	apiRequest := define.BsmgPutScheduleRequest{}
	apiResponse := define.OnlyResult{}

	err = c.Bind(&apiRequest)
	if err != nil {
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, apiResponse)
	}

	idxData, _ := strconv.Atoi(apiRequest.Data.RptIdx.RptIdx)
	idx := int32(idxData)

	// (무엇이 바뀌었는지 특정할 수 없으므로 전부 삭제 후 재 삽입)
	// 기존 스케쥴 삭제
	err = server.dbManager.DBGorm.DeleteSchedule(idx)
	if err != nil {
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	for _, scheduleString := range apiRequest.Data.BsmgScheduleInfo {
		schedule := scheduleString.ParseSchedule()
		schedule.Rpt_Idx = idx

		// 신규 스케쥴 삽입
		err = server.dbManager.DBGorm.InsertSchedule(schedule)
		if err != nil {
			apiResponse.Result.ResultCode = define.ErrorDataBase
			return c.JSON(http.StatusOK, apiResponse)
		}
	}

	apiResponse.Result.ResultCode = define.Success
	return c.JSON(http.StatusOK, apiResponse)
}

func deleteReportReq(c echo.Context) (err error) {
	log.Println("deleteReportReq ")
	apiRespone := define.OnlyResult{}

	rptIdxParam, _ := strconv.Atoi(c.Param("rptIdx"))
	rptIdx := int32(rptIdxParam)

	err = server.dbManager.DBGorm.DeleteSchedule(rptIdx)
	if err != nil {
		apiRespone.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiRespone)
	}

	err = server.dbManager.DBGorm.DeleteReport(rptIdx)
	if err != nil {
		apiRespone.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiRespone)
	}

	apiRespone.Result.ResultCode = define.Success
	return c.JSON(http.StatusOK, apiRespone)
}
