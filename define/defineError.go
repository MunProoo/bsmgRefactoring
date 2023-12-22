package define

const (
	Success = iota
	ErrorInvalidParameter
	ErrorSession
	ResultIsNull
	LoginIDNotExist
	LoginPWMismatch
	NotLoggedIn
	ParsingFailed
	NotAuthorizedUser
	DataBaseError
)
