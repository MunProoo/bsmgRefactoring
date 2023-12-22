package database

import (
	"BsmgRefactoring/define"
	"log"

	"github.com/blue1004jy/gorm"
)

const (
	DBNAME = "BSMG"
)

type DatabaseManager struct {
	DBGorm DBInterface
}

type DBGormMaria struct {
	DB       *gorm.DB
	DBConfig define.DBConfig
}

func (dbManager *DatabaseManager) InitDBManager() (err error) {
	// 필요? -------------------------
	// 메모리에 저장하자 AES 256해서

	// mariaDB 연결
	log.Println("Connect DB ... ")
	dbManager.DBGorm = &DBGormMaria{
		DBConfig: define.DBConfig{
			DatabaseIP:   "127.0.0.1",
			DatabaseID:   "root",
			DatabasePW:   "0000",
			DatabasePort: "3306",
		},
	}
	err = dbManager.DBGorm.ConnectMariaDB()
	if err != nil {
		// 로그남기기
		log.Printf("ConnectMariaDB Failed . err = %v\n", err)
		return err
	}

	// BSMG Database 연결
	dbExist := false
	err = dbManager.DBGorm.IsExistBSMG()
	if err == nil {
		dbExist = true
	}

	// database 생성
	if !dbExist {
		err = dbManager.DBGorm.CreateDataBase()
		if err != nil {
			// 로그
			log.Printf("CreateDataBase Failed . err = %v\n", err)
		}
		dbManager.DBGorm.InsertDefaultAttr1()
		dbManager.DBGorm.InsertDefaultAttr2()
	}

	// database 연결
	log.Println("Connect BSMG ... ")
	err = dbManager.DBGorm.ConnectBSMG()
	if err != nil {
		// 로그
		// Database connect Failed
		log.Printf("Database connect Failed . err = %v\n", err)
		return
	}

	// 테이블 한번에 묶기
	log.Println("Create Tables ... ")
	err = dbManager.CreateTables()
	if err != nil {
		log.Printf("CreateTables : %v \n", err)
		return err
	}

	return nil
}

func (dbManager *DatabaseManager) CreateTables() (err error) {
	err = dbManager.DBGorm.CreateRankTable()
	if err != nil {
		return err
	}
	err = dbManager.DBGorm.CreatePartTable()
	if err != nil {
		return err
	}
	err = dbManager.DBGorm.CreateMemberTable()
	if err != nil {
		return err
	}
	err = dbManager.DBGorm.CreateAttr1Table()
	if err != nil {
		return err
	}
	err = dbManager.DBGorm.CreateAttr2Table()
	if err != nil {
		return err
	}
	err = dbManager.DBGorm.CreateDailyReportTable()
	if err != nil {
		return err
	}
	err = dbManager.DBGorm.CreateScheduleTable()
	if err != nil {
		return err
	}
	err = dbManager.DBGorm.CreateWeekReportTable()
	if err != nil {
		return err
	}

	return
}
