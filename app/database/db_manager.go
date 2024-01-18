package database

import (
	"BsmgRefactoring/define"
	"BsmgRefactoring/middleware"
	"BsmgRefactoring/utils"
	"fmt"
	"strings"
	"time"

	"github.com/blue1004jy/gorm"
)

type DatabaseManagerInterface interface {
	InitDBManager(config define.DBConfig) (err error)
	CreateTables() (err error)
	MakeWeekRpt(bef7d, bef1d, now string, t time.Time) (err error)
	DBInterface
}

type DatabaseManager struct {
	DBGorm DBInterface
}

type DBGormMaria struct {
	DB       *gorm.DB
	DBConfig define.DBConfig `json:"database"`
}

func (dbManager *DatabaseManager) InitDBManager(config define.DBConfig) (err error) {
	// mariaDB 연결
	// dbManager.Log.Info("Connect DB ... ")
	middleware.PrintI(middleware.LogArg{"layer": "database", "fn": "NewDBManager", "text": "Connect DB ..."})

	dbManager.DBGorm = &DBGormMaria{
		DBConfig: config,
	}
	err = dbManager.DBGorm.ConnectMariaDB()
	if err != nil {
		// 로그남기기
		// dbManager.Log.Error("ConnectMariaDB Failed.", "error", err)
		middleware.PrintE(middleware.LogArg{"layer": "database", "fn": "InitDBManager", "text": "ConnectMariaDB Failed. ...", "err": err})
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
			// dbManager.Log.Error("CreateDataBase Failed ", "error", err)
			middleware.PrintE(middleware.LogArg{"layer": "database", "fn": "InitDBManager", "text": "CreateDataBase Failed ...", "err": err})
		}

		err = dbManager.DBGorm.ConnectBSMG()
		if err != nil {
			// 로그
			// Database connect Failed
			// dbManager.Log.Error("Database connect Failed ", "error", err)
			middleware.PrintE(middleware.LogArg{"layer": "database", "fn": "InitDBManager", "text": "Database connect Failed ...", "err": err})
			return
		}
		// 테이블 생성
		// log.Println("Create Tables ... ")
		middleware.PrintI(middleware.LogArg{"message": "Create Tables ... "})
		err = dbManager.CreateTables()
		if err != nil {
			// dbManager.Log.Error("CreateTables", "error", err)
			middleware.PrintE(middleware.LogArg{"layer": "database", "fn": "InitDBManager", "text": "CreateTables Failed ...", "err": err})
			return err
		}

		dbManager.DBGorm.InsertDefaultAttr1()
		dbManager.DBGorm.InsertDefaultAttr2()
	}

	// database 연결
	// dbManager.Log.Info("Connect BSMG ... ")
	middleware.PrintI(middleware.LogArg{"message": "Connect BSMG ... "})
	err = dbManager.DBGorm.ConnectBSMG()
	if err != nil {
		// 로그
		// Database connect Failed
		// dbManager.Log.Error("Database connect Failed", "error", err)
		middleware.PrintE(middleware.LogArg{"layer": "database", "fn": "InitDBManager", "text": "Database connect Failed..", "err": err})
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
	// dbManager.Log.Info("dbManager.MakeWeekRpt is proceed")
	middleware.PrintI(middleware.LogArg{"layer": "database", "fn": "MakeWeekRpt", "text": "MakeWeekRpt is proceed"})
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
			middleware.PrintE(middleware.LogArg{"layer": "database", "fn": "MakeWeekRpt", "text": "InsertWeekReport is failed", "err": err})
		}
	}

	return err
}

func NewDBManager(config define.DBConfig) DatabaseManagerInterface {
	dbManager := DatabaseManager{}

	// mariaDB 연결
	// dbManager.Log.Info("Connect DB ... ")
	middleware.PrintI(middleware.LogArg{"layer": "database", "fn": "NewDBManager", "text": "Connect DB ..."})

	dbManager.DBGorm = &DBGormMaria{
		DBConfig: config,
	}
	err := dbManager.DBGorm.ConnectMariaDB()
	if err != nil {
		// 로그남기기
		// dbManager.Log.Error("ConnectMariaDB Failed.", "error", err)
		middleware.PrintE(middleware.LogArg{"layer": "database", "fn": "NewDBManager", "text": "ConnectMariaDB failed", "err": err})
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
			// dbManager.Log.Error("CreateDataBase Failed ", "error", err)
			middleware.PrintE(middleware.LogArg{"layer": "database", "fn": "NewDBManager", "text": "CreateDataBase failed", "err": err})
		}

		err = dbManager.DBGorm.ConnectBSMG()
		if err != nil {
			// 로그
			// Database connect Failed
			// dbManager.Log.Error("Database connect Failed ", "error", err)
			middleware.PrintE(middleware.LogArg{"layer": "database", "fn": "NewDBManager", "text": "CreateDataBase failed", "err": err})
			return nil
		}
		// 테이블 생성
		// log.Println("Create Tables ... ")
		middleware.PrintE(middleware.LogArg{"layer": "database", "fn": "NewDBManager", "text": "Create Tables", "err": err})
		err = dbManager.CreateTables()
		if err != nil {
			// dbManager.Log.Error("CreateTables", "error", err)
			middleware.PrintE(middleware.LogArg{"layer": "database", "fn": "NewDBManager", "text": "CreateTables", "err": err})
		}

		dbManager.DBGorm.InsertDefaultAttr1()
		dbManager.DBGorm.InsertDefaultAttr2()
	}

	// database 연결
	// dbManager.Log.Info("Connect BSMG ... ")
	middleware.PrintI(middleware.LogArg{"layer": "database", "fn": "NewDBManager", "text": "Connect BSMG"})
	err = dbManager.DBGorm.ConnectBSMG()
	if err != nil {
		// 로그
		// Database connect Failed
		// dbManager.Log.Error("Database connect Failed", "error", err)
		middleware.PrintE(middleware.LogArg{"layer": "database", "fn": "NewDBManager", "text": "Database connect Failed", "err": err})
	}
	return &dbManager
}
