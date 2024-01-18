package usecase

import "BsmgRefactoring/define"

func (uc structBsmgUsecase) ConnectMariaDB() (err error) {
	return uc.rp.ConnectMariaDB()
}
func (uc structBsmgUsecase) IsConnected() (err error) {
	return uc.rp.IsConnected()
}
func (uc structBsmgUsecase) IsExistBSMG() (err error) {
	return uc.rp.IsExistBSMG()
}

// create
func (uc structBsmgUsecase) CreateDataBase() (err error) {
	return uc.rp.CreateDataBase()
}
func (uc structBsmgUsecase) ConnectBSMG() (err error) {
	return uc.rp.ConnectBSMG()
}
func (uc structBsmgUsecase) CreateMemberTable() (err error) {
	return uc.rp.CreateMemberTable()
}
func (uc structBsmgUsecase) CreateRankTable() (err error) {
	return uc.rp.CreateRankTable()
}
func (uc structBsmgUsecase) CreatePartTable() (err error) {
	return uc.rp.CreatePartTable()
}
func (uc structBsmgUsecase) CreateAttr1Table() (err error) {
	return uc.rp.CreateAttr1Table()
}
func (uc structBsmgUsecase) CreateAttr2Table() (err error) {
	return uc.rp.CreateAttr2Table()
}
func (uc structBsmgUsecase) CreateDailyReportTable() (err error) {
	return uc.rp.CreateDailyReportTable()
}
func (uc structBsmgUsecase) CreateScheduleTable() (err error) {
	return uc.rp.CreateScheduleTable()
}
func (uc structBsmgUsecase) CreateWeekReportTable() (err error) {
	return uc.rp.CreateWeekReportTable()
}

// insert
func (uc structBsmgUsecase) InsertDefaultAttr1() {
	uc.rp.InsertDefaultAttr1()
}
func (uc structBsmgUsecase) InsertDefaultAttr2() {
	uc.rp.InsertDefaultAttr2()
}
func (uc structBsmgUsecase) InsertDefaultRank() {
	uc.rp.InsertDefaultRank()
}
func (uc structBsmgUsecase) InsertDefaultPart() {
	uc.rp.InsertDefaultPart()
}

func (uc structBsmgUsecase) InsertMember(member define.BsmgMemberInfo) (err error) {
	return uc.rp.InsertMember(member)
}
func (uc structBsmgUsecase) InsertDailyReport(report define.BsmgReportInfo) (err error) {
	return uc.rp.InsertDailyReport(report)
}
func (uc structBsmgUsecase) InsertSchedule(schedule define.BsmgScheduleInfo) (err error) {
	return uc.rp.InsertSchedule(schedule)
}
func (uc structBsmgUsecase) InsertWeekReport(wRptInfo define.BsmgWeekRptInfo) (err error) {
	return uc.rp.InsertWeekReport(wRptInfo)
}

// Select
func (uc structBsmgUsecase) SelectMemberInfo(member *define.BsmgMemberInfo) (err error) {
	return uc.rp.SelectMemberInfo(member)
}
func (uc structBsmgUsecase) SelectRankList() (rankList []define.BsmgRankInfo, err error) {
	return uc.rp.SelectRankList()
}
func (uc structBsmgUsecase) SelectPartist() (partList []define.BsmgPartInfo, err error) {
	return uc.rp.SelectPartist()
}

