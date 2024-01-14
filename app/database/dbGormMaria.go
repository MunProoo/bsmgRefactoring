package database

import (
	"BsmgRefactoring/define"
	"BsmgRefactoring/utils"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/blue1004jy/gorm"
	"github.com/inconshreveable/log15"
)

type DatabaseManager struct {
	DBGorm DBInterface
	Log    log15.Logger
}

type DBGormMaria struct {
	DB       *gorm.DB
	DBConfig define.DBConfig `json:"database"`
}

func (dbManager *DatabaseManager) InitDBManager(config define.DBConfig) (err error) {
	dbManager.Log = InitLog()

	// mariaDB 연결
	dbManager.Log.Info("Connect DB ... ")

	dbManager.DBGorm = &DBGormMaria{
		DBConfig: config,
	}
	err = dbManager.DBGorm.ConnectMariaDB()
	if err != nil {
		// 로그남기기
		dbManager.Log.Error("ConnectMariaDB Failed.", "error", err)
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
			dbManager.Log.Error("CreateDataBase Failed ", "error", err)
		}

		err = dbManager.DBGorm.ConnectBSMG()
		if err != nil {
			// 로그
			// Database connect Failed
			dbManager.Log.Error("Database connect Failed ", "error", err)
			return
		}
		// 테이블 생성
		log.Println("Create Tables ... ")
		err = dbManager.CreateTables()
		if err != nil {
			dbManager.Log.Error("CreateTables", "error", err)
			return err
		}

		dbManager.DBGorm.InsertDefaultAttr1()
		dbManager.DBGorm.InsertDefaultAttr2()
	}

	// database 연결
	dbManager.Log.Info("Connect BSMG ... ")
	err = dbManager.DBGorm.ConnectBSMG()
	if err != nil {
		// 로그
		// Database connect Failed
		dbManager.Log.Error("Database connect Failed", "error", err)
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
	dbManager.Log.Info("dbManager.MakeWeekRpt is proceed")
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

		var findOmission *utils.OmissionMap
		findOmission = utils.InitOmissionMap(t) // 업무보고 없는 날짜 map에 할당할 것.
		weekContent := strings.Builder{}        // 주간보고 내용물
		for _, report := range rptList {
			weekContent.WriteString("📆")
			weekContent.WriteString(report.Rpt_date[:8] + "\n")
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

func InitLog() log15.Logger {
	log := log15.New("module", "database")

	// 스트림 핸들러 (터미널에 출력)
	streamHandler := log15.StreamHandler(os.Stdout, log15.TerminalFormat())

	// 파일 핸들러 (프로덕션 환경에서만 중요한 오류 기록)
	fileHandler := log15.LvlFilterHandler(
		log15.LvlError,
		log15.Must.FileHandler("errors.json", log15.JsonFormat()),
	)

	// 로거에 여러 핸들러 추가 (디버깅 중에는 모든 로그를 스트림에, 프로덕션 환경에서는 중요한 오류만 파일에)
	log.SetHandler(log15.MultiHandler(streamHandler, fileHandler))

	return log
}
