package define

type DBConfig struct {
	DatabaseIP       string
	DatabasePort     string
	DatabaseID       string
	DatabasePW       string
	ServerListenPort int
}

// 멤버 구조체 객체
type BsmgMemberInfo struct {
	Mem_Idx      int32  `json:"mem_idx" gorm:"type:int;AUTO_INCREMENT;primary_key"` // 추후 서버 메모리에 담을때 사용할 예정
	Mem_ID       string `json:"mem_id" gorm:"type:varchar(20);unique_key"`
	Mem_Password string `json:"mem_pw" gorm:"type:varchar(50)"`
	Mem_Name     string `json:"mem_name" gorm:"type:nvarchar(50)"`
	Mem_Rank     int32  `json:"mem_rank" gorm:"type:int"`
	Mem_Part     int32  `json:"mem_part" gorm:"type:int"`
}

// 직급 구조체
type BsmgRankInfo struct {
	Rank_Idx  int32  `json:"rank_idx" gorm:"type:int;auto_increment;primary_key"`
	Rank_Name string `json:"rank_name" gorm:"type:nvarchar(20)"`
}

// 부서 구조체
type BsmgPartInfo struct {
	Part_Idx  int32  `json:"part_idx" gorm:"type:int;auto_increment;primary_key"`
	Part_Name string `json:"part_name" gorm:"type:nvarchar(20)"`
}

// 업무속성1 : Category ('솔루션', '제품')
type BsmgAttr1Info struct {
	Attr1_Idx      int32  `json:"attr1_idx" gorm:"type:int;auto_increment;primary_key"`
	Attr1_Category string `json:"attr1_category" gorm:"type:nvarchar(20)"`
}

// 업무속성2 : Name (솔루션 or 제품의 '이름')
type BsmgAttr2Info struct {
	Attr2_Idx  int32  `json:"attr2_idx" gorm:"type:int;auto_increment;primary_key"`
	Attr1_Idx  int32  `json:"attr1_idx" gorm:"type:int"`
	Attr2_Name string `json:"attr2_name" gorm:"type:nvarchar(100)"`
}

// 일일 업무보고서 객체
type BsmgReportInfo struct {
	Rpt_Idx      int32  `json:"rpt_idx" gorm:"type:int;auto_increment;primary_key;not null"` // 인덱스
	Rpt_Reporter string `json:"rpt_reporter" gorm:"type:varchar(20)"`                        // 보고자
	Rpt_date     string `json:"rpt_date" gorm:"type:varchar(30)"`                            // 보고 일자
	Rpt_toRpt    string `json:"rpt_toRpt" gorm:"type:nvarchar(20)"`                          // 보고 대상
	Rpt_ref      string `json:"rpt_ref" gorm:"type:nvarchar(100)"`                           // 참조 대상
	Rpt_title    string `json:"rpt_title" gorm:"type:nvarchar(40)"`                          // 업무보고 제목
	Rpt_content  string `json:"rpt_content" gorm:"type:text"`                                // 업무보고 내용
	Rpt_attr1    int32  `json:"rpt_attr1" gorm:"type:int"`                                   // 업무속성1(솔루션/제품)
	Rpt_attr2    int32  `json:"rpt_attr2" gorm:"type:int"`                                   // 업무속성2 (이름)
	Rpt_etc      string `json:"rpt_etc" gorm:"type:nvarchar(50)"`                            // 기타 특이사항
	Rpt_confirm  bool   `json:"rpt_confirm" gorm:"type:tinyint(1)"`                          // 보고서 확정 상태
}

// 일일 업무보고서 일정 객체
type BsmgScheduleInfo struct {
	Rpt_Idx    int32  `json:"rpt_idx" gorm:"type:int;"`             // 일일 업무보고 인덱스 (1:N이므로 Not PK)
	Sc_Content string `json:"sc_content" gorm:"type:nvarchar(100)"` // 업무 일정 내용
}

// 주간 업무 보고서 객체
type BsmgWeekRptInfo struct {
	WRpt_Idx          int32  `json:"wRpt_idx" gorm:"type:int;primary_key"`      // 업무보고 인덱스
	WRpt_Reporter     string `json:"wRpt_reporter" gorm:"type:nvarchar(20)"`    // 업무보고자
	WRpt_Date         string `json:"wRpt_date" gorm:"type:nvarchar(30)"`        // 업무보고 일자
	WRpt_ToRpt        string `json:"wRpt_toRpt" gorm:"type:nvarchar(20)"`       // 업무보고 대상
	WRpt_Title        string `json:"wRpt_title" gorm:"type:nvarchar(40)"`       // 업무보고 제목
	WRpt_Content      string `json:"wRpt_content" gorm:"type:text"`             // 업무 내용
	WRpt_Part         int32  `json:"wRpt_part" gorm:"type:int"`                 // 부서
	WRpt_OmissionDate string `json:"wRpt_omissionDate" gorm:"type:varchar(50)"` // 보고서 누락 날짜
}
