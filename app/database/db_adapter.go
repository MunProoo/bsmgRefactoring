package database

import "BsmgRefactoring/define"

// 쉬운 mockup 객체 생성위해 dbManager를 interface로 감싸도록 하기 위해 adapter 패턴 적용

func (dbManager *DatabaseManager) ConnectMariaDB() (err error) {
	return dbManager.DBGorm.ConnectMariaDB()
}
func (dbManager *DatabaseManager) IsConnected() (err error) {
	return dbManager.DBGorm.IsConnected()
}
func (dbManager *DatabaseManager) IsExistBSMG() (err error) {
	return dbManager.DBGorm.IsExistBSMG()
}

// create
func (dbManager *DatabaseManager) CreateDataBase() (err error) {
	return dbManager.DBGorm.CreateDataBase()
}
func (dbManager *DatabaseManager) ConnectBSMG() (err error) {
	return dbManager.DBGorm.ConnectBSMG()
}
func (dbManager *DatabaseManager) CreateMemberTable() (err error) {
	return dbManager.DBGorm.CreateMemberTable()
}
func (dbManager *DatabaseManager) CreateRankTable() (err error) {
	return dbManager.DBGorm.CreateRankTable()
}
func (dbManager *DatabaseManager) CreatePartTable() (err error) {
	return dbManager.DBGorm.CreatePartTable()
}
func (dbManager *DatabaseManager) CreateAttr1Table() (err error) {
	return dbManager.DBGorm.CreateAttr1Table()
}
func (dbManager *DatabaseManager) CreateAttr2Table() (err error) {
	return dbManager.DBGorm.CreateAttr2Table()
}
func (dbManager *DatabaseManager) CreateDailyReportTable() (err error) {
	return dbManager.DBGorm.CreateDailyReportTable()
}
func (dbManager *DatabaseManager) CreateScheduleTable() (err error) {
	return dbManager.DBGorm.CreateScheduleTable()
}
func (dbManager *DatabaseManager) CreateWeekReportTable() (err error) {
	return dbManager.DBGorm.CreateWeekReportTable()
}

// insert
func (dbManager *DatabaseManager) InsertDefaultAttr1() {
	dbManager.DBGorm.InsertDefaultAttr1()
}
func (dbManager *DatabaseManager) InsertDefaultAttr2() {
	dbManager.DBGorm.InsertDefaultAttr2()
}
func (dbManager *DatabaseManager) InsertDefaultRank() {
	dbManager.DBGorm.InsertDefaultRank()
}
func (dbManager *DatabaseManager) InsertDefaultPart() {
	dbManager.DBGorm.InsertDefaultPart()
}

func (dbManager *DatabaseManager) InsertMember(member define.BsmgMemberInfo) (err error) {
	return dbManager.DBGorm.InsertMember(member)
}
func (dbManager *DatabaseManager) InsertDailyReport(report define.BsmgReportInfo) (err error) {
	return dbManager.DBGorm.InsertDailyReport(report)
}
func (dbManager *DatabaseManager) InsertSchedule(schedule define.BsmgScheduleInfo) (err error) {
	return dbManager.DBGorm.InsertSchedule(schedule)
}
func (dbManager *DatabaseManager) InsertWeekReport(wRptInfo define.BsmgWeekRptInfo) (err error) {
	return dbManager.DBGorm.InsertWeekReport(wRptInfo)
}

// Select
func (dbManager *DatabaseManager) SelectMemberInfo(member *define.BsmgMemberInfo) (err error) {
	return dbManager.DBGorm.SelectMemberInfo(member)
}
func (dbManager *DatabaseManager) SelectRankList() (rankList []define.BsmgRankInfo, err error) {
	return dbManager.DBGorm.SelectRankList()
}
func (dbManager *DatabaseManager) SelectPartist() (partList []define.BsmgPartInfo, err error) {
	return dbManager.DBGorm.SelectPartist()
}
func (dbManager *DatabaseManager) SelectUserList() (userList []define.BsmgMemberInfo, err error) {
	return dbManager.DBGorm.SelectUserList()
}

