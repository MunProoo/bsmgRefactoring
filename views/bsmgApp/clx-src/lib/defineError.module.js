/************************************************
 * defineError.module.js
 * Created at 2023. 12. 29. 오전 11:24:39.
 *
 * @author SW2Team
 ************************************************/

globals.ErrorInvalidParameter			= 1 // 요청 파라미터 에러
globals.ErrorSession            		= 2 // 서버 세션 체크 에러
globals.ErrorDataBase                   = 3 // Database 에러
globals.ErrorDuplicatedID               = 4 // ID 중복
globals.ErrorLoginFailed     			= 5 // 로그인 실패
globals.ErrorNotLoggedIn                = 6 // 로그인하지 않은 상태
globals.ErrorNotAuthorizedUser          = 7 // 권한이 없음


globals.getErrorString = function( errCode ){
	var errMsg = "";
	errCode = Number(errCode)
	switch ( errCode ){				
		case ErrorInvalidParameter:  				errMsg = "요청값이 잘못 되었습니다."; break;
		case ErrorSession: 					errMsg ="세션이 끊겼습니다."; break;
		case ErrorDataBase: 					errMsg = "Database 에러입니다."; break;
		case ErrorLoginFailed:                 		errMsg = "아이디 혹은 비밀번호가 다릅니다."; break;
		case ErrorNotLoggedIn: 								errMsg = "ID가 중복되었습니다."; break; 
		case ErrorNotAuthorizedUser: 								errMsg = "권한이 없습니다."; break; 
		default : errMsg = "정의되지 않은 에러"; break;
	}
	return errMsg;
}