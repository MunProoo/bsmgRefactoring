package define

import "time"

// Type Definition
type OmissionMap map[string]bool

// 보고가 있는 날짜 체크해서 true로
func (om *OmissionMap) SetRptDate(rpt_date string) {
	rpt_date = rpt_date[:8] // 시간은 제외하고 날짜만 보겠다.
	_, exists := (*om)[rpt_date]
	if exists {
		(*om)[rpt_date] = true
	}
}

// 보고 누락된 날짜 취합하여 return
func (om *OmissionMap) ExtractMap() string {
	omissionDate := ""
	for key, val := range *om {
		if !val {
			omissionDate += key + ", "
		}
	}
	/*
		보고가 아예 없는 경우 슬라이싱에서 panic 발생 여지 있음.
		하지만 보고가 아예 없는 경우 이 메서드 안 탈 것임.

		음.. 하지만 방어코드 작성이 나쁠건 없으니까 ..
	*/
	if len(omissionDate) > 3 {
		omissionDate = omissionDate[:len(omissionDate)-2]
	}

	return omissionDate
}

// 주간 업무 기간 세팅 (빠진 날짜 찾기위해)
func InitOmissionMap(t time.Time) (findOmission *OmissionMap) {
	// 초기 Init : 주말 제외
	findOmission = &OmissionMap{}
	for i := 0; i < 7; i++ {
		date := t.AddDate(0, 0, -7+i)
		if date.Weekday() == 6 || date.Weekday() == 0 { // 토요일이거나 일요일이면 true
			(*findOmission)[t.AddDate(0, 0, -7+i).Format("20060102")] = true
		} else {
			(*findOmission)[t.AddDate(0, 0, -7+i).Format("20060102")] = false
		}

	}
	return findOmission
}
