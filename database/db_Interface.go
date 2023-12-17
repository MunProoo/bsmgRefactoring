package database

import "BsmgRefactoring/define"

type DBInterface interface {
	ConnectMariaDB() (err error) // DB 연결
	IsConnected() (err error)    // DB연결 확인
	IsExistBSMG() (err error)    // Database 존재여부
	CreateDataBase() (err error) // Database 생성
	ConnectBSMG() (err error)    // Database 연결
	CreateMemberTable()
	InsertMember(member define.BsmgMemberInfo) (err error)
}
