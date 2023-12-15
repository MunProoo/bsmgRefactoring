package define

// 세션에 들어가는 값 (ID)
type SessionValue struct {
	ID         string
	RememberID int32
}

type SessionData struct {
	Authenticated bool
	Member        BsmgMemberInfo
}
