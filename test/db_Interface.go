package main

type DBInterface interface {
	ConnectDB() (err error)      // DB 연결
	IsExistBSMG() (err error)    // Database 존재여부
	CreateDataBase() (err error) // Database 생성
	CreateMemberTable()
}
