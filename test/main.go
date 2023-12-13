package main

import (
	"BsmgRefactoring/define"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBinterface 통해서 메서드를 할당하기

func main() {
	db, err := initDBManager()

	fmt.Printf("%v and %v", db, err)

}

type DatabaseManager struct {
	DBConfig define.DBConfig
	DB       *gorm.DB
}

func (dbm *DatabaseManager) CreateDataBase() error {

	dbm.DB.Table("INFORMATION_SCHEMA.SCHEMATA").Where("SCHEMA_NAME = ")
	return nil
}

func initDBManager() (*DatabaseManager, error) {

	// 메모리에 저장하자 AES 256해서
	DBConfig := define.DBConfig{
		DatabaseIP:   "127.0.0.1",
		DatabaseID:   "root",
		DatabasePW:   "0000",
		DatabasePort: "3306",
		DatabaseName: "",
	}
	var DBManager *DatabaseManager
	DBManager.DBConfig = DBConfig

	// connectDB로 빼자
	connectionString := "root:12345@tcp(127.0.0.1:3306)/"

	var err error
	DBManager.DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		// 로그남기기
		return nil, err
	}

	return DBManager, nil

}
