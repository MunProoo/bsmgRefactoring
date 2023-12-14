package database

type DBInterface interface {
	ConnectMariaDB() (err error) // DB 연결
	IsExistBSMG() (err error)    // Database 존재여부
	CreateDataBase() (err error) // Database 생성
	ConnectBSMG() (err error)    // Database 연결
	CreateMemberTable()
}
