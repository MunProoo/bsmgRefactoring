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
	id := dbm.DBConfig.DatabaseID
	pw := dbm.DBConfig.DatabasePW
	ip := dbm.DBConfig.DatabaseIP
	port := dbm.DBConfig.DatabasePort
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/", id, pw, ip, port)

	// dbm.DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	dbm.DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Printf("ConnectMariaDB %v \n", err)
		// 로그
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
	err := DBGorm.DB.Table("INFORMATION_SCHEMA.SCHEMATA").Where("SCHEMA_NAME = ?", DBNAME).Count(&count).Error
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
	id := dbm.DBConfig.DatabaseID
	pw := dbm.DBConfig.DatabasePW
	ip := dbm.DBConfig.DatabaseIP
	port := dbm.DBConfig.DatabasePort
	connectionString := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`, id, pw, ip, port, DBNAME)
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

// BSMG 직급 테이블 생성
func (dbm *DBGormMaria) CreateRankTable() (err error) {
	log.Println("CreateRankTable")
	exist := dbm.DB.HasTable(&define.BsmgRankInfo{})
	if !exist {
		err := dbm.DB.CreateTable(&define.BsmgRankInfo{}).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
	return
}

// BSMG 부서 테이블 생성
func (dbm *DBGormMaria) CreatePartTable() (err error) {
	log.Println("CreatePartTable")
	exist := dbm.DB.HasTable(&define.BsmgPartInfo{})
	if !exist {
		err := dbm.DB.CreateTable(&define.BsmgPartInfo{}).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
	return
}

// BSMG 멤버 테이블 생성
func (dbm *DBGormMaria) CreateMemberTable() (err error) {
	log.Println("CreateMemberTable")
	exist := dbm.DB.HasTable(&define.BsmgMemberInfo{})
	if !exist {
		err = dbm.DB.CreateTable(&define.BsmgMemberInfo{}).Error
		if err != nil {
			log.Printf("%v \n", err)
			return
		}
		err = dbm.DB.Model(&define.BsmgMemberInfo{}).AddForeignKey("mem_rank", "bsmg_rank_infos(rank_idx)", "NO ACTION", "CASCADE").Error
		if err != nil {
			log.Printf("%v \n", err)
		}
		err = dbm.DB.Model(&define.BsmgMemberInfo{}).AddForeignKey("mem_part", "bsmg_part_infos(part_idx)", "NO ACTION", "CASCADE").Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
	return
}

// BSMG 업무속성(카테고리) 테이블 생성
func (dbm *DBGormMaria) CreateAttr1Table() (err error) {
	log.Println("CreateAttr1Table")
	exist := dbm.DB.HasTable(&define.BsmgAttr1Info{})
	if !exist {
		err := dbm.DB.CreateTable(&define.BsmgAttr1Info{}).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
	return
}

// BSMG 업무속성(이름) 테이블 생성
func (dbm *DBGormMaria) CreateAttr2Table() (err error) {
	log.Println("CreateAttr2Table")
	exist := dbm.DB.HasTable(&define.BsmgAttr2Info{})
	if !exist {
		err := dbm.DB.CreateTable(&define.BsmgAttr2Info{}).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
		err = dbm.DB.Model(&define.BsmgAttr2Info{}).AddForeignKey("attr1_idx", "bsmg_attr1_infos(attr1_idx)", "NO ACTION", "CASCADE").Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
	return
}

// BSMG 일일보고 테이블 생성
func (dbm *DBGormMaria) CreateDailyReportTable() (err error) {
	log.Println("CreateDailyReportTable")
	exist := dbm.DB.HasTable(&define.BsmgReportInfo{})
	if !exist {
		err := dbm.DB.CreateTable(&define.BsmgReportInfo{}).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
		err = dbm.DB.Model(&define.BsmgReportInfo{}).AddForeignKey("rpt_attr1", "bsmg_attr1_infos(attr1_idx)", "NO ACTION", "CASCADE").Error
		if err != nil {
			log.Printf("%v \n", err)
		}
		err = dbm.DB.Model(&define.BsmgReportInfo{}).AddForeignKey("rpt_attr2", "bsmg_attr2_infos(attr2_idx)", "NO ACTION", "CASCADE").Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
	return
}

// BSMG 일정 테이블 생성
func (dbm *DBGormMaria) CreateScheduleTable() (err error) {
	log.Println("CreateScheduleTable")
	exist := dbm.DB.HasTable(&define.BsmgScheduleInfo{})
	if !exist {
		err := dbm.DB.CreateTable(&define.BsmgScheduleInfo{}).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
	return
}

// BSMG 주간보고 테이블 생성
func (dbm *DBGormMaria) CreateWeekReportTable() (err error) {
	log.Println("CreateWeekReportTable")
	exist := dbm.DB.HasTable(&define.BsmgWeekRptInfo{})
	if !exist {
		err := dbm.DB.CreateTable(&define.BsmgWeekRptInfo{}).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
		err = dbm.DB.Model(&define.BsmgWeekRptInfo{}).AddForeignKey("w_rpt_part", "bsmg_part_infos(part_idx)", "NO ACTION", "CASCADE").Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
	return
}

// 사용자 등록
func (dbm *DBGormMaria) InsertMember(member define.BsmgMemberInfo) (err error) {
	nextIdx := dbm.FindMinIdx()
	if nextIdx == 0 {
		// 0은 INSERT 에러, doesn't have a default value
		nextIdx++
	}
	member.Mem_Idx = nextIdx

	err = dbm.DB.Debug().Create(&member).Error
	// queryString := fmt.Sprintf(`INSERT INTO bsmg_member_infos (mem_index, mem_id, mem_password, mem_name,
	// 	mem_rank, mem_part) VALUES (%d,%s,%s,%s,%s,%s)`, member.Mem_Index, member.Mem_ID, member.Mem_Password,
	// 	member.Mem_Name, member.Mem_Rank, member.Mem_Part)
	// err = dbm.DB.Debug().Exec(queryString).Error

	if err != nil {
		log.Printf("%v \n", err)
		return
	}
	return
}

// 가장 작은 user index 번호 return
func (dbm *DBGormMaria) FindMinIdx() int32 {
	queryString := `SELECT MIN(t1.mem_idx + 1) AS NextIdx
	FROM bsmg_member_infos t1
	LEFT JOIN bsmg_member_infos t2 ON t1.mem_idx + 1 = t2.mem_idx
	WHERE t2.mem_idx IS NULL;`

	NextIdx := define.NextIdx{}
	err := dbm.DB.Debug().Raw(queryString).Scan(&NextIdx).Error
	fmt.Println(err)
	fmt.Println(NextIdx)
	if err != nil {
		log.Printf("%v \n", err)
	}

	return NextIdx.Idx.Int32
}

func (dbm *DBGormMaria) SelectRankList() (rankList []define.BsmgRankInfo, err error) {
	dbWhere := dbm.DB.Model(define.BsmgRankInfo{})
	err = dbWhere.Debug().Find(&rankList).Error
	if err != nil {
		return nil, err
	}
	return
}

func (dbm *DBGormMaria) SelectPartist() (partList []define.BsmgPartInfo, err error) {
	dbWhere := dbm.DB.Model(define.BsmgPartInfo{})
	err = dbWhere.Debug().Find(&partList).Error
	if err != nil {
		return nil, err
	}
	return
}
