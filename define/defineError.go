package define

const (
	Success = iota
	ErrorInvalidParameter
	ErrorSession

	ErrorDataBase

	ErrorDuplicatedID
	ErrorLoginFailed
	ErrorNotLoggedIn

	ErrorNotAuthorizedUser
)
