package define

type BsmgMemberResult struct {
	MemberList []*BsmgMemberInfo `json:"Src_memberList"` // ds_memberList
	MemberInfo *BsmgMemberInfo   `json:"dm_memberInfo"`
	TotalCount TotalCountData    `json:"TotalCount"`
	Result     Result            `json:"Result"`
}
