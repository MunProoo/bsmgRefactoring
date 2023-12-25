package define

// 웹 클라이언트에 응답용 구조체만 정의

type BsmgMemberResponse struct {
	MemberInfo BsmgMemberInfo `json:"dm_memberInfo"`
	Result     Result         `json:"Result"`
}

type BsmgMemberListResponse struct {
	MemberList []BsmgMemberInfo `json:"Src_memberList"` // ds_memberList
	TotalCount TotalCountData   `json:"TotalCount"`
	Result     Result           `json:"Result"`
}

// 일일 업무보고 조회시
type BsmgReportResult struct {
	ReportList   []BsmgReportInfo   `json:"ds_rptList"`
	ScheduleList []BsmgScheduleInfo `json:"ds_schedule"`
	ReportInfo   *BsmgReportInfo    `json:"dm_reportInfo"`
	TotalCount   TotalCountData     `json:"totalCount"`
	Result       Result             `json:"Result"`
}

// getRptList 응답
type BsmgReportListResponse struct {
	ReportList []BsmgReportInfo `json:"ds_rptList"`
	TotalCount TotalCountData   `json:"totalCount"`
	Result     Result           `json:"Result"`
}

// getReportDetail
type BsmgReportInfoResponse struct {
	ReportInfo BsmgReportInfo `json:"dm_reportInfo"`
	Result     Result         `json:"Result"`
}

// getSchedull
type BsmgScheduleListResponse struct {
	ScheduleList []BsmgScheduleInfo `json:"ds_schedule"`
	Result       Result             `json:"Result"`
}

// getWeekReportList
type BsmgWeekReportListResponse struct {
	WeekReportList []BsmgWeekRptInfo `json:"ds_weekRptList"`
	TotalCount     TotalCountData    `json:"totalCount"`
	Result         Result            `json:"Result"`
}

// getWeekReportInfo
type BsmgWeekReportInfoResponse struct {
	WeekReportInfo BsmgWeekRptInfo `json:"dm_weekRptInfo"`
	Result         Result          `json:"Result"`
}
