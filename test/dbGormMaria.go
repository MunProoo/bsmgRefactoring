package main

import (
	"BsmgRefactoring/define"

	"github.com/blue1004jy/gorm"
)

type DatabaseManager struct {
	DBGorm   DBInterface
	DBConfig define.DBConfig
}

type DBGormMaria struct {
	DB *gorm.DB
}

func InitDBManager() (*DatabaseManager, error) {

	// 메모리에 저장하자 AES 256해서
	DBManager := &DatabaseManager{
		DBConfig: define.DBConfig{},
		DBGorm:   &DBGormMaria{},
	}

	DBConfig := define.DBConfig{
		DatabaseIP:   "127.0.0.1",
		DatabaseID:   "root",
		DatabasePW:   "0000",
		DatabasePort: "3306",
		DatabaseName: "",
	}
	DBManager.DBConfig = DBConfig

	err := DBManager.DBGorm.ConnectDB()
	if err != nil {
		// 로그남기기
		return nil, err
	}

	return DBManager, nil
}
