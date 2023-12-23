package define

import "strconv"

// 웹 클라이언트의 Request를 받는 구조체만 정의
type BsmgMemberLoginRequest struct {
	Data struct {
		MemberInfo BsmgMemberInfo `json:"dm_memberInfo"`
	} `json:"data"`
}

type BsmgMemberRequest struct {
	Data struct {
		MemberList []BsmgMemberInfo `json:"Src_memberList"` // ds_memberList
		MemberInfo BsmgMemberInfo   `json:"dm_memberInfo"`
		TotalCount TotalCountData   `json:"TotalCount"`
		Result     Result           `json:"Result"`
	} `json:"data"`
}

type BsmgPutMemberRequest struct {
	Data struct {
		MemberList []BsmgMemberInfoStringField `json:"ds_putMember"` // ds_putMember
	} `json:"data"`
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

func (bs *BsmgMemberInfoStringField) ParseMember() (member BsmgMemberInfo) {
	member.Mem_ID = bs.Mem_ID
	member.Mem_Password = bs.Mem_Password
	rank, _ := strconv.Atoi(bs.Mem_Rank)
	member.Mem_Rank = int32(rank)
	part, _ := strconv.Atoi(bs.Mem_Part)
	member.Mem_Part = int32(part)
	member.Mem_Name = bs.Mem_Name

	return
}

type BsmgReportInfoRequest struct {
	Data struct {
		BsmgReportInfo BsmgReportInfoStringField `json:"dm_reportInfo"`
	} `json:"data"`
}

type BsmgPostScheduleRequest struct {
	Data struct {
		BsmgReportInfo   BsmgReportInfoStringField `json:"dm_reportInfo"`
		BsmgScheduleInfo []BsmgScheduleInfoString  `json:"ds_schedule"`
	} `json:"data"`
}

type BsmgPutScheduleRequest struct {
	Data struct {
		BsmgScheduleInfo []BsmgScheduleInfoString `json:"ds_schedule"`
		RptIdx           struct {
			RptIdx string `json:"rpt_idx"`
		} `json:"dm_rptIdx"`
	} `json:"data"`
}
