package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/shinYeongHyeon/go-times"
)

// 트리에서 가져온 value를 다시 attr2_idx로
// ex) "1-11" -> 11
func GetAttr2Idx(rpt_attr2String string) int32 {
	idx := strings.Index(rpt_attr2String, "-")
	rpt_attr2String = rpt_attr2String[idx+1:]
	rpt_attr2, err := strconv.ParseInt(rpt_attr2String, 10, 32)
	if err != nil {
		return 0
	}

	return int32(rpt_attr2)
}

// 일일 업무보고 취합할 기간
func GetDate() (string, string, string, time.Time) {
	// t := time.Now()
	t := time.Now().AddDate(0, 0, 3)
	now := t.Format("20060102000000")
	bef7d := t.AddDate(0, 0, -7).Format("20060102000000") //7일전 (저번주목욜)
	bef1d := t.AddDate(0, 0, -1).Format("20060102000000") // 1일전 (이번주수욜)
	return bef7d, bef1d, now, t
}

// 주간업무보고 제목
func GetWeekRptTitle(mem_name string, t time.Time) string {
	t = t.AddDate(0, 0, -1) // t객체가 3일 후인 다음주로 들어오므로 이번 주로 변경

	month := t.Format("200601020000")[4:6]
	nWeek := times.GetNthWeekOfMonth(t)

	return fmt.Sprintf("%s월 %d주차 %s 주간 업무보고", month, nWeek, mem_name)

}
