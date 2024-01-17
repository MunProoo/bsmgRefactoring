package usecase

import (
	"BsmgRefactoring/define"
	"fmt"

	"github.com/robfig/cron"
)

func (su structBsmgUsecase) SelectUserList() (userList []define.BsmgMemberInfo, err error) {
	/*
		비즈니스 로직 담당
		특정 권한을 가진 사용자만 리스트를 가져오도록 하는 인가
		가져온 리스트에 대한 필터링, 정렬 작업 등
	*/

	fmt.Println("usecase GetBsmgUserList")
	return su.br.SelectUserList()
}

// 일일 업무보고 -> 주간 업무보고 취합 스케쥴링 생성
func (su structBsmgUsecase) CreateCron(config define.ScheduleConfig) {
	CronSpec := config

	server.WeekRptMaker = cron.New()
	server.WeekRptMaker.AddFunc(CronSpec, server.MakeWeekRpt)
	server.WeekRptMaker.Start()
}
