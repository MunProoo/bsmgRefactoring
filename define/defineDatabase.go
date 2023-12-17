package define

type DBConfig struct {
	DatabaseIP       string
	DatabasePort     string
	DatabaseID       string
	DatabasePW       string
	DatabaseName     string
	ServerListenPort int
}

// 멤버 구조체 객체
type BsmgMemberInfo struct {
	Mem_Index    int32  `gorm:"type:int;AUTO_INCREMENT;primary_key;not_null"`
	Mem_ID       string `json:"mem_id" gorm:"type:varchar(20);unique_key"`
	Mem_Password string `json:"mem_pw" gorm:"type:varchar(50)"`
	Mem_Name     string `json:"mem_name" gorm:"type:varchar(50)"`
	Mem_Rank     string `json:"mem_rank" gorm:"type:int"`
	Mem_Part     string `json:"mem_part" gorm:"type:int"`
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
