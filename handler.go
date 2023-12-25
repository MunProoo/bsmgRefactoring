package main

import "github.com/labstack/echo/v4"

func initRouteGroup(bsmgGroup *echo.Group) {
	bsmgLoginGroup := bsmgGroup.Group("/login")
	initLoginRoute(bsmgLoginGroup)

	bsmgUserGroup := bsmgGroup.Group("/user")
	initUserRoute(bsmgUserGroup)

	bsmgSettingGroup := bsmgGroup.Group("/setting")
	initSettingRoute(bsmgSettingGroup)

	bsmgReportGroup := bsmgGroup.Group("/report")
	initReportRoute(bsmgReportGroup)
}

// Login Group
func initLoginRoute(loginGroup *echo.Group) {
	loginGroup.GET("/chkLogin", getChkLoginRequest)
	loginGroup.POST("/login", postLoginRequest)
	loginGroup.POST("/logout", postLogoutRequest)
}

// User Group
func initUserRoute(userGroup *echo.Group) {
	userGroup.GET("/userList", getUserListRequest)
	userGroup.GET("/idCheck", getIdCheckRequest)
	userGroup.GET("/userSearch", getUserSearchRequest)

	// 사용자 등록
	userGroup.POST("/", postUserReq)

	// 사용자 정보 수정
	userGroup.PUT("", putUserReq)

	// 사용자 삭제
	userGroup.DELETE("/deleteUser/:memID", deleteUserReq)
}

// Setting Group
func initSettingRoute(settingGroup *echo.Group) {
	settingGroup.GET("/attrTree", getAttrTreeReq)
	settingGroup.GET("/rankPart", getRankPartReq)
	settingGroup.GET("/weekRptCategory", getWeekRptCategoryReq)
	settingGroup.GET("/getToRpt", getToRptReq)
}

// Report Group
func initReportRoute(reportGroup *echo.Group) {
	reportGroup.GET("/reportList", getReportSearchReq)
	reportGroup.GET("/reportSearch", getReportSearchReq)
	reportGroup.GET("/reportAttrSearch", getReportAttrSearchReq)
	reportGroup.GET("/reportInfo", getReportInfoReq)
	reportGroup.GET("/getSchdule", getScheduleReq)
	reportGroup.GET("/getWeekRptList", getWeekRptListReq)
	reportGroup.GET("/getWeekRptSearch", getWeekRptSearchReq)
	reportGroup.GET("/getWeekRptCategory", getWeekRptCategory)
	reportGroup.GET("/getWeekRptInfo", getWeekRptInfoReq)
	reportGroup.GET("/confirmRpt", getConfirmRptReq)

	// POST
	reportGroup.POST("/report", postReportReq)
	reportGroup.POST("/registSch", postRegistScheduleReq)

	// PUT
	reportGroup.PUT("/putRpt", putReportReq)
	reportGroup.PUT("/putSchedule", putScheduleReq)
	// reportGroup.PUT("/putWeekRpt", putWeekRptReq)

	// DELETE
	reportGroup.DELETE("/deleteRpt/:rptIdx", deleteReportReq)
	// reportGroup.DELETE("/deleteWeekRpt", deleteWeekRptReq)
}
