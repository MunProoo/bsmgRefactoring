package define

import (
	"BsmgRefactoring/utils"
	"strconv"
)

// ResultCode만 필요할때
type OnlyResult struct {
	Result Result `json:"Result"`
}

// 페이징처리를 위한 count(쿼리로 불러온 열의 수)     -> 굳이 구조체로 만들어야 하나? : 놉. eXBuilder랑 통신하는 규격만 맞추면 된다.
type TotalCountData struct {
	Count int32 `json:"Count"`
}

type Result struct {
	ResultCode int32 `json:"ResultCode"`
}

type SearchData struct {
	SearchCombo int32
	SearchInput string
}

type AttrSearchData struct {
	AttrValue    int32
	AttrCategory int32
}

// 메인 화면의 tree 구조를 위한 결과물
type BsmgTreeResult struct {
	AttrTreeList []AttrTree `json:"ds_List"`
	PartTreeList []PartTree `json:"ds_partTree"`
	Result       Result     `json:"Result"`
}

// 업무속성을 트리 구조로 만든 객체
type AttrTree struct {
	Label  string `json:"label"`
	Value  string `json:"value"`
	Parent string `json:"parent"`
}

// 부서를 트리 구조로 만든 객체
type PartTree struct {
	Label  string `json:"label"`
	Value  string `json:"value"`
	Parent string `json:"parent"`
}

type BsmgRankPartResult struct {
	RankList []BsmgRankInfo `json:"ds_rank"`
	PartList []BsmgPartInfo `json:"ds_part"`
	Result   Result         `json:"Result"`
}

// 부서 변경시 보고대상 바로 해당 팀의 팀장급으로
type BsmgTeamLeaderResult struct {
	Part   PartStruct `json:"dm_part"`
	Result Result     `json:"Result"`
}
type PartStruct struct {
	PartIdx    int32  `json:"part_idx"`
	TeamLeader string `json:"team_leader"`
}

// 페이징처리
type PageInfo struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

// 주간 업무보고 조회시
type BsmgWeekRptResult struct {
	WeekReportList []BsmgWeekRptInfo `json:"ds_weekRptList"`
	WeekReportInfo BsmgWeekRptInfo   `json:"dm_weekRptInfo"`
	TotalCount     TotalCountData    `json:"totalCount"`
	Result         Result            `json:"Result"`
}

type BsmgReportInfoStringField struct {
	Rpt_idx      string `json:"rpt_idx"`
	Rpt_Reporter string `json:"rpt_reporter"`
	Rpt_date     string `json:"rpt_date" gorm:"type:varchar(30)"`   // 보고 일자
	Rpt_toRpt    string `json:"rpt_toRpt" gorm:"type:nvarchar(20)"` // 보고 대상
	Rpt_ref      string `json:"rpt_ref" gorm:"type:nvarchar(100)"`  // 참조 대상
	Rpt_title    string `json:"rpt_title" gorm:"type:nvarchar(40)"` // 업무보고 제목
	Rpt_content  string `json:"rpt_content" gorm:"type:text"`       // 업무보고 내용
	Rpt_attr1    string `json:"rpt_attr1" gorm:"type:int"`          // 업무속성1(솔루션/제품)
	Rpt_attr2    string `json:"rpt_attr2" gorm:"type:int"`          // 업무속성2 (이름)
	Rpt_etc      string `json:"rpt_etc" gorm:"type:nvarchar(50)"`   // 기타 특이사항
}

func (stringReport *BsmgReportInfoStringField) ParseReport() (report BsmgReportInfo) {
	idx, _ := strconv.Atoi(stringReport.Rpt_idx)
	report.Rpt_Idx = int32(idx)
	report.Rpt_Reporter = stringReport.Rpt_Reporter
	report.Rpt_date = stringReport.Rpt_date
	report.Rpt_toRpt = stringReport.Rpt_toRpt
	report.Rpt_ref = stringReport.Rpt_ref
	report.Rpt_title = stringReport.Rpt_title
	report.Rpt_content = stringReport.Rpt_content
	attr1, _ := strconv.Atoi(stringReport.Rpt_attr1)

	report.Rpt_attr1 = int32(attr1)
	attr2 := utils.GetAttr2Idx(stringReport.Rpt_attr2)

	report.Rpt_attr2 = int32(attr2)
	report.Rpt_etc = stringReport.Rpt_etc

	return
}

type BsmgScheduleInfoString struct {
	Rpt_Idx    string `json:"rpt_idx"`
	Sc_Content string `json:"sc_content"`
}

func (bs *BsmgScheduleInfoString) ParseSchedule() (schedule BsmgScheduleInfo) {
	idx, _ := strconv.Atoi(bs.Rpt_Idx)
	schedule.Rpt_Idx = int32(idx)
	schedule.Sc_Content = bs.Sc_Content
	return
}

type BsmgIncludeNameReport struct {
	BsmgReportInfo
	Reporter_Name string
	ToRpt_Name    string
}

func (brName *BsmgIncludeNameReport) ChangeIDToName() {
	brName.Rpt_Reporter = brName.Reporter_Name
	brName.Rpt_toRpt = brName.ToRpt_Name
}

type BsmgIncludeNameWeekReport struct {
	BsmgWeekRptInfo
	Reporter_Name string
	ToRpt_Name    string
}

func (brName *BsmgIncludeNameWeekReport) ChangeIDToName() {
	brName.WRpt_Reporter = brName.Reporter_Name
	brName.WRpt_ToRpt = brName.ToRpt_Name
}
