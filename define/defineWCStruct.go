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
