package router

import (
	"BsmgRefactoring/handler"

	"github.com/labstack/echo/v4"
)

func InitRouteGroup(bsmgGroup *echo.Group, h handler.BsmgHandler) {
	bsmgLoginGroup := bsmgGroup.Group("/login")
	initLoginRoute(bsmgLoginGroup, h)

	bsmgUserGroup := bsmgGroup.Group("/user")
	initUserRoute(bsmgUserGroup, h)

	bsmgSettingGroup := bsmgGroup.Group("/setting")
	initSettingRoute(bsmgSettingGroup, h)

	bsmgReportGroup := bsmgGroup.Group("/report")
	initReportRoute(bsmgReportGroup, h)
}

// Login Group
func initLoginRoute(loginGroup *echo.Group, h handler.BsmgHandler) {
	// loginGroup.GET("/chkLogin", getChkLoginRequest)
	loginGroup.POST("/logout", h.PostLogoutRequest)
}

// User Group
func initUserRoute(userGroup *echo.Group, h handler.BsmgHandler) {
	userGroup.GET("/userList", h.GetUserListRequest)
	userGroup.GET("/idCheck", h.GetIdCheckRequest)
	userGroup.GET("/userSearch", h.GetUserSearchRequest)

	// 사용자 등록
	userGroup.POST("/", h.PostUserReq)

	// 사용자 정보 수정
	userGroup.PUT("", h.PutUserReq)

	// 사용자 삭제
	userGroup.DELETE("/deleteUser/:memID", h.DeleteUserReq)
}

// Setting Group
func initSettingRoute(settingGroup *echo.Group, h handler.BsmgHandler) {
	settingGroup.GET("/attrTree", h.GetAttrTreeReq)
	settingGroup.GET("/rankPart", h.GetRankPartReq)
	settingGroup.GET("/weekRptCategory", h.GetPartTree) // 주간보고 속성 트리 (부서 별로 볼 수 있는 기능)
	settingGroup.GET("/getToRpt", h.GetToRptReq)
	settingGroup.GET("/attr1", h.GetAttr1Req) // 업무 카테고리만 return
}

// Report Group
func initReportRoute(reportGroup *echo.Group, h handler.BsmgHandler) {
	reportGroup.GET("/reportList", h.GetReportSearchReq)
	reportGroup.GET("/reportSearch", h.GetReportSearchReq)
	reportGroup.GET("/reportAttrSearch", h.GetReportAttrSearchReq)
	reportGroup.GET("/reportInfo", h.GetReportInfoReq)
	reportGroup.GET("/getSchdule", h.GetScheduleReq)
	reportGroup.GET("/getWeekRptList", h.GetWeekRptSearchReq)
	reportGroup.GET("/getWeekRptSearch", h.GetWeekRptSearchReq)
	reportGroup.GET("/getWeekRptCategory", h.GetWeekRptCategorySearch)
	reportGroup.GET("/getWeekRptInfo", h.GetWeekRptInfoReq)
	reportGroup.GET("/confirmRpt", h.GetConfirmRptReq)

	// POST
	reportGroup.POST("/report", h.PostReportReq)
	reportGroup.POST("/registSch", h.PostRegistScheduleReq)

	// PUT
	reportGroup.PUT("/putRpt", h.PutReportReq)
	reportGroup.PUT("/putSchedule", h.PutScheduleReq)
	reportGroup.PUT("/putWeekRpt", h.PutWeekRptReq)

	// DELETE
	reportGroup.DELETE("/deleteRpt/:rptIdx", h.DeleteReportReq)
	reportGroup.DELETE("/deleteWeekRpt/:wRptIdx", h.DeleteWeekRptReq)
}
