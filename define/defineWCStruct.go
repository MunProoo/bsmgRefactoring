package define

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
	AttrTreeList []*AttrTree `json:"ds_List"`
	PartTreeList []*PartTree `json:"ds_partTree"`
	Result       Result      `json:"Result"`
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
	RankList []*BsmgRankInfo `json:"ds_rank"`
	PartList []*BsmgPartInfo `json:"ds_part"`
	Result   Result          `json:"Result"`
}

// 직급 구조체
type BsmgRankInfo struct {
	Rank_Idx  int32  `json:"rank_idx"`
	Rank_Name string `json:"rank_name"`
}

// 부서 구조체
type BsmgPartInfo struct {
	Part_Idx  int32  `json:"part_idx"`
	Part_Name string `json:"part_name"`
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
	ReportList   []*BsmgReportInfo   `json:"ds_rptList"`
	ScheduleList []*BsmgScheduleInfo `json:"ds_schedule"`
	ReportInfo   *BsmgReportInfo     `json:"dm_reportInfo"`
	TotalCount   TotalCountData      `json:"totalCount"`
	Result       Result              `json:"Result"`
}

// 주간 업무보고 조회시
type BsmgWeekRptResult struct {
	WeekReportList []*BsmgWeekRptInfo `json:"ds_weekRptList"`
	WeekReportInfo *BsmgWeekRptInfo   `json:"dm_weekRptInfo"`
	TotalCount     TotalCountData     `json:"totalCount"`
	Result         Result             `json:"Result"`
}

// 일일 업무보고서 객체
type BsmgReportInfo struct {
	Rpt_Idx      int32  `json:"rpt_idx"`
	Rpt_Reporter string `json:"rpt_reporter"`
	Rpt_date     string `json:"rpt_date"`
	Rpt_toRpt    string `json:"rpt_toRpt"`
	Rpt_ref      string `json:"rpt_ref"`
	Rpt_schedule string `json:"rpt_schedule"`
	Rpt_title    string `json:"rpt_title"`
	Rpt_content  string `json:"rpt_content"`
	Rpt_etc      string `json:"rpt_etc"`
	Rpt_attr1    string `json:"rpt_attr1"`
	Rpt_attr2    string `json:"rpt_attr2"`
	Rpt_confirm  bool   `json:"rpt_confirm"`
}

// 일일 업무보고서 일정 객체
type BsmgScheduleInfo struct {
	Rpt_Idx    int32  `json:"rpt_idx"`
	Sc_Content string `json:"sc_content"`
}

// 주간 업무 보고서 객체
type BsmgWeekRptInfo struct {
	WRpt_Idx          int32  `json:"wRpt_idx"` // Struct 필드는 항상 대문자로 시작해야 한다
	WRpt_Reporter     string `json:"wRpt_reporter"`
	WRpt_Date         string `json:"wRpt_date"`
	WRpt_ToRpt        string `json:"wRpt_toRpt"`
	WRpt_Title        string `json:"wRpt_title"`
	WRpt_Content      string `json:"wRpt_content"`
	WRpt_OmissionDate string `json:"wRpt_omissionDate"`
	WRpt_Part         int32  `json:"wRpt_part"`
}
