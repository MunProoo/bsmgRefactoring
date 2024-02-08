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
	// * 0 - All
	// * 1 - Title
	// * 2 - Content
	// * 3 - Reporter
	SearchCombo int32  `json:"@d1#search_combo" enums:"0,1,2,3"`
	SearchInput string `json:"@d1#search_input"`
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

// 일일 업무보고 json Bind용 (MacOS에선 변수타입 때문에 bind 안되므로)
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
	Rpt_toRptID  string `json:"rpt_toRptID"`                        // 실제 보고대상 ID
}

// string으로 받은 필드들을 실제 DB 타입에 맞게 변환
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
	report.Rpt_toRpt = stringReport.Rpt_toRptID

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
	To_Rpt_Name   string
}

func (brName *BsmgIncludeNameReport) AllocationInfo() (BsmgReportInfo BsmgReportInfoForWeb) {
	BsmgReportInfo.Rpt_Idx = brName.BsmgReportInfo.Rpt_Idx
	BsmgReportInfo.Rpt_Reporter = brName.BsmgReportInfo.Rpt_Reporter
	BsmgReportInfo.Rpt_date = brName.BsmgReportInfo.Rpt_date
	BsmgReportInfo.Rpt_toRpt = brName.BsmgReportInfo.Rpt_toRpt
	BsmgReportInfo.Rpt_ref = brName.BsmgReportInfo.Rpt_ref
	BsmgReportInfo.Rpt_title = brName.BsmgReportInfo.Rpt_title
	BsmgReportInfo.Rpt_content = brName.BsmgReportInfo.Rpt_content
	BsmgReportInfo.Rpt_confirm = brName.BsmgReportInfo.Rpt_confirm
	BsmgReportInfo.Rpt_attr1 = brName.Rpt_attr1
	BsmgReportInfo.Rpt_attr2 = brName.Rpt_attr2
	BsmgReportInfo.Rpt_etc = brName.Rpt_etc

	BsmgReportInfo.Rpt_ReporterName = brName.Reporter_Name
	BsmgReportInfo.Rpt_toRptName = brName.To_Rpt_Name
	return
}

type BsmgIncludeNameWeekReport struct {
	BsmgWeekRptInfo
	Reporter_Name string
	To_Rpt_Name   string
}

func (brName *BsmgIncludeNameWeekReport) AllocationInfo() (BsmgReportInfo BsmgWeekRptInfoForWeb) {
	BsmgReportInfo.WRpt_Idx = brName.BsmgWeekRptInfo.WRpt_Idx
	BsmgReportInfo.WRpt_Reporter = brName.BsmgWeekRptInfo.WRpt_Reporter
	BsmgReportInfo.WRpt_Date = brName.BsmgWeekRptInfo.WRpt_Date
	BsmgReportInfo.WRpt_ToRpt = brName.BsmgWeekRptInfo.WRpt_ToRpt
	BsmgReportInfo.WRpt_Title = brName.BsmgWeekRptInfo.WRpt_Title
	BsmgReportInfo.WRpt_Content = brName.BsmgWeekRptInfo.WRpt_Content
	BsmgReportInfo.WRpt_Part = brName.BsmgWeekRptInfo.WRpt_Part
	BsmgReportInfo.WRpt_OmissionDate = brName.BsmgWeekRptInfo.WRpt_OmissionDate

	BsmgReportInfo.WRpt_ReporterName = brName.Reporter_Name
	BsmgReportInfo.WRpt_ToRptName = brName.To_Rpt_Name
	return
}

// 주간 업무 보고 json Binding용
type BsmgWeekRptInfoStringField struct {
	WRpt_Idx          string `json:"wRpt_idx"`
	WRpt_Reporter     string `json:"wRpt_reporter"`
	WRpt_Date         string `json:"wRpt_date"`
	WRpt_ToRpt        string `json:"wRpt_toRpt"`
	WRpt_Title        string `json:"wRpt_title"`
	WRpt_Content      string `json:"wRpt_content"`
	WRpt_Part         string `json:"wRpt_part"`
	WRpt_OmissionDate string `json:"wRpt_omissionDate"`
}

