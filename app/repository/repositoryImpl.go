package repository

import (
	"BsmgRefactoring/define"
)

func (rp *structBsmgRepository) ConnectMariaDB() (err error) {
	return rp.dm.ConnectMariaDB()
}
func (rp *structBsmgRepository) IsConnected() (err error) {
	return rp.dm.IsConnected()
}
func (rp *structBsmgRepository) IsExistBSMG() (err error) {
	return rp.dm.IsExistBSMG()
}

// create
func (rp *structBsmgRepository) CreateDataBase() (err error) {
	return rp.dm.CreateDataBase()
}
func (rp *structBsmgRepository) ConnectBSMG() (err error) {
	return rp.dm.ConnectBSMG()
}
func (rp *structBsmgRepository) CreateMemberTable() (err error) {
	return rp.dm.CreateMemberTable()
}
func (rp *structBsmgRepository) CreateRankTable() (err error) {
	return rp.dm.CreateRankTable()
}
func (rp *structBsmgRepository) CreatePartTable() (err error) {
	return rp.dm.CreatePartTable()
}
func (rp *structBsmgRepository) CreateAttr1Table() (err error) {
	return rp.dm.CreateAttr1Table()
}
func (rp *structBsmgRepository) CreateAttr2Table() (err error) {
	return rp.dm.CreateAttr2Table()
}
func (rp *structBsmgRepository) CreateDailyReportTable() (err error) {
	return rp.dm.CreateDailyReportTable()
}
func (rp *structBsmgRepository) CreateScheduleTable() (err error) {
	return rp.dm.CreateScheduleTable()
}
func (rp *structBsmgRepository) CreateWeekReportTable() (err error) {
	return rp.dm.CreateWeekReportTable()
}

// insert
func (rp *structBsmgRepository) InsertDefaultAttr1() {
	rp.dm.InsertDefaultAttr1()
}
func (rp *structBsmgRepository) InsertDefaultAttr2() {
	rp.dm.InsertDefaultAttr2()
}
func (rp *structBsmgRepository) InsertDefaultRank() {
	rp.dm.InsertDefaultRank()
}
func (rp *structBsmgRepository) InsertDefaultPart() {
	rp.dm.InsertDefaultPart()
}

func (rp *structBsmgRepository) InsertMember(member define.BsmgMemberInfo) (err error) {
	rp.Mutex.Lock()
	defer rp.Mutex.Unlock()
	return rp.dm.InsertMember(member)
}
func (rp *structBsmgRepository) InsertDailyReport(report define.BsmgReportInfo) (err error) {
	rp.Mutex.Lock()
	defer rp.Mutex.Unlock()
	return rp.dm.InsertDailyReport(report)
}
func (rp *structBsmgRepository) InsertSchedule(schedule define.BsmgScheduleInfo) (err error) {
	rp.Mutex.Lock()
	defer rp.Mutex.Unlock()
	return rp.dm.InsertSchedule(schedule)
}
func (rp *structBsmgRepository) InsertWeekReport(wRptInfo define.BsmgWeekRptInfo) (err error) {
	return rp.dm.InsertWeekReport(wRptInfo)
}

// Select
func (rp *structBsmgRepository) SelectMemberInfo(member *define.BsmgMemberInfo) (err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.SelectMemberInfo(member)
}
func (rp *structBsmgRepository) SelectRankList() (rankList []define.BsmgRankInfo, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.SelectRankList()
}
func (rp *structBsmgRepository) SelectPartist() (partList []define.BsmgPartInfo, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.SelectPartist()
}
func (rp *structBsmgRepository) SelectUserList() (userList []define.BsmgMemberInfo, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.SelectUserList()
}