func (dbManager *DatabaseManager) SelectReportList(pageInfo define.PageInfo, searchData define.SearchData) (rptList []define.BsmgReportInfo, totalCount int32, err error) {
	return dbManager.DBGorm.SelectReportList(pageInfo, searchData)
}
func (dbManager *DatabaseManager) SelecAttrSearchReportList(pageInfo define.PageInfo, attrData define.AttrSearchData) (rptList []define.BsmgReportInfo, totalCount int32, err error) {
	return dbManager.DBGorm.SelecAttrSearchReportList(pageInfo, attrData)
}
func (dbManager *DatabaseManager) SelectReportInfo(idx int) (rptInfo define.BsmgReportInfo, err error) {
	return dbManager.DBGorm.SelectReportInfo(idx)
}
func (dbManager *DatabaseManager) SelectLatestRptIdx(reporter string) (rptIdx int32, err error) {
	return dbManager.DBGorm.SelectLatestRptIdx(reporter)
}
func (dbManager *DatabaseManager) SelectScheduleList(rptIdx int32) (scheduleList []define.BsmgScheduleInfo, err error) {
	return dbManager.DBGorm.SelectScheduleList(rptIdx)
}
func (dbManager *DatabaseManager) CheckMemberIDDuplicate(memID string) (isExist bool, err error) {
	return dbManager.DBGorm.CheckMemberIDDuplicate(memID)
}
func (dbManager *DatabaseManager) SelectMemberListSearch(searchData define.SearchData) (memberList []define.BsmgMemberInfo, err error) {
	return dbManager.DBGorm.SelectMemberListSearch(searchData)
}
func (dbManager *DatabaseManager) SelectReportListAWeek(Mem_ID, bef7d, bef1d string) (reportList []define.BsmgReportInfo, err error) {
	return dbManager.DBGorm.SelectReportListAWeek(Mem_ID, bef7d, bef1d)
}

func (dbManager *DatabaseManager) SelectPartLeader(Mem_Part int32) (partLeader string, err error) {
	return dbManager.DBGorm.SelectPartLeader(Mem_Part)
}
func (dbManager *DatabaseManager) SelectWeekReportList(pageInfo define.PageInfo, searchData define.SearchData) (rptList []define.BsmgWeekRptInfo, totalCount int32, err error) {
	return dbManager.DBGorm.SelectWeekReportList(pageInfo, searchData)
}
func (dbManager *DatabaseManager) SelectWeekReportCategorySearch(pageInfo define.PageInfo, partIdx int) (rptList []define.BsmgWeekRptInfo, totalCount int32, err error) {
	return dbManager.DBGorm.SelectWeekReportCategorySearch(pageInfo, partIdx)
}
func (dbManager *DatabaseManager) SelectWeekReportInfo(wRptIdx int) (rptInfo define.BsmgWeekRptInfo, err error) {
	return dbManager.DBGorm.SelectWeekReportInfo(wRptIdx)
}
func (dbManager *DatabaseManager) SelectAttr1List() (attr1List []define.BsmgAttr1Info, err error) {
	return dbManager.DBGorm.SelectAttr1List()
}

// Update
func (dbManager *DatabaseManager) UpdateUser(member define.BsmgMemberInfo) error {
	return dbManager.DBGorm.UpdateUser(member)
}
func (dbManager *DatabaseManager) UpdateReportInfo(report define.BsmgReportInfo) (err error) {
	return dbManager.DBGorm.UpdateReportInfo(report)
}
func (dbManager *DatabaseManager) UpdateWeekReportInfo(report define.BsmgWeekRptInfo) (err error) {
	return dbManager.DBGorm.UpdateWeekReportInfo(report)
}
func (dbManager *DatabaseManager) ConfirmRpt(rptIdx int32) (err error) {
	return dbManager.DBGorm.ConfirmRpt(rptIdx)
}

// Delete
func (dbManager *DatabaseManager) DeleteSchedule(rptIdx int32) (err error) {
	return dbManager.DBGorm.DeleteSchedule(rptIdx)
}
func (dbManager *DatabaseManager) DeleteReport(rptIdx int32) (err error) {
	return dbManager.DBGorm.DeleteReport(rptIdx)
}
func (dbManager *DatabaseManager) DeleteWeekReport(wRptIdx int) (err error) {
	return dbManager.DBGorm.DeleteWeekReport(wRptIdx)
}
func (dbManager *DatabaseManager) DeleteMember(memId string) (err error) {
	return dbManager.DBGorm.DeleteMember(memId)
}

// util
func (dbManager *DatabaseManager) FindMinIdx() int32 {
	return dbManager.DBGorm.FindMinIdx()
}
func (dbManager *DatabaseManager) Release() {
	dbManager.DBGorm.Release()
}
func (dbManager *DatabaseManager) MakeAttrTree() (attrTreeList []define.AttrTree, err error) {
	return dbManager.DBGorm.MakeAttrTree()
}
func (dbManager *DatabaseManager) MakePartTree() (partTreeList []define.PartTree, err error) {
	return dbManager.DBGorm.MakePartTree()
}
