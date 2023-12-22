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
	SearchCombo string
	SearchInput string
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
	Part   Part   `json:"dm_part"`
	Result Result `json:"Result"`
}
type Part struct {
	PartIdx    int32  `json:"part_idx"`
	TeamLeader string `json:"team_leader"`
}

// 페이징처리
type PageInfo struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

// 일일 업무보고 조회시
type BsmgReportResult struct {
	ReportList   []BsmgReportInfo   `json:"ds_rptList"`
	ScheduleList []BsmgScheduleInfo `json:"ds_schedule"`
	ReportInfo   *BsmgReportInfo    `json:"dm_reportInfo"`
	TotalCount   TotalCountData     `json:"totalCount"`
	Result       Result             `json:"Result"`
}

// 주간 업무보고 조회시
type BsmgWeekRptResult struct {
	WeekReportList []BsmgWeekRptInfo `json:"ds_weekRptList"`
	WeekReportInfo BsmgWeekRptInfo   `json:"dm_weekRptInfo"`
	TotalCount     TotalCountData    `json:"totalCount"`
	Result         Result            `json:"Result"`
}

type BsmgMemberRequest struct {
	Data struct {
		MemberList []BsmgMemberInfo `json:"Src_memberList"` // ds_memberList
		MemberInfo BsmgMemberInfo   `json:"dm_memberInfo"`
		TotalCount TotalCountData   `json:"TotalCount"`
		Result     Result           `json:"Result"`
	} `json:"data"`
}

type BsmgMemberResponse struct {
	MemberList []BsmgMemberInfo `json:"Src_memberList"` // ds_memberList
	MemberInfo BsmgMemberInfo   `json:"dm_memberInfo"`
	TotalCount TotalCountData   `json:"TotalCount"`
	Result     Result           `json:"Result"`
}

// mac에선 json을 binding하려면 무조건 String이어야 함
// request binding용으로 전부 string인 구조체 선언 필요
type BsmgMemberInfoStringField struct {
	Mem_Idx      string `json:"mem_idx" gorm:"type:int;AUTO_INCREMENT;primary_key"`
	Mem_ID       string `json:"mem_id" gorm:"type:varchar(20);unique_key"`
	Mem_Password string `json:"mem_pw" gorm:"type:varchar(50)"`
	Mem_Name     string `json:"mem_name" gorm:"type:nvarchar(50)"`
	Mem_Rank     string `json:"mem_rank" gorm:"type:int"`
	Mem_Part     string `json:"mem_part" gorm:"type:int"`
}

type BsmgPutMemberRequest struct {
	Data struct {
		MemberList []BsmgMemberInfoStringField `json:"ds_putMember"` // ds_putMember
	} `json:"data"`
}

type BsmgReportInfoStringField struct {
	Rpt_date    string `json:"rpt_date" gorm:"type:varchar(30)"`   // 보고 일자
	Rpt_toRpt   string `json:"rpt_toRpt" gorm:"type:nvarchar(20)"` // 보고 대상
	Rpt_ref     string `json:"rpt_ref" gorm:"type:nvarchar(100)"`  // 참조 대상
	Rpt_title   string `json:"rpt_title" gorm:"type:nvarchar(40)"` // 업무보고 제목
	Rpt_content string `json:"rpt_content" gorm:"type:text"`       // 업무보고 내용
	Rpt_attr1   string `json:"rpt_attr1" gorm:"type:int"`          // 업무속성1(솔루션/제품)
	Rpt_attr2   string `json:"rpt_attr2" gorm:"type:int"`          // 업무속성2 (이름)
	Rpt_etc     string `json:"rpt_etc" gorm:"type:nvarchar(50)"`   // 기타 특이사항
}

type BsmgPostReportRequest struct {
	Data struct {
		BsmgReportInfo BsmgReportInfoStringField `json:"dm_reportInfo"`
	} `json:"data"`
}

func (stringReport *BsmgReportInfoStringField) ParseReport() (report BsmgReportInfo) {
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
