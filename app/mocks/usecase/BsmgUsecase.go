// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	define "BsmgRefactoring/define"

	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// BsmgUsecase is an autogenerated mock type for the BsmgUsecase type
type BsmgUsecase struct {
	mock.Mock
}

// AuthorizationCheck provides a mock function with given fields: c
func (_m *BsmgUsecase) AuthorizationCheck(c echo.Context) (define.BsmgMemberResponse, int) {
	ret := _m.Called(c)

	var r0 define.BsmgMemberResponse
	if rf, ok := ret.Get(0).(func(echo.Context) define.BsmgMemberResponse); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(define.BsmgMemberResponse)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(echo.Context) int); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// CheckLoginIng provides a mock function with given fields: c
func (_m *BsmgUsecase) CheckLoginIng(c echo.Context) (define.BsmgMemberResponse, int) {
	ret := _m.Called(c)

	var r0 define.BsmgMemberResponse
	if rf, ok := ret.Get(0).(func(echo.Context) define.BsmgMemberResponse); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(define.BsmgMemberResponse)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(echo.Context) int); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// CheckMemberIDDuplicate provides a mock function with given fields: memID
func (_m *BsmgUsecase) CheckMemberIDDuplicate(memID string) (bool, error) {
	ret := _m.Called(memID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(memID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(memID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckUserLogin provides a mock function with given fields: c, apiRequest
func (_m *BsmgUsecase) CheckUserLogin(c echo.Context, apiRequest define.BsmgMemberLoginRequest) define.BsmgMemberResponse {
	ret := _m.Called(c, apiRequest)

	var r0 define.BsmgMemberResponse
	if rf, ok := ret.Get(0).(func(echo.Context, define.BsmgMemberLoginRequest) define.BsmgMemberResponse); ok {
		r0 = rf(c, apiRequest)
	} else {
		r0 = ret.Get(0).(define.BsmgMemberResponse)
	}

	return r0
}

// ConfirmRpt provides a mock function with given fields: rptIdx
func (_m *BsmgUsecase) ConfirmRpt(rptIdx int32) error {
	ret := _m.Called(rptIdx)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32) error); ok {
		r0 = rf(rptIdx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConnectBSMG provides a mock function with given fields:
func (_m *BsmgUsecase) ConnectBSMG() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConnectMariaDB provides a mock function with given fields:
func (_m *BsmgUsecase) ConnectMariaDB() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateAttr1Table provides a mock function with given fields:
func (_m *BsmgUsecase) CreateAttr1Table() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateAttr2Table provides a mock function with given fields:
func (_m *BsmgUsecase) CreateAttr2Table() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateDailyReportTable provides a mock function with given fields:
func (_m *BsmgUsecase) CreateDailyReportTable() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateDataBase provides a mock function with given fields:
func (_m *BsmgUsecase) CreateDataBase() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMemberTable provides a mock function with given fields:
func (_m *BsmgUsecase) CreateMemberTable() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatePartTable provides a mock function with given fields:
func (_m *BsmgUsecase) CreatePartTable() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateRankTable provides a mock function with given fields:
func (_m *BsmgUsecase) CreateRankTable() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateScheduleTable provides a mock function with given fields:
func (_m *BsmgUsecase) CreateScheduleTable() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateWeekReportTable provides a mock function with given fields:
func (_m *BsmgUsecase) CreateWeekReportTable() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMember provides a mock function with given fields: memId
func (_m *BsmgUsecase) DeleteMember(memId string) error {
	ret := _m.Called(memId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(memId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReport provides a mock function with given fields: rptIdx
func (_m *BsmgUsecase) DeleteReport(rptIdx int32) error {
	ret := _m.Called(rptIdx)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32) error); ok {
		r0 = rf(rptIdx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReportReq provides a mock function with given fields: c, rptIdx
func (_m *BsmgUsecase) DeleteReportReq(c echo.Context, rptIdx int32) define.OnlyResult {
	ret := _m.Called(c, rptIdx)

	var r0 define.OnlyResult
	if rf, ok := ret.Get(0).(func(echo.Context, int32) define.OnlyResult); ok {
		r0 = rf(c, rptIdx)
	} else {
		r0 = ret.Get(0).(define.OnlyResult)
	}

	return r0
}

// DeleteSchedule provides a mock function with given fields: rptIdx
func (_m *BsmgUsecase) DeleteSchedule(rptIdx int32) error {
	ret := _m.Called(rptIdx)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32) error); ok {
		r0 = rf(rptIdx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteWeekReport provides a mock function with given fields: wRptIdx
func (_m *BsmgUsecase) DeleteWeekReport(wRptIdx int) error {
	ret := _m.Called(wRptIdx)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(wRptIdx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteWeekRptReq provides a mock function with given fields: c, wRptIdx
func (_m *BsmgUsecase) DeleteWeekRptReq(c echo.Context, wRptIdx int) define.OnlyResult {
	ret := _m.Called(c, wRptIdx)

	var r0 define.OnlyResult
	if rf, ok := ret.Get(0).(func(echo.Context, int) define.OnlyResult); ok {
		r0 = rf(c, wRptIdx)
	} else {
		r0 = ret.Get(0).(define.OnlyResult)
	}

	return r0
}

// FindMinIdx provides a mock function with given fields:
func (_m *BsmgUsecase) FindMinIdx() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// GetIdCheckRequest provides a mock function with given fields: memID
func (_m *BsmgUsecase) GetIdCheckRequest(memID string) define.OnlyResult {
	ret := _m.Called(memID)

	var r0 define.OnlyResult
	if rf, ok := ret.Get(0).(func(string) define.OnlyResult); ok {
		r0 = rf(memID)
	} else {
		r0 = ret.Get(0).(define.OnlyResult)
	}

	return r0
}

// GetUserListRequest provides a mock function with given fields:
func (_m *BsmgUsecase) GetUserListRequest() define.BsmgMemberListResponse {
	ret := _m.Called()

	var r0 define.BsmgMemberListResponse
	if rf, ok := ret.Get(0).(func() define.BsmgMemberListResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(define.BsmgMemberListResponse)
	}

	return r0
}

// InsertDailyReport provides a mock function with given fields: report
func (_m *BsmgUsecase) InsertDailyReport(report define.BsmgReportInfo) error {
	ret := _m.Called(report)

	var r0 error
	if rf, ok := ret.Get(0).(func(define.BsmgReportInfo) error); ok {
		r0 = rf(report)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertDefaultAttr1 provides a mock function with given fields:
func (_m *BsmgUsecase) InsertDefaultAttr1() {
	_m.Called()
}

// InsertDefaultAttr2 provides a mock function with given fields:
func (_m *BsmgUsecase) InsertDefaultAttr2() {
	_m.Called()
}

// InsertDefaultPart provides a mock function with given fields:
func (_m *BsmgUsecase) InsertDefaultPart() {
	_m.Called()
}

// InsertDefaultRank provides a mock function with given fields:
func (_m *BsmgUsecase) InsertDefaultRank() {
	_m.Called()
}

// InsertMember provides a mock function with given fields: member
func (_m *BsmgUsecase) InsertMember(member define.BsmgMemberInfo) error {
	ret := _m.Called(member)

	var r0 error
	if rf, ok := ret.Get(0).(func(define.BsmgMemberInfo) error); ok {
		r0 = rf(member)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertSchedule provides a mock function with given fields: schedule
func (_m *BsmgUsecase) InsertSchedule(schedule define.BsmgScheduleInfo) error {
	ret := _m.Called(schedule)

	var r0 error
	if rf, ok := ret.Get(0).(func(define.BsmgScheduleInfo) error); ok {
		r0 = rf(schedule)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertWeekReport provides a mock function with given fields: wRptInfo
func (_m *BsmgUsecase) InsertWeekReport(wRptInfo define.BsmgWeekRptInfo) error {
	ret := _m.Called(wRptInfo)

	var r0 error
	if rf, ok := ret.Get(0).(func(define.BsmgWeekRptInfo) error); ok {
		r0 = rf(wRptInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsConnected provides a mock function with given fields:
func (_m *BsmgUsecase) IsConnected() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsExistBSMG provides a mock function with given fields:
func (_m *BsmgUsecase) IsExistBSMG() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MakeAttrTree provides a mock function with given fields:
func (_m *BsmgUsecase) MakeAttrTree() ([]define.AttrTree, error) {
	ret := _m.Called()

	var r0 []define.AttrTree
	if rf, ok := ret.Get(0).(func() []define.AttrTree); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]define.AttrTree)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakePartTree provides a mock function with given fields:
func (_m *BsmgUsecase) MakePartTree() ([]define.PartTree, error) {
	ret := _m.Called()

	var r0 []define.PartTree
	if rf, ok := ret.Get(0).(func() []define.PartTree); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]define.PartTree)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeWeekRpt provides a mock function with given fields:
func (_m *BsmgUsecase) MakeWeekRpt() {
	_m.Called()
}

// PostRegistScheduleReq provides a mock function with given fields: c, apiRequest
func (_m *BsmgUsecase) PostRegistScheduleReq(c echo.Context, apiRequest define.BsmgPostScheduleRequest) define.OnlyResult {
	ret := _m.Called(c, apiRequest)

	var r0 define.OnlyResult
	if rf, ok := ret.Get(0).(func(echo.Context, define.BsmgPostScheduleRequest) define.OnlyResult); ok {
		r0 = rf(c, apiRequest)
	} else {
		r0 = ret.Get(0).(define.OnlyResult)
	}

	return r0
}

// PostReportReq provides a mock function with given fields: c, apiRequest
func (_m *BsmgUsecase) PostReportReq(c echo.Context, apiRequest define.BsmgReportInfoRequest) define.BsmgReportInfoResponse {
	ret := _m.Called(c, apiRequest)

	var r0 define.BsmgReportInfoResponse
	if rf, ok := ret.Get(0).(func(echo.Context, define.BsmgReportInfoRequest) define.BsmgReportInfoResponse); ok {
		r0 = rf(c, apiRequest)
	} else {
		r0 = ret.Get(0).(define.BsmgReportInfoResponse)
	}

	return r0
}

// PostUserReq provides a mock function with given fields: c, apiRequest
func (_m *BsmgUsecase) PostUserReq(c echo.Context, apiRequest define.BsmgMemberRequest) define.OnlyResult {
	ret := _m.Called(c, apiRequest)

	var r0 define.OnlyResult
	if rf, ok := ret.Get(0).(func(echo.Context, define.BsmgMemberRequest) define.OnlyResult); ok {
		r0 = rf(c, apiRequest)
	} else {
		r0 = ret.Get(0).(define.OnlyResult)
	}

	return r0
}

// PutReportReq provides a mock function with given fields: c, report
func (_m *BsmgUsecase) PutReportReq(c echo.Context, report define.BsmgReportInfo) define.BsmgReportInfoResponse {
	ret := _m.Called(c, report)

	var r0 define.BsmgReportInfoResponse
	if rf, ok := ret.Get(0).(func(echo.Context, define.BsmgReportInfo) define.BsmgReportInfoResponse); ok {
		r0 = rf(c, report)
	} else {
		r0 = ret.Get(0).(define.BsmgReportInfoResponse)
	}

	return r0
}

// PutScheduleReq provides a mock function with given fields: apiRequest, idx
func (_m *BsmgUsecase) PutScheduleReq(apiRequest define.BsmgPutScheduleRequest, idx int32) define.OnlyResult {
	ret := _m.Called(apiRequest, idx)

	var r0 define.OnlyResult
	if rf, ok := ret.Get(0).(func(define.BsmgPutScheduleRequest, int32) define.OnlyResult); ok {
		r0 = rf(apiRequest, idx)
	} else {
		r0 = ret.Get(0).(define.OnlyResult)
	}

	return r0
}

// PutUserReq provides a mock function with given fields: c, apiRequest
func (_m *BsmgUsecase) PutUserReq(c echo.Context, apiRequest define.BsmgPutMemberRequest) define.OnlyResult {
	ret := _m.Called(c, apiRequest)

	var r0 define.OnlyResult
	if rf, ok := ret.Get(0).(func(echo.Context, define.BsmgPutMemberRequest) define.OnlyResult); ok {
		r0 = rf(c, apiRequest)
	} else {
		r0 = ret.Get(0).(define.OnlyResult)
	}

	return r0
}

// PutWeekRptReq provides a mock function with given fields: c, report
func (_m *BsmgUsecase) PutWeekRptReq(c echo.Context, report define.BsmgWeekRptInfo) define.OnlyResult {
	ret := _m.Called(c, report)

	var r0 define.OnlyResult
	if rf, ok := ret.Get(0).(func(echo.Context, define.BsmgWeekRptInfo) define.OnlyResult); ok {
		r0 = rf(c, report)
	} else {
		r0 = ret.Get(0).(define.OnlyResult)
	}

	return r0
}

// Release provides a mock function with given fields:
func (_m *BsmgUsecase) Release() {
	_m.Called()
}

// SelecAttrSearchReportList provides a mock function with given fields: pageInfo, attrData
func (_m *BsmgUsecase) SelecAttrSearchReportList(pageInfo define.PageInfo, attrData define.AttrSearchData) ([]define.BsmgReportInfoForWeb, int32, error) {
	ret := _m.Called(pageInfo, attrData)

	var r0 []define.BsmgReportInfoForWeb
	if rf, ok := ret.Get(0).(func(define.PageInfo, define.AttrSearchData) []define.BsmgReportInfoForWeb); ok {
		r0 = rf(pageInfo, attrData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]define.BsmgReportInfoForWeb)
		}
	}

	var r1 int32
	if rf, ok := ret.Get(1).(func(define.PageInfo, define.AttrSearchData) int32); ok {
		r1 = rf(pageInfo, attrData)
	} else {
		r1 = ret.Get(1).(int32)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(define.PageInfo, define.AttrSearchData) error); ok {
		r2 = rf(pageInfo, attrData)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SelectAttr1List provides a mock function with given fields:
func (_m *BsmgUsecase) SelectAttr1List() ([]define.BsmgAttr1Info, error) {
	ret := _m.Called()

	var r0 []define.BsmgAttr1Info
	if rf, ok := ret.Get(0).(func() []define.BsmgAttr1Info); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]define.BsmgAttr1Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectLatestRptIdx provides a mock function with given fields: reporter
func (_m *BsmgUsecase) SelectLatestRptIdx(reporter string) (int32, error) {
	ret := _m.Called(reporter)

	var r0 int32
	if rf, ok := ret.Get(0).(func(string) int32); ok {
		r0 = rf(reporter)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(reporter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectMemberInfo provides a mock function with given fields: member
func (_m *BsmgUsecase) SelectMemberInfo(member *define.BsmgMemberInfo) error {
	ret := _m.Called(member)

	var r0 error
	if rf, ok := ret.Get(0).(func(*define.BsmgMemberInfo) error); ok {
		r0 = rf(member)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectMemberListSearch provides a mock function with given fields: searchData
func (_m *BsmgUsecase) SelectMemberListSearch(searchData define.SearchData) ([]define.BsmgMemberInfo, error) {
	ret := _m.Called(searchData)

	var r0 []define.BsmgMemberInfo
	if rf, ok := ret.Get(0).(func(define.SearchData) []define.BsmgMemberInfo); ok {
		r0 = rf(searchData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]define.BsmgMemberInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(define.SearchData) error); ok {
		r1 = rf(searchData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectPartLeader provides a mock function with given fields: Mem_Part
func (_m *BsmgUsecase) SelectPartLeader(Mem_Part int32) (string, error) {
	ret := _m.Called(Mem_Part)

	var r0 string
	if rf, ok := ret.Get(0).(func(int32) string); ok {
		r0 = rf(Mem_Part)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(Mem_Part)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectPartist provides a mock function with given fields:
func (_m *BsmgUsecase) SelectPartist() ([]define.BsmgPartInfo, error) {
	ret := _m.Called()

	var r0 []define.BsmgPartInfo
	if rf, ok := ret.Get(0).(func() []define.BsmgPartInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]define.BsmgPartInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectRankList provides a mock function with given fields:
func (_m *BsmgUsecase) SelectRankList() ([]define.BsmgRankInfo, error) {
	ret := _m.Called()

	var r0 []define.BsmgRankInfo
	if rf, ok := ret.Get(0).(func() []define.BsmgRankInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]define.BsmgRankInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectReportInfo provides a mock function with given fields: idx
func (_m *BsmgUsecase) SelectReportInfo(idx int) (define.BsmgReportInfoForWeb, error) {
	ret := _m.Called(idx)

	var r0 define.BsmgReportInfoForWeb
	if rf, ok := ret.Get(0).(func(int) define.BsmgReportInfoForWeb); ok {
		r0 = rf(idx)
	} else {
		r0 = ret.Get(0).(define.BsmgReportInfoForWeb)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectReportList provides a mock function with given fields: pageInfo, searchData
func (_m *BsmgUsecase) SelectReportList(pageInfo define.PageInfo, searchData define.SearchData) ([]define.BsmgReportInfoForWeb, int32, error) {
	ret := _m.Called(pageInfo, searchData)

	var r0 []define.BsmgReportInfoForWeb
	if rf, ok := ret.Get(0).(func(define.PageInfo, define.SearchData) []define.BsmgReportInfoForWeb); ok {
		r0 = rf(pageInfo, searchData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]define.BsmgReportInfoForWeb)
		}
	}

	var r1 int32
	if rf, ok := ret.Get(1).(func(define.PageInfo, define.SearchData) int32); ok {
		r1 = rf(pageInfo, searchData)
	} else {
		r1 = ret.Get(1).(int32)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(define.PageInfo, define.SearchData) error); ok {
		r2 = rf(pageInfo, searchData)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SelectReportListAWeek provides a mock function with given fields: Mem_ID, bef7d, bef1d
func (_m *BsmgUsecase) SelectReportListAWeek(Mem_ID string, bef7d string, bef1d string) ([]define.BsmgReportInfo, error) {
	ret := _m.Called(Mem_ID, bef7d, bef1d)

	var r0 []define.BsmgReportInfo
	if rf, ok := ret.Get(0).(func(string, string, string) []define.BsmgReportInfo); ok {
		r0 = rf(Mem_ID, bef7d, bef1d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]define.BsmgReportInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(Mem_ID, bef7d, bef1d)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectReportListReq provides a mock function with given fields: searchData, pageInfo
func (_m *BsmgUsecase) SelectReportListReq(searchData define.SearchData, pageInfo define.PageInfo) define.BsmgReportListResponse {
	ret := _m.Called(searchData, pageInfo)

	var r0 define.BsmgReportListResponse
	if rf, ok := ret.Get(0).(func(define.SearchData, define.PageInfo) define.BsmgReportListResponse); ok {
		r0 = rf(searchData, pageInfo)
	} else {
		r0 = ret.Get(0).(define.BsmgReportListResponse)
	}

	return r0
}

// SelectScheduleList provides a mock function with given fields: rptIdx
func (_m *BsmgUsecase) SelectScheduleList(rptIdx int32) ([]define.BsmgScheduleInfo, error) {
	ret := _m.Called(rptIdx)

	var r0 []define.BsmgScheduleInfo
	if rf, ok := ret.Get(0).(func(int32) []define.BsmgScheduleInfo); ok {
		r0 = rf(rptIdx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]define.BsmgScheduleInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(rptIdx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectUserList provides a mock function with given fields:
func (_m *BsmgUsecase) SelectUserList() ([]define.BsmgMemberInfo, error) {
	ret := _m.Called()

	var r0 []define.BsmgMemberInfo
	if rf, ok := ret.Get(0).(func() []define.BsmgMemberInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]define.BsmgMemberInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectWeekReportCategorySearch provides a mock function with given fields: pageInfo, partIdx
func (_m *BsmgUsecase) SelectWeekReportCategorySearch(pageInfo define.PageInfo, partIdx int) ([]define.BsmgWeekRptInfoForWeb, int32, error) {
	ret := _m.Called(pageInfo, partIdx)

	var r0 []define.BsmgWeekRptInfoForWeb
	if rf, ok := ret.Get(0).(func(define.PageInfo, int) []define.BsmgWeekRptInfoForWeb); ok {
		r0 = rf(pageInfo, partIdx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]define.BsmgWeekRptInfoForWeb)
		}
	}

	var r1 int32
	if rf, ok := ret.Get(1).(func(define.PageInfo, int) int32); ok {
		r1 = rf(pageInfo, partIdx)
	} else {
		r1 = ret.Get(1).(int32)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(define.PageInfo, int) error); ok {
		r2 = rf(pageInfo, partIdx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SelectWeekReportInfo provides a mock function with given fields: wRptIdx
func (_m *BsmgUsecase) SelectWeekReportInfo(wRptIdx int) (define.BsmgWeekRptInfoForWeb, error) {
	ret := _m.Called(wRptIdx)

	var r0 define.BsmgWeekRptInfoForWeb
	if rf, ok := ret.Get(0).(func(int) define.BsmgWeekRptInfoForWeb); ok {
		r0 = rf(wRptIdx)
	} else {
		r0 = ret.Get(0).(define.BsmgWeekRptInfoForWeb)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(wRptIdx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectWeekReportList provides a mock function with given fields: pageInfo, searchData
func (_m *BsmgUsecase) SelectWeekReportList(pageInfo define.PageInfo, searchData define.SearchData) ([]define.BsmgWeekRptInfoForWeb, int32, error) {
	ret := _m.Called(pageInfo, searchData)

	var r0 []define.BsmgWeekRptInfoForWeb
	if rf, ok := ret.Get(0).(func(define.PageInfo, define.SearchData) []define.BsmgWeekRptInfoForWeb); ok {
		r0 = rf(pageInfo, searchData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]define.BsmgWeekRptInfoForWeb)
		}
	}

	var r1 int32
	if rf, ok := ret.Get(1).(func(define.PageInfo, define.SearchData) int32); ok {
		r1 = rf(pageInfo, searchData)
	} else {
		r1 = ret.Get(1).(int32)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(define.PageInfo, define.SearchData) error); ok {
		r2 = rf(pageInfo, searchData)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateReportInfo provides a mock function with given fields: report
func (_m *BsmgUsecase) UpdateReportInfo(report define.BsmgReportInfo) error {
	ret := _m.Called(report)

	var r0 error
	if rf, ok := ret.Get(0).(func(define.BsmgReportInfo) error); ok {
		r0 = rf(report)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: member
func (_m *BsmgUsecase) UpdateUser(member define.BsmgMemberInfo) error {
	ret := _m.Called(member)

	var r0 error
	if rf, ok := ret.Get(0).(func(define.BsmgMemberInfo) error); ok {
		r0 = rf(member)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateWeekReportInfo provides a mock function with given fields: report
func (_m *BsmgUsecase) UpdateWeekReportInfo(report define.BsmgWeekRptInfo) error {
	ret := _m.Called(report)

	var r0 error
	if rf, ok := ret.Get(0).(func(define.BsmgWeekRptInfo) error); ok {
		r0 = rf(report)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserLogout provides a mock function with given fields: c
func (_m *BsmgUsecase) UserLogout(c echo.Context) define.OnlyResult {
	ret := _m.Called(c)

	var r0 define.OnlyResult
	if rf, ok := ret.Get(0).(func(echo.Context) define.OnlyResult); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(define.OnlyResult)
	}

	return r0
}
