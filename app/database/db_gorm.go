package database

import (
	"fmt"

	"github.com/blue1004jy/gorm"
	_ "github.com/go-sql-driver/mysql"
)

// DB연결 초기화
func (dbm *DBGormMaria) Release() {
	if dbm.DB != nil {
		dbm.DB.Close()
		dbm.DB = nil
	}
}

func (dbm *DBGormMaria) ConnectMariaDB() (err error) {
	dbm.Release()

	// config 파일로 받아오도록 수정
	id := dbm.DBConfig.DatabaseID
	pw := dbm.DBConfig.DatabasePW
	ip := dbm.DBConfig.DatabaseIP
	port := dbm.DBConfig.DatabasePort
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/", id, pw, ip, port)

	// dbm.DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	dbm.DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	// Open 실패해도 err nil이므로 Ping으로 연결 체크
	// if err = ito.exDbConn.Ping(); err != nil {
	if err = dbm.DB.DB().Ping(); err != nil {
		// out.Printe(out.LogArg{"pn": "ITO", "fn": "connectDB", "text": "connect fail", "err": err})
		dbm.DB.Close()
		dbm.DB = nil
		return err
	}
	// 로그
	// 연결 성공
	return nil
}

func (dbm *DBGormMaria) IsConnected() (err error) {
	err = dbm.DB.DB().Ping()
	if err != nil {
		// out.Printe(out.LogArg{"pn": "ITO", "fn": "connectDB", "text": "connect fail", "err": err})
		dbm.DB.Close()
		dbm.DB = nil
		return err
	}
	return nil
}

// DB 존재여부 확인
func (DBGorm *DBGormMaria) IsExistBSMG() error {
	dbname := DBGorm.DBConfig.DatabaseName
	var count int64
	err := DBGorm.DB.Table("INFORMATION_SCHEMA.SCHEMATA").Where("SCHEMA_NAME = ?", dbname).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		// 로그
		// 에러
		err = fmt.Errorf("'BSMG' is not exist")
		return err
	}
	return nil
}

// BSMG DB 생성
func (dbm *DBGormMaria) CreateDataBase() error {
	dbname := dbm.DBConfig.DatabaseName
	createQuery := fmt.Sprintf("CREATE DATABASE %s", dbname)
	err := dbm.DB.Exec(createQuery).Error
	if err != nil {
		// 로그
		// 연결 실패
		fmt.Println("CreateDataBase Failed . ExecuteQuery = ", createQuery)
		return err
	}

	return nil
}

// BSMG 연결
func (dbm *DBGormMaria) ConnectBSMG() (err error) {
	dbm.Release()
	id := dbm.DBConfig.DatabaseID
	pw := dbm.DBConfig.DatabasePW
	ip := dbm.DBConfig.DatabaseIP
	port := dbm.DBConfig.DatabasePort
	dbname := dbm.DBConfig.DatabaseName
	connectionString := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`, id, pw, ip, port, dbname)
	dbm.DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	if err = dbm.DB.DB().Ping(); err != nil {
		// out.Printe(out.LogArg{"pn": "ITO", "fn": "connectDB", "text": "connect fail", "err": err})
		dbm.DB.Close()
		dbm.DB = nil
		return err
	}
	return err

}
