package main

import (
	"BsmgRefactoring/database"
	"BsmgRefactoring/define"
	"encoding/json"
	"os"
	"sync"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type ServerProcessor struct {
	dbManager    database.DatabaseManager
	State        uint16 // 서버의 상태
	mutex        sync.RWMutex
	WeekRptMaker *cron.Cron // 주간보고 스케쥴러
	Config       *define.Config
	reqCh        chan interface{}
	resCh        chan interface{}
}

func (server *ServerProcessor) ConnectDataBase() (err error) {
	err = server.dbManager.InitDBManager(server.Config.DBConfig)
	if err != nil {
		// 로그
		log.Printf("InitDBManager Failed . err = %v\n", err)
		return err
	}
	return
}

func (server *ServerProcessor) LoadConfig() error {
	data, err := os.ReadFile("config.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, server.Config)
	if err != nil {
		return err
	}

	return err
}

// 일일 업무보고 -> 주간 업무보고 취합 스케쥴링 생성
func (server *ServerProcessor) CreateCron() {
	CronSpec := server.Config.ScheduleConfig.Spec

	server.WeekRptMaker = cron.New()
	server.WeekRptMaker.AddFunc(CronSpec, server.MakeWeekRpt)
	server.WeekRptMaker.Start()
}
