package usecase

import (
	"BsmgRefactoring/define"
	"BsmgRefactoring/middleware"
	"BsmgRefactoring/utils"
	"fmt"
	"strings"
)

func (uc structBsmgUsecase) SelectUserList() (userList []define.BsmgMemberInfo, err error) {
	/*
		비즈니스 로직 담당
		특정 권한을 가진 사용자만 리스트를 가져오도록 하는 인가
		가져온 리스트에 대한 필터링, 정렬 작업 등
	*/

	fmt.Println("usecase GetBsmgUserList")
	return uc.rp.SelectUserList()
}

// 일일 업무보고 -> 주간 업무보고 취합 스케쥴링 생성
func (uc structBsmgUsecase) MakeWeekRpt(config define.ScheduleConfig) {
	middleware.PrintI(middleware.LogArg{"layer": "UseCase", "fn": "MakeWeekRpt", "text": "MakeWeekRpt is proceed"})
	bef7d, bef1d, now, t := utils.GetDate()

	userList, err := uc.rp.SelectUserList()
	if err != nil {
		middleware.PrintE(middleware.LogArg{"layer": "UseCase", "fn": "MakeWeekRpt", "text": "SelectUserList is Failed", "err": err})
		return
	}

	for _, user := range userList {
		rptList, err := uc.rp.SelectReportListAWeek(user.Mem_ID, bef7d, bef1d)
		if err != nil {
			middleware.PrintE(middleware.LogArg{"layer": "UseCase", "fn": "MakeWeekRpt", "text": "SelectReportListAWeek is Failed", "err": err})
			return
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

		partLeader, err := uc.rp.SelectPartLeader(user.Mem_Part) // 부서 팀장님의 아이디
		if err != nil {
			middleware.PrintE(middleware.LogArg{"layer": "UseCase", "fn": "MakeWeekRpt", "text": "SelectPartLeader is Failed", "err": err})
			return
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
		err = uc.rp.InsertWeekReport(weekRptInfo)
		if err != nil {
			middleware.PrintE(middleware.LogArg{"layer": "database", "fn": "MakeWeekRpt", "text": "InsertWeekReport is failed", "err": err})
		}
	}

}
