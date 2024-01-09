package database

import "BsmgRefactoring/app/define"

type DBInterface interface {
	// DB Connect
	ConnectMariaDB() (err error) // DB 연결
	IsConnected() (err error)    // DB연결 확인
	IsExistBSMG() (err error)    // Database 존재여부
	CreateDataBase() (err error) // Database 생성
	ConnectBSMG() (err error)    // Database 연결

	// Create table
	CreateMemberTable() (err error)
	CreateRankTable() (err error)
	CreatePartTable() (err error)
	CreateAttr1Table() (err error)
	CreateAttr2Table() (err error)
	CreateDailyReportTable() (err error)
	CreateScheduleTable() (err error)
	CreateWeekReportTable() (err error)

	// Insert
	InsertDefaultAttr1()
	InsertDefaultAttr2()
	InsertDefaultRank()
	InsertDefaultPart()
	InsertMember(member define.BsmgMemberInfo) (err error)
	InsertDailyReport(report define.BsmgReportInfo) (err error)
	InsertSchedule(schedule define.BsmgScheduleInfo) (err error)
	InsertWeekReport(wRptInfo define.BsmgWeekRptInfo) (err error)

	// Select
	SelectMemberInfo(member *define.BsmgMemberInfo) (err error)
	SelectRankList() (rankList []define.BsmgRankInfo, err error)
	SelectPartist() (partList []define.BsmgPartInfo, err error)
	SelectUserList() (userList []define.BsmgMemberInfo, err error)
	SelectReportList(pageInfo define.PageInfo, searchData define.SearchData) (rptList []define.BsmgReportInfo, totalCount int32, err error)
	SelecAttrSearchReportList(pageInfo define.PageInfo, attrData define.AttrSearchData) (rptList []define.BsmgReportInfo, totalCount int32, err error)
	SelectReportInfo(idx int) (rptInfo define.BsmgReportInfo, err error)
	SelectLatestRptIdx(reporter string) (rptIdx int32, err error)
	SelectScheduleList(rptIdx int32) (scheduleList []define.BsmgScheduleInfo, err error)
	CheckMemberIDDuplicate(memID string) (isExist bool, err error)
	SelectMemberListSearch(searchData define.SearchData) (memberList []define.BsmgMemberInfo, err error)
	SelectReportListAWeek(Mem_ID, bef7d, bef1d string) (reportList []define.BsmgReportInfo, err error)
	SelectPartLeader(Mem_Part int32) (partLeader string, err error)
	SelectWeekReportList(pageInfo define.PageInfo, searchData define.SearchData) (rptList []define.BsmgWeekRptInfo, totalCount int32, err error)
	SelectWeekReportCategorySearch(pageInfo define.PageInfo, partIdx int) (rptList []define.BsmgWeekRptInfo, totalCount int32, err error)
	SelectWeekReportInfo(wRptIdx int) (rptInfo define.BsmgWeekRptInfo, err error)
	SelectAttr1List() (attr1List []define.BsmgAttr1Info, err error)

	// Update
	UpdateUser(member define.BsmgMemberInfo) error
	UpdateReportInfo(report define.BsmgReportInfo) (err error)
	UpdateWeekReportInfo(report define.BsmgWeekRptInfo) (err error)
	ConfirmRpt(rptIdx int32) (err error)

	// Delete
	DeleteSchedule(rptIdx int32) (err error)
	DeleteReport(rptIdx int32) (err error)
	DeleteWeekReport(wRptIdx int) (err error)
	DeleteMember(memId string) (err error)

	// util
	FindMinIdx() int32
	Release()
	MakeAttrTree() (attrTreeList []define.AttrTree, err error)
	MakePartTree() (partTreeList []define.PartTree, err error)
}
