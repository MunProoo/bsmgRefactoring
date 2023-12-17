package database

import (
	"BsmgRefactoring/define"
	"fmt"
	"log"

	"github.com/blue1004jy/gorm"
	_ "github.com/go-sql-driver/mysql"
)

// DB연결 초기화
func (dbm *DBGormMaria) release() {
	if dbm.DB != nil {
		dbm.DB.Close()
		dbm.DB = nil
	}
}

func (dbm *DBGormMaria) ConnectMariaDB() (err error) {
	dbm.release()

	// config 파일로 받아오도록 수정
	connectionString := "root:0000@tcp(127.0.0.1:3306)/"
	// dbm.DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	dbm.DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		fmt.Printf("%v \n", err)
		// 로그
		// 연결 실패

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
	var count int64
	err := DBGorm.DB.Table("INFORMATION_SCHEMA.SCHEMATA").Where("SCHEMA_NAME = 'BSMG'").Count(&count).Error
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
	createQuery := fmt.Sprintf("CREATE DATABASE %s", DBNAME)
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
	dbm.release()

	connectionString := "root:0000@tcp(127.0.0.1:3306)/BSMG?charset=utf8&parseTime=True&loc=Local"
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

// BSMG 멤버 테이블 생성
func (dbm *DBGormMaria) CreateMemberTable() {
	exist := dbm.DB.HasTable(&define.BsmgMemberInfo{})
	if !exist {
		err := dbm.DB.CreateTable(&define.BsmgMemberInfo{}).Error
		if err != nil {
			// 로그
			// 생성 실패
			fmt.Printf("%v \n", err)
		}
	}
}

// GORM 사용법 확인 필요
func (dbm *DBGormMaria) InsertMember(member define.BsmgMemberInfo) (err error) {
	nextIdx := dbm.findMinIdx()
	member.Mem_Index = nextIdx

	// 왜 mem_index 인식 못하고
	// Error 1364 (HY000): Field 'mem_index' doesn't have a default value 뜰까?
	// err = dbm.DB.Debug().Create(member).Error
	queryString := fmt.Sprintf(`INSERT INTO 'bsmg_member_infos' (mem_index, mem_id, mem_password, mem_name,
		mem_rank, mem_part) VALUES (%d,%s,%s,%s,%s,%s)`, member.Mem_Index, member.Mem_ID, member.Mem_Password,
		member.Mem_Name, member.Mem_Rank, member.Mem_Part)
	err = dbm.DB.Debug().Exec(queryString).Error

	if err != nil {
		log.Printf("%v \n", err)
		return
	}
	return
}

// 가장 작은 user index 번호 return
func (dbm *DBGormMaria) findMinIdx() int32 {
	queryString := `SELECT MIN(t1.mem_index + 1) AS next_available_id
	FROM bsmg_member_infos t1
	LEFT JOIN bsmg_member_infos t2 ON t1.mem_index + 1 = t2.mem_index
	WHERE t2.mem_index IS NULL`

	var nextIdx int32
	dbm.DB.Exec(queryString).Scan(&nextIdx)
	return nextIdx
}
