package main

import (
	"BsmgRefactoring/define"
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// 업무보고 리스트 정보
func getReportListReq(c echo.Context) error {
	log.Println("getReportListReq")

	var result define.BsmgReportResult

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
		log.Printf("%v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}
	fmt.Printf("%v \n", RptList)

	var totalCount int = 0
	// DB 처리

	result.TotalCount.Count = int32(totalCount)
	result.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, result)
}

// 업무보고 리스트 검색 + 정보 이걸 왜 나눠놓은 거지?
func getReportSearchReq(c echo.Context) error {
	log.Println("getReportSearchReq")

	var result define.BsmgReportResult

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

	search := &define.SearchData{}
	search = parseSearchRequest(parser)

	fmt.Printf("%v \n", search)

	RptList, err := getPageInfo(c.Request().URL.RawQuery)
	if err != nil {
		log.Printf("%v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}
	fmt.Printf("%v \n", RptList)

	var totalCount int = 0
	// DB 처리

	result.TotalCount.Count = int32(totalCount)
	result.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, result)
}

// 업무 속성에 따른 보고 정보
func getReportAttrSearchReq(c echo.Context) error {
	log.Println("getReportAttrSearchReq")

	var result define.BsmgReportResult
	var attrValue int32
	var attrCategory int32

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
		log.Printf("%v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	attrValue, err = parser.getInt32Value(0, "attrValue", 0)
	if err != nil {
		log.Printf("%v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}
	attrCategory, err = parser.getInt32Value(0, "attrCategory", 0)
	if err != nil {
		log.Printf("%v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	fmt.Printf("%v \n", RptList)
	fmt.Printf("%v \n", attrValue)
	fmt.Printf("%v \n", attrCategory)

	var totalCount int = 0
	// DB 처리

	result.TotalCount.Count = int32(totalCount)
	result.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, result)
}

// 업무 보고 Detail 정보
func getReportInfoReq(c echo.Context) error {
	log.Println("getReportInfoReq")

	var result define.BsmgReportResult

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

	rpt_idx, err := parser.getInt32Value(0, "rpt_idx", 0)
	if err != nil {
		log.Printf("%v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	fmt.Printf("%v \n", rpt_idx)

	var totalCount int = 0
	// DB 처리

	result.TotalCount.Count = int32(totalCount)
	result.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, result)
}

// 업무보고에 들어있는 일정 정보
func getScheduleReq(c echo.Context) error {
	log.Println("getReportInfoReq")

	var result define.BsmgReportResult

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

	rpt_idx, err := parser.getInt32Value(0, "rpt_idx", 0)
	if err != nil {
		log.Printf("%v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	fmt.Printf("%v \n", rpt_idx)

	var totalCount int = 0
	// DB 처리

	result.TotalCount.Count = int32(totalCount)
	result.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, result)
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
func getWeekRptSearchReq(c echo.Context) error {
	log.Println("getWeekRptSearchReq")

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
	search := &define.SearchData{}
	search = parseSearchRequest(parser)

	fmt.Printf("%v \n", weekRptList)
	fmt.Printf("%v \n", search)

	var totalCount int = 0
	// DB 처리

	result.TotalCount.Count = int32(totalCount)
	result.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, result)
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
func getConfirmRptReq(c echo.Context) error {
	log.Println("getConfirmRptReq")

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

	rpt_idx, err := parser.getInt32Value(0, "rpt_idx", 0)
	if err != nil {
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}
	fmt.Printf("%v \n", rpt_idx)

	var totalCount int = 0
	// DB 처리

	result.TotalCount.Count = int32(totalCount)
	result.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, result)
}

func postReportReq(c echo.Context) error {
	log.Println("postReportReq")

	apiRequest := define.BsmgPostReportRequest{}
	apiResponse := define.OnlyResult{}

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
	err = server.dbManager.DBGorm.CreateDailyReport(report)
	if err != nil {
		apiResponse.Result.ResultCode = define.DataBaseError
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.Result.ResultCode = define.Success
	return c.JSON(http.StatusOK, apiResponse)

}