func (rp *structBsmgRepository) SelectReportList(pageInfo define.PageInfo, searchData define.SearchData) (rptList []define.BsmgReportInfoForWeb, totalCount int32, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.SelectReportList(pageInfo, searchData)
}
func (rp *structBsmgRepository) SelecAttrSearchReportList(pageInfo define.PageInfo, attrData define.AttrSearchData) (rptList []define.BsmgReportInfoForWeb, totalCount int32, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.SelecAttrSearchReportList(pageInfo, attrData)
}
func (rp *structBsmgRepository) SelectReportInfo(idx int) (reportInfoForWeb define.BsmgReportInfoForWeb, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.SelectReportInfo(idx)
}
func (rp *structBsmgRepository) SelectLatestRptIdx(reporter string) (rptIdx int32, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.SelectLatestRptIdx(reporter)
}
func (rp *structBsmgRepository) SelectScheduleList(rptIdx int32) (scheduleList []define.BsmgScheduleInfo, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.SelectScheduleList(rptIdx)
}
func (rp *structBsmgRepository) CheckMemberIDDuplicate(memID string) (isExist bool, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.CheckMemberIDDuplicate(memID)
}
func (rp *structBsmgRepository) SelectMemberListSearch(searchData define.SearchData) (memberList []define.BsmgMemberInfo, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.SelectMemberListSearch(searchData)
}
func (rp *structBsmgRepository) SelectReportListAWeek(Mem_ID, bef7d, bef1d string) (reportList []define.BsmgReportInfo, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.SelectReportListAWeek(Mem_ID, bef7d, bef1d)
}

func (rp *structBsmgRepository) SelectPartLeader(Mem_Part int32) (partLeader string, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.SelectPartLeader(Mem_Part)
}
func (rp *structBsmgRepository) SelectWeekReportList(pageInfo define.PageInfo, searchData define.SearchData) (rptList []define.BsmgWeekRptInfoForWeb, totalCount int32, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.SelectWeekReportList(pageInfo, searchData)
}
func (rp *structBsmgRepository) SelectWeekReportCategorySearch(pageInfo define.PageInfo, partIdx int) (rptList []define.BsmgWeekRptInfoForWeb, totalCount int32, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.SelectWeekReportCategorySearch(pageInfo, partIdx)
}
func (rp *structBsmgRepository) SelectWeekReportInfo(wRptIdx int) (rptInfo define.BsmgWeekRptInfoForWeb, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.SelectWeekReportInfo(wRptIdx)
}
func (rp *structBsmgRepository) SelectAttr1List() (attr1List []define.BsmgAttr1Info, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.SelectAttr1List()
}

// Update
func (rp *structBsmgRepository) UpdateUser(member define.BsmgMemberInfo) error {
	rp.Mutex.Lock()
	defer rp.Mutex.Unlock()
	return rp.dm.UpdateUser(member)
}
func (rp *structBsmgRepository) UpdateReportInfo(report define.BsmgReportInfo) (err error) {
	rp.Mutex.Lock()
	defer rp.Mutex.Unlock()
	return rp.dm.UpdateReportInfo(report)
}
func (rp *structBsmgRepository) UpdateWeekReportInfo(report define.BsmgWeekRptInfo) (err error) {
	rp.Mutex.Lock()
	defer rp.Mutex.Unlock()
	return rp.dm.UpdateWeekReportInfo(report)
}
func (rp *structBsmgRepository) ConfirmRpt(rptIdx int32) (err error) {
	rp.Mutex.Lock()
	defer rp.Mutex.Unlock()
	return rp.dm.ConfirmRpt(rptIdx)
}

// Delete
func (rp *structBsmgRepository) DeleteSchedule(rptIdx int32) (err error) {
	rp.Mutex.Lock()
	defer rp.Mutex.Unlock()
	return rp.dm.DeleteSchedule(rptIdx)
}
func (rp *structBsmgRepository) DeleteReport(rptIdx int32) (err error) {
	rp.Mutex.Lock()
	defer rp.Mutex.Unlock()
	return rp.dm.DeleteReport(rptIdx)
}
func (rp *structBsmgRepository) DeleteWeekReport(wRptIdx int) (err error) {
	rp.Mutex.Lock()
	defer rp.Mutex.Unlock()
	return rp.dm.DeleteWeekReport(wRptIdx)
}
func (rp *structBsmgRepository) DeleteMember(memId string) (err error) {
	rp.Mutex.Lock()
	defer rp.Mutex.Unlock()
	return rp.dm.DeleteMember(memId)
}

// util
func (rp *structBsmgRepository) FindMinIdx() int32 {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.FindMinIdx()
}
func (rp *structBsmgRepository) Release() {
	rp.dm.Release()
}
func (rp *structBsmgRepository) MakeAttrTree() (attrTreeList []define.AttrTree, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.MakeAttrTree()
}
func (rp *structBsmgRepository) MakePartTree() (partTreeList []define.PartTree, err error) {
	rp.Mutex.RLock()
	defer rp.Mutex.RUnlock()
	return rp.dm.MakePartTree()
}