func (bwr *BsmgWeekRptInfoStringField) ParseReport() (report BsmgWeekRptInfo) {
	idx, _ := strconv.Atoi(bwr.WRpt_Idx)
	report.WRpt_Idx = int32(idx)
	report.WRpt_Reporter = bwr.WRpt_Reporter
	report.WRpt_Date = bwr.WRpt_Date
	report.WRpt_ToRpt = bwr.WRpt_ToRpt
	report.WRpt_Title = bwr.WRpt_Title
	report.WRpt_Content = bwr.WRpt_Content
	part, _ := strconv.Atoi(bwr.WRpt_Part)
	report.WRpt_Part = int32(part)
	report.WRpt_OmissionDate = bwr.WRpt_OmissionDate
	return
}

// 웹에서 아이디 대신 이름으로 보여주기 위한 객체
type BsmgReportInfoForWeb struct {
	Rpt_Idx          int32  `json:"rpt_idx"`      // 인덱스
	Rpt_Reporter     string `json:"rpt_reporter"` // 보고자
	Rpt_date         string `json:"rpt_date"`     // 보고 일자
	Rpt_toRpt        string `json:"rpt_toRpt"`    // 보고 대상
	Rpt_ref          string `json:"rpt_ref"`      // 참조 대상
	Rpt_title        string `json:"rpt_title"`    // 업무보고 제목
	Rpt_content      string `json:"rpt_content"`  // 업무보고 내용
	Rpt_attr1        int32  `json:"rpt_attr1"`    // 업무속성1(솔루션/제품)
	Rpt_attr2        int32  `json:"rpt_attr2"`    // 업무속성2 (이름)
	Rpt_etc          string `json:"rpt_etc"`      // 기타 특이사항
	Rpt_confirm      bool   `json:"rpt_confirm"`  // 보고서 확정 상태
	Rpt_ReporterName string `json:"rpt_reporter_name"`
	Rpt_toRptName    string `json:"rpt_toRpt_name"`
}

func (rptForWeb *BsmgReportInfoForWeb) ParseReport(report BsmgReportInfo) {
	rptForWeb.Rpt_Idx = report.Rpt_Idx
	rptForWeb.Rpt_Reporter = report.Rpt_Reporter
	rptForWeb.Rpt_date = report.Rpt_date
	rptForWeb.Rpt_toRpt = report.Rpt_toRpt
	rptForWeb.Rpt_ref = report.Rpt_ref
	rptForWeb.Rpt_title = report.Rpt_title
	rptForWeb.Rpt_content = report.Rpt_content
	rptForWeb.Rpt_attr1 = report.Rpt_attr1
	rptForWeb.Rpt_attr2 = report.Rpt_attr2
	rptForWeb.Rpt_etc = report.Rpt_etc
	rptForWeb.Rpt_confirm = report.Rpt_confirm
	rptForWeb.Rpt_ReporterName = ""
	rptForWeb.Rpt_toRptName = ""
}

// 웹에서 아이디 대신 이름으로 보여주기 위한 객체
type BsmgWeekRptInfoForWeb struct {
	WRpt_Idx          int32  `json:"wRpt_idx"`           // 인덱스
	WRpt_Reporter     string `json:"wRpt_reporter"`      // 보고자
	WRpt_Date         string `json:"wRpt_date"`          // 보고 일자
	WRpt_ToRpt        string `json:"wRpt_toRpt"`         // 보고 대상 아이디
	WRpt_Title        string `json:"wRpt_title"`         // 제목
	WRpt_Content      string `json:"wRpt_content"`       // 내용
	WRpt_Part         int32  `json:"wRpt_part"`          // 부서
	WRpt_OmissionDate string `json:"wRpt_omissionDate"`  // 보고 누락 날짜
	WRpt_ReporterName string `json:"wRpt_reporter_name"` // 보고자 이름
	WRpt_ToRptName    string `json:"wRpt_toRpt_name"`    // 보고대상 이름
}

func (rptForWeb *BsmgWeekRptInfoForWeb) ParseReport(report BsmgWeekRptInfo) {
	rptForWeb.WRpt_Idx = report.WRpt_Idx
	rptForWeb.WRpt_Reporter = report.WRpt_Reporter
	rptForWeb.WRpt_Date = report.WRpt_Date
	rptForWeb.WRpt_ToRpt = report.WRpt_ToRpt
	rptForWeb.WRpt_Title = report.WRpt_Title
	rptForWeb.WRpt_Content = report.WRpt_Content
	rptForWeb.WRpt_Part = report.WRpt_Part
	rptForWeb.WRpt_OmissionDate = report.WRpt_OmissionDate

	rptForWeb.WRpt_ReporterName = ""
	rptForWeb.WRpt_ToRptName = ""
}
