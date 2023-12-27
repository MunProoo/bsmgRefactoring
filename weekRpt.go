package main

import (
	"BsmgRefactoring/utils"
)

// 비즈니스 로직이라 볼 수 있으므로 이 파일은 따로 service에 놔야할 거 같은데 DB작업을 하려면 Server 변수가 필요하므로
// ServerProcessor도 같이 옮겨줘야함..
// 어떻게 옮겨야 할까

// 주간보고로 취합
func (server *ServerProcessor) MakeWeekRpt() {
	bef7d, bef1d, now, t := utils.GetDate()
	err := server.dbManager.MakeWeekRpt(bef7d, bef1d, now, t)
	if err != nil {
		log.Printf("%s MakeWeekRpt Faield : %v", now, err)
		return
	}

}
