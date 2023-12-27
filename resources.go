package main

import (
	"BsmgRefactoring/database"
	"sync"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// func initConfig() (config *define.DBConfig, err error) {
// 	curPath, err := filepath.Abs(filepath.Dir(os.Args[0]))

// 	if _, err := os.Stat(curPath + "/dbConfig.json"); os.IsNotExist(err) {
// 		file, err := os.OpenFile(curPath+"/dbConfig.json", os.O_CREATE|os.O_RDWR, 0666)
// 		if err != nil {

// 		}
// 	}
// 	return
// }

type ServerProcessor struct {
	dbManager    database.DatabaseManager
	State        uint16 // 서버의 상태
	mutex        sync.RWMutex
	WeekRptMaker *cron.Cron
	reqCh        chan interface{}
	resCh        chan interface{}
}

func (server *ServerProcessor) ConnectDataBase() (err error) {
	err = server.dbManager.InitDBManager()
	if err != nil {
		// 로그
		log.Printf("InitDBManager Failed . err = %v\n", err)
		return err
	}
	return
}

// TODO : CronSpec과 DB정보 저장한 Config파일 만들어서 읽어오도록
const (
	//            Minute   Hour   Day      Month                  Day of Week
	//		 		0-59   0-23   1-31   1-12 or JAN-DEC		  0-6 or SUN-SAT
	// CronSpec = "47 15 * * FRI"
	CronSpec = "01 16 * * MON"
	// CronSpec = "0 11 * * THU"
)

// 일일 업무복 -> 주간 업무보고 취합 스케쥴링 생성
func (server *ServerProcessor) CreateCron() (err error) {
	server.WeekRptMaker = cron.New()
	// server.WeekRptMaker.AddFunc(CronSpec, makeWeekRpt)

	return nil
}
