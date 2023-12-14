package define

type Result struct {
	ResultCode int32 `json:"ResultCode"`
}

// 멤버 구조체 객체
type BsmgMemberInfo struct {
	Mem_ID       string `json:"mem_id"`
	Mem_Password string `json:"mem_pw"`
	Mem_Name     string `json:"mem_name"`
	Mem_Rank     string `json:"mem_rank"`
	Mem_Part     string `json:"mem_part"`
	Mem_Index    string `gorm:"AUTO_INCREMENT;primary_key"`
}

// 페이징처리를 위한 count(쿼리로 불러온 열의 수)     -> 굳이 구조체로 만들어야 하나? : 놉. eXBuilder랑 통신하는 규격만 맞추면 된다.
type TotalCountData struct {
	Count int32 `json:"Count"`
}

type BsmgMemberResult struct {
	MemberList []*BsmgMemberInfo `json:"Src_memberList"` // ds_memberList
	MemberInfo *BsmgMemberInfo   `json:"dm_memberInfo"`
	TotalCount TotalCountData    `json:"TotalCount"`
	Result     Result            `json:"Result"`
}