func (uc structBsmgUsecase) SelectReportList(pageInfo define.PageInfo, searchData define.SearchData) (rptList []define.BsmgReportInfo, totalCount int32, err error) {
	return uc.rp.SelectReportList(pageInfo, searchData)
}
func (uc structBsmgUsecase) SelecAttrSearchReportList(pageInfo define.PageInfo, attrData define.AttrSearchData) (rptList []define.BsmgReportInfo, totalCount int32, err error) {
	return uc.rp.SelecAttrSearchReportList(pageInfo, attrData)
}
func (uc structBsmgUsecase) SelectReportInfo(idx int) (rptInfo define.BsmgReportInfo, err error) {
	return uc.rp.SelectReportInfo(idx)
}
func (uc structBsmgUsecase) SelectLatestRptIdx(reporter string) (rptIdx int32, err error) {
	return uc.rp.SelectLatestRptIdx(reporter)
}
func (uc structBsmgUsecase) SelectScheduleList(rptIdx int32) (scheduleList []define.BsmgScheduleInfo, err error) {
	return uc.rp.SelectScheduleList(rptIdx)
}
func (uc structBsmgUsecase) CheckMemberIDDuplicate(memID string) (isExist bool, err error) {
	return uc.rp.CheckMemberIDDuplicate(memID)
}
func (uc structBsmgUsecase) SelectMemberListSearch(searchData define.SearchData) (memberList []define.BsmgMemberInfo, err error) {
	return uc.rp.SelectMemberListSearch(searchData)
}
func (uc structBsmgUsecase) SelectReportListAWeek(Mem_ID, bef7d, bef1d string) (reportList []define.BsmgReportInfo, err error) {
	return uc.rp.SelectReportListAWeek(Mem_ID, bef7d, bef1d)
}

func (uc structBsmgUsecase) SelectPartLeader(Mem_Part int32) (partLeader string, err error) {
	return uc.rp.SelectPartLeader(Mem_Part)
}
func (uc structBsmgUsecase) SelectWeekReportList(pageInfo define.PageInfo, searchData define.SearchData) (rptList []define.BsmgWeekRptInfo, totalCount int32, err error) {
	return uc.rp.SelectWeekReportList(pageInfo, searchData)
}
func (uc structBsmgUsecase) SelectWeekReportCategorySearch(pageInfo define.PageInfo, partIdx int) (rptList []define.BsmgWeekRptInfo, totalCount int32, err error) {
	return uc.rp.SelectWeekReportCategorySearch(pageInfo, partIdx)
}
func (uc structBsmgUsecase) SelectWeekReportInfo(wRptIdx int) (rptInfo define.BsmgWeekRptInfo, err error) {
	return uc.rp.SelectWeekReportInfo(wRptIdx)
}
func (uc structBsmgUsecase) SelectAttr1List() (attr1List []define.BsmgAttr1Info, err error) {
	return uc.rp.SelectAttr1List()
}

// Update
func (uc structBsmgUsecase) UpdateUser(member define.BsmgMemberInfo) error {
	return uc.rp.UpdateUser(member)
}
func (uc structBsmgUsecase) UpdateReportInfo(report define.BsmgReportInfo) (err error) {
	return uc.rp.UpdateReportInfo(report)
}
func (uc structBsmgUsecase) UpdateWeekReportInfo(report define.BsmgWeekRptInfo) (err error) {
	return uc.rp.UpdateWeekReportInfo(report)
}
func (uc structBsmgUsecase) ConfirmRpt(rptIdx int32) (err error) {
	return uc.rp.ConfirmRpt(rptIdx)
}

// Delete
func (uc structBsmgUsecase) DeleteSchedule(rptIdx int32) (err error) {
	return uc.rp.DeleteSchedule(rptIdx)
}
func (uc structBsmgUsecase) DeleteReport(rptIdx int32) (err error) {
	return uc.rp.DeleteReport(rptIdx)
}
func (uc structBsmgUsecase) DeleteWeekReport(wRptIdx int) (err error) {
	return uc.rp.DeleteWeekReport(wRptIdx)
}
func (uc structBsmgUsecase) DeleteMember(memId string) (err error) {
	return uc.rp.DeleteMember(memId)
}

// util
func (uc structBsmgUsecase) FindMinIdx() int32 {
	return uc.rp.FindMinIdx()
}
func (uc structBsmgUsecase) Release() {
	uc.rp.Release()
}
func (uc structBsmgUsecase) MakeAttrTree() (attrTreeList []define.AttrTree, err error) {
	return uc.rp.MakeAttrTree()
}
func (uc structBsmgUsecase) MakePartTree() (partTreeList []define.PartTree, err error) {
	return uc.rp.MakePartTree()
}
