package database

import "BsmgRefactoring/define"

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
	CreateDailyReport(report define.BsmgReportInfo) (err error)

	// Select
	SelectRankList() (rankList []define.BsmgRankInfo, err error)
	SelectPartist() (partList []define.BsmgPartInfo, err error)
	SelectUserList() (userList []define.BsmgMemberInfo, err error)

	// Update
	UpdateUser(setVal map[string]interface{}, memID string) error

	// util
	FindMinIdx() int32
	Release()
	MakeAttrTree() (attrTreeList []define.AttrTree, err error)
}
