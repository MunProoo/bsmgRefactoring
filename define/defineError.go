package define

// Web에 Alert를 위한 에러 코드

const (
	Success = iota
	ErrorInvalidParameter
	ErrorSession

	ErrorDataBase

	ErrorDuplicatedID
	ErrorLoginFailed
	ErrorNotLoggedIn

	ErrorNotAuthorizedUser
	ErrorTokenCreationFailed
	ErrorCookieExtractionFailed
	ErrorInvalidToken
)
