package database

import (
	"BsmgRefactoring/define"
	"BsmgRefactoring/utils"
	"fmt"
	"log"
	"strings"
	"time"

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
			DatabasePW:   "12345",
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

		// 테이블 생성
		log.Println("Create Tables ... ")
		err = dbManager.CreateTables()
		if err != nil {
			log.Printf("CreateTables : %v \n", err)
			return err
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

func (dbManager *DatabaseManager) MakeWeekRpt(bef7d, bef1d, now string, t time.Time) (err error) {
	log.Println("dbManager.MakeWeekRpt is proceed")
	userList, err := dbManager.DBGorm.SelectUserList()
	if err != nil {
		return err
	}

	for _, user := range userList {
		rptList, err := dbManager.DBGorm.SelectReportListAWeek(user.Mem_ID, bef7d, bef1d)
		if err != nil {
			return err
		}

		if len(rptList) == 0 {
			continue
		}

		var findOmission *define.OmissionMap
		findOmission = define.InitOmissionMap(t) // 업무보고 없는 날짜 map에 할당할 것.
		weekContent := strings.Builder{}         // 주간보고 내용물
		for _, report := range rptList {
			weekContent.WriteString("📆")
			weekContent.WriteString(report.Rpt_date + "\n")
			weekContent.WriteString(report.Rpt_content + "\n")

			findOmission.SetRptDate(report.Rpt_date) // 보고가 있는 날짜는 map에서 true로 변경
		}
		omissionDate := findOmission.ExtractMap() // 보고 없는 날짜 취합

		partLeader, err := dbManager.DBGorm.SelectPartLeader(user.Mem_Part) // 부서 팀장님의 아이디
		if err != nil {
			return err
		}

		fmt.Println(weekContent.String())

		weekRptInfo := define.BsmgWeekRptInfo{
			WRpt_Reporter:     user.Mem_ID,
			WRpt_Date:         now,
			WRpt_Title:        utils.GetWeekRptTitle(user.Mem_Name, t),
			WRpt_ToRpt:        partLeader,
			WRpt_Content:      weekContent.String(),
			WRpt_Part:         user.Mem_Part,
			WRpt_OmissionDate: omissionDate,
		}
		err = dbManager.DBGorm.InsertWeekReport(weekRptInfo)
		if err != nil {
			return err
		}
	}

	return err
}
