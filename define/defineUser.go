package define

type Result struct {
	ResultCode int32 `json:"ResultCode"`
}

// 멤버 구조체 객체
type BsmgMemberInfo struct {
	Mem_Index    string `gorm:"type:int;AUTO_INCREMENT;primary_key"`
	Mem_ID       string `json:"mem_id" gorm:"type:varchar(20);unique_key"`
	Mem_Password string `json:"mem_pw" gorm:"type:varchar(50)"`
	Mem_Name     string `json:"mem_name" gorm:"type:varchar(50)"`
	Mem_Rank     string `json:"mem_rank" gorm:"type:int"`
	Mem_Part     string `json:"mem_part" gorm:"type:int"`
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
