package server

import (
	"encoding/json"

	// "log"
	"os"

	"github.com/robfig/cron"
)

func (server *ServerProcessor) ConnectDataBase() (err error) {
	err = server.DBManager.InitDBManager(server.Config.DBConfig)
	if err != nil {
		// 로그
		// server.log.Error("InitDBManager Failed ", "error", err)
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
