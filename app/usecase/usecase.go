package usecase

import (
	"BsmgRefactoring/database"
	"BsmgRefactoring/define"
	"BsmgRefactoring/repository"

	"github.com/labstack/echo/v4"
)

// 애플리케이션/비즈니스 로직이 담겨있는 레이어

// 각 상황에 따라 이럴 땐 이렇게, 저럴 땐 저렇게 하는 등의 애플리케이션 내에서 어떤 가치를 만들어내는 자동화된 코드의 흐름이다.

type BsmgUsecase interface {
	CheckLoginIng(c echo.Context) (apiResponse define.BsmgMemberResponse, ResultCode int)
	PutUserReq(c echo.Context, apiRequest define.BsmgPutMemberRequest) (apiResponse define.OnlyResult)
	PostUserReq(c echo.Context, apiRequest define.BsmgMemberRequest) (apiResponse define.OnlyResult)
	GetIdCheckRequest(memID string) (apiResponse define.OnlyResult)
	GetUserListRequest() (result define.BsmgMemberListResponse)
	DeleteWeekRptReq(c echo.Context, wRptIdx int) (apiResponse define.OnlyResult)
	PutWeekRptReq(c echo.Context, report define.BsmgWeekRptInfo) (apiResponse define.OnlyResult)
	DeleteReportReq(c echo.Context, rptIdx int32) (apiRespone define.OnlyResult)
	PutScheduleReq(apiRequest define.BsmgPutScheduleRequest, idx int32) (apiResponse define.OnlyResult)
	PutReportReq(c echo.Context, report define.BsmgReportInfo) (apiResponse define.BsmgReportInfoResponse)
	PostRegistScheduleReq(c echo.Context, apiRequest define.BsmgPostScheduleRequest) (apiRespone define.OnlyResult)
	PostReportReq(c echo.Context, apiRequest define.BsmgReportInfoRequest) (apiResponse define.BsmgReportInfoResponse)
	SelectReportListReq(searchData define.SearchData, pageInfo define.PageInfo) (apiResponse define.BsmgReportListResponse)
	UserLogout(c echo.Context) (apiResponse define.OnlyResult)
	CheckUserLogin(c echo.Context, apiRequest define.BsmgMemberLoginRequest) (apiResponse define.BsmgMemberResponse)
	AuthorizationCheck(c echo.Context) (apiResponse define.BsmgMemberResponse, ResultCode int) // 권한 체크
	MakeWeekRpt()
	database.DBInterface
}

type structBsmgUsecase struct {
	rp repository.BsmgRepository
}

func NewBsmgUsecase(br repository.BsmgRepository) BsmgUsecase {
	return &structBsmgUsecase{br}
}
