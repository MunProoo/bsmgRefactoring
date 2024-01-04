package define

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

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

// getToRpt
// 부서 변경시 보고대상 바로 해당 팀의 팀장급으로
type BsmgTeamLeaderResponse struct {
	Part   PartStruct `json:"dm_part"`
	Result Result     `json:"Result"`
}

// getAttr1Req
type BsmgAttr1Response struct {
	Attr1List []BsmgAttr1Info `json:"ds_attr1"`
	Result    Result          `json:"Result"`
}

func (bm *BsmgMemberInfo) ParsingClaim(claims jwt.MapClaims) error {
	var exist bool

	bm.Mem_ID, exist = claims["mem_id"].(string)
	if !exist {
		return errors.New("token doesn't have a member")
	}
	bm.Mem_Name, exist = claims["mem_name"].(string)
	if !exist {
		return errors.New("token doesn't have a member")
	}

	// 무조건 숫자는 float으로 옴
	memPartFloat, exist := claims["mem_part"].(float64)
	if !exist {
		return errors.New("token doesn't have a member")
	}
	bm.Mem_Part = int32(memPartFloat)

	// 무조건 숫자는 float으로 옴
	memRankFloat, exist := claims["mem_rank"].(float64)
	if !exist {
		return errors.New("token doesn't have a member")
	}
	bm.Mem_Rank = int32(memRankFloat)

	return nil
}
