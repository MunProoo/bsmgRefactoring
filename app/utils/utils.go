package utils

import (
	"fmt"
	"os"
	"path/filepath"
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
	t = t.AddDate(0, 0, -3) // t객체가 3일 후인 다음주로 들어오므로 이번 주로 변경

	month := t.Format("200601020000")[4:6]
	nWeek := times.GetNthWeekOfMonth(t)

	return fmt.Sprintf("%s월 %d주차 %s 주간 업무보고", month, nWeek, mem_name)

}

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

// 로그 파일 생성 (/logs/curTimeDir/ 날짜.logs)
func CreateLogFile() (file *os.File, err error) {
	curPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(err)

	// logs 디렉토리 생성
	if _, err = os.Stat(curPath + "/logs"); os.IsNotExist(err) {
		if err = os.Mkdir(curPath+"/logs", os.FileMode(0755)); err != nil {
			return
		}
	}

	// 날짜 디렉토리 생성
	timeInfo := time.Now()
	dirName := fmt.Sprintf("%s/logs/%0.4d%0.2d%0.2d", curPath, timeInfo.Year(), timeInfo.Month(), timeInfo.Day())
	if _, err = os.Stat(dirName); os.IsNotExist(err) {
		if err = os.Mkdir(dirName, os.FileMode(0755)); err != nil {
			return
		}
	}

	fileIndex := 0
	fileNamePrefix := fmt.Sprintf("%s/%s_%0.4d%0.2d%0.2d", dirName, "BSMG", timeInfo.Year(), timeInfo.Month(), timeInfo.Day())

	for fileIndex = 0; ; fileIndex++ {
		fileName := fmt.Sprintf("%s_%d.log", fileNamePrefix, fileIndex)
		if _, err = os.Stat(fileName); os.IsNotExist(err) {
			file, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // 모두에게 읽기 쓰기 권한
			if err != nil {
				return
			}
			fmt.Println(file.Name())
			return file, err
		}
	}

}
