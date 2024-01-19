package define

// User Search Combobox Value
const (
	SearchUserAll = iota
	SearchUserName
	SearchUserRank
	SearchUserPart
)

// DailyReport Search Combobox Value
const (
	SearchReportAll = iota
	SearchReportTitle
	SearchReportContent
	SearchReportReporter
)

// WeekReport Search Combobox Value
const (
	SearchWeekReportAll = iota
	SearchWeekReportTitle
	SearchWeekReportContent
	SearchWeekReportReporter
)

// WeekReport Category
const (
	AllCategory = 0
)

// Server State
const (
	StateInit = iota
	StateConnected
	StateRunning
	StateDisconnected
)

// PartLeader 의 Idx 값..
const (
	PartLeader = 3
)

// 주간 업무보고 카테고리 이름
const (
	WeekCategoryName = "부서별 주간 업무보고"
)

// 직급 define
const (
	Rank1 = iota + 1
	Rank2
	Rank3
	Rank4
)
