package usecase

import (
	"BsmgRefactoring/define"
	"BsmgRefactoring/middleware"
	"BsmgRefactoring/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

/*
	비즈니스 로직 담당
	특정 권한을 가진 사용자만 리스트를 가져오도록 하는 인가
	가져온 리스트에 대한 필터링, 정렬 작업 등
*/

// 권한 확인
func (uc structBsmgUsecase) AuthorizationCheck(c echo.Context) (apiResponse define.BsmgMemberResponse, ResultCode int) {
	// 인가 체크
	claims, err, loginPossible := middleware.CheckToken(c, uc.loginUserAgentMap)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"Token Err": err})
		ResultCode = define.ErrorInvalidToken
		return
	} else if !loginPossible {
		ResultCode = define.ErrorLoginFailed
		return
	}

	MemberInfo := define.BsmgMemberInfo{}
	err = MemberInfo.ParsingClaim(claims)
	if err != nil || MemberInfo.Mem_Rank > define.Rank3 {
		middleware.PrintE(middleware.LogArg{"Authorization Err": err})
		ResultCode = define.ErrorNotAuthorizedUser
		return
	}
	ResultCode = define.Success
	apiResponse.MemberInfo = MemberInfo
	return
}

// 로그인만 확인
func (uc structBsmgUsecase) CheckLoginIng(c echo.Context) (apiResponse define.BsmgMemberResponse, ResultCode int) {
	// 인가 체크
	claims, err, loginPossible := middleware.CheckToken(c, uc.loginUserAgentMap)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"Token Err": err})
		ResultCode = define.ErrorInvalidToken
		return
	} else if !loginPossible {
		ResultCode = define.ErrorLoginFailed
		return
	}

	MemberInfo := define.BsmgMemberInfo{}
	err = MemberInfo.ParsingClaim(claims)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"Authorization Err": err})
		ResultCode = define.ErrorNotAuthorizedUser
		return
	}
	ResultCode = define.Success
	apiResponse.MemberInfo = MemberInfo
	return
}

func (uc structBsmgUsecase) SelectUserList() (userList []define.BsmgMemberInfo, err error) {

	fmt.Println("usecase GetBsmgUserList")
	return uc.rp.SelectUserList()
}

// 일일 업무보고 -> 주간 업무보고 취합 스케쥴링 생성
func (uc structBsmgUsecase) MakeWeekRpt() {
	middleware.PrintI(middleware.LogArg{"layer": "UseCase", "fn": "MakeWeekRpt", "text": "MakeWeekRpt is proceed"})
	bef7d, bef1d, now, t := utils.GetDate()

	userList, err := uc.rp.SelectUserList()
	if err != nil {
		middleware.PrintE(middleware.LogArg{"layer": "UseCase", "fn": "MakeWeekRpt", "text": "SelectUserList is Failed", "err": err})
		return
	}

	for _, user := range userList {
		rptList, err := uc.rp.SelectReportListAWeek(user.Mem_ID, bef7d, bef1d)
		if err != nil {
			middleware.PrintE(middleware.LogArg{"layer": "UseCase", "fn": "MakeWeekRpt", "text": "SelectReportListAWeek is Failed", "err": err})
			return
		}

		if len(rptList) == 0 {
			continue
		}

		// var findOmission *utils.OmissionMap
		findOmission := utils.InitOmissionMap(t) // 업무보고 없는 날짜 map에 할당할 것.
		weekContent := strings.Builder{}         // 주간보고 내용물
		for _, report := range rptList {
			weekContent.WriteString("📆")
			weekContent.WriteString(report.Rpt_date[:8] + "\n")
			weekContent.WriteString(report.Rpt_content + "\n")

			findOmission.SetRptDate(report.Rpt_date) // 보고가 있는 날짜는 map에서 true로 변경
		}
		omissionDate := findOmission.ExtractMap() // 보고 없는 날짜 취합

		partLeader, err := uc.rp.SelectPartLeader(user.Mem_Part) // 부서 팀장님의 아이디
		if err != nil {
			middleware.PrintE(middleware.LogArg{"layer": "UseCase", "fn": "MakeWeekRpt", "text": "SelectPartLeader is Failed", "err": err})
			return
		}

		fmt.Println(weekContent.String())

		weekRptInfo := define.BsmgWeekRptInfo{
			WRpt_Reporter:     user.Mem_ID,
			WRpt_Date:         now,
			WRpt_Title:        utils.GetWeekRptTitle(user.Mem_Name, t),
			WRpt_ToRpt:        partLeader,
			WRpt_Content:      weekContent.String(),
			WRpt_Part:         user.Mem_Part,
			WRpt_OmissionDate: omissionDate,
		}
		err = uc.rp.InsertWeekReport(weekRptInfo)
		if err != nil {
			middleware.PrintE(middleware.LogArg{"layer": "database", "fn": "MakeWeekRpt", "text": "InsertWeekReport is failed", "err": err})
			return
		}
	}

}

// 로그인 체크
func (uc structBsmgUsecase) CheckUserLogin(c echo.Context, apiRequest define.BsmgMemberLoginRequest) (apiResponse define.BsmgMemberResponse) {
	member := apiRequest.Data.MemberInfo
	err := uc.SelectMemberInfo(&member)
	if err != nil {
		if err.Error() == "record not found" {
			// 웹에서 에러코드 통해서 아이디 혹은 비밀번호가 틀립니다로
			middleware.PrintE(middleware.LogArg{"pn": "handlerLogin", "err": err})
			apiResponse.Result.ResultCode = define.ErrorLoginFailed
			return
		}
	}

	// 비밀번호 매칭 확인
	match, err := middleware.ComparePasswordAndHash(apiRequest.Data.MemberInfo.Mem_Password, member.Mem_Password)
	if err != nil || !match {
		middleware.PrintE(middleware.LogArg{"pn": "handlerLogin", "err": err})
		apiResponse.Result.ResultCode = define.ErrorLoginFailed
		return
	}
	// 인증 성공 ----------------------

	// 중복 로그인 방지
	if useAgent, exist := uc.loginUserAgentMap[member.Mem_ID]; exist {
		present_user_agent := c.Request().UserAgent()
		// 같은 브라우저면 X
		if useAgent == present_user_agent {
			middleware.PrintE(middleware.LogArg{"pn": "handlerLogin", "text": "This ID Login is Duplicated"})
			apiResponse.Result.ResultCode = define.ErrorLoginDuplication
			return
		}

		// 다른 브라우저면 로그인 OK (기존 연결 해제)
		uc.loginUserAgentMap[member.Mem_ID] = present_user_agent
	}

	apiResponse.Result.ResultCode = define.Success
	apiResponse.MemberInfo = member

	// JWT 토큰 생성
	// Set custom claims
	claims := &middleware.MemberClaims{
		Mem_ID:   member.Mem_ID,
		Mem_Name: member.Mem_Name,
		Mem_Rank: member.Mem_Rank,
		Mem_Part: member.Mem_Part,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	// AccessToken, RefreshToken , 쿠키 생성
	middleware.MakeJwtToken(c, claims)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerLogin", "err": err})
		apiResponse.Result.ResultCode = define.ErrorTokenCreationFailed
		return
	}
	// 로그인 여부 저장
	uc.loginUserAgentMap[member.Mem_ID] = c.Request().UserAgent()

	// 세션 저장
	// session := middleware.CreateSession(c, &apiResponse.MemberInfo)

	return
}

func (uc structBsmgUsecase) UserLogout(c echo.Context) (apiResponse define.OnlyResult) {
	middleware.DeleteCookie(c, middleware.AccessCookieName)
	middleware.DeleteCookie(c, middleware.RefreshCookieName)

	middleware.DeleteSession(c)
	apiResponse.Result.ResultCode = define.Success
	return
}

func (uc structBsmgUsecase) SelectReportListReq(searchData define.SearchData, pageInfo define.PageInfo) (apiResponse define.BsmgReportListResponse) {
	var err error
	var totalCount int32
	apiResponse.ReportList, totalCount, err = uc.SelectReportList(pageInfo, searchData)
	// apiResponse.ReportList, totalCount, err = server.DBManager.DBGorm.SelectReportList(pageInfo, searchData)

	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return
	}

	apiResponse.TotalCount.Count = totalCount
	apiResponse.Result.ResultCode = define.Success
	return
}

func (uc structBsmgUsecase) PostReportReq(c echo.Context, apiRequest define.BsmgReportInfoRequest) (apiResponse define.BsmgReportInfoResponse) {
	// 세션으로 클라이언트 정보 Get
	session, err := session.Get(middleware.SessionKey, c)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorSession
		return
	}
	client := session.Values["Member"].(define.BsmgMemberInfo)

	// parsing
	report := apiRequest.Data.BsmgReportInfo.ParseReport()
	report.Rpt_Reporter = client.Mem_ID

	// DB 처리
	err = uc.InsertDailyReport(report)
	// err = server.DBManager.DBGorm.InsertDailyReport(report)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return
	}

	// 스케쥴 등록을 위한 idx 반환
	idx, err := uc.SelectLatestRptIdx(report.Rpt_Reporter)
	// idx, err := server.DBManager.DBGorm.SelectLatestRptIdx(report.Rpt_Reporter)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return
	}
	report.Rpt_Idx = idx

	apiResponse.ReportInfo.ParseReport(report)
	apiResponse.Result.ResultCode = define.Success
	return
}

func (uc structBsmgUsecase) PostRegistScheduleReq(c echo.Context, apiRequest define.BsmgPostScheduleRequest) (apiRespone define.OnlyResult) {
	var err error
	idx, _ := strconv.Atoi(apiRequest.Data.BsmgReportInfo.Rpt_idx)

	for _, scheduleString := range apiRequest.Data.BsmgScheduleInfo {
		schedule := scheduleString.ParseSchedule()

		schedule.Rpt_Idx = int32(idx)
		err = uc.InsertSchedule(schedule)
		// err = server.DBManager.DBGorm.InsertSchedule(schedule)
		if err != nil {
			middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
			apiRespone.Result.ResultCode = define.ErrorDataBase
		}
	}

	apiRespone.Result.ResultCode = define.Success
	return
}

func (uc structBsmgUsecase) PutReportReq(c echo.Context, report define.BsmgReportInfo) (apiResponse define.BsmgReportInfoResponse) {
	// 세션으로 클라이언트 정보 Get
	session, err := session.Get(middleware.SessionKey, c)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "usecase", "err": err})
		apiResponse.Result.ResultCode = define.ErrorSession
		return
	}
	client := session.Values["Member"].(define.BsmgMemberInfo)

	// 본인만 수정 가능
	if client.Mem_ID != report.Rpt_Reporter {
		middleware.PrintE(middleware.LogArg{"pn": "usecase", "text": "This User is not him"})
		apiResponse.Result.ResultCode = define.ErrorNotAuthorizedUser
		return
	}

	err = uc.UpdateReportInfo(report)
	// err = server.DBManager.DBGorm.UpdateReportInfo(report)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "usecase", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return
	}

	apiResponse.Result.ResultCode = define.Success
	return
}

func (uc structBsmgUsecase) PutScheduleReq(apiRequest define.BsmgPutScheduleRequest, idx int32) (apiResponse define.OnlyResult) {
	// (무엇이 바뀌었는지 특정할 수 없으므로 전부 삭제 후 재 삽입)
	// 기존 스케쥴 삭제
	err := uc.DeleteSchedule(idx)
	// err = server.DBManager.DBGorm.DeleteSchedule(idx)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return
	}

	for _, scheduleString := range apiRequest.Data.BsmgScheduleInfo {
		schedule := scheduleString.ParseSchedule()
		schedule.Rpt_Idx = idx

		// 신규 스케쥴 삽입
		err = uc.InsertSchedule(schedule)
		// err = server.DBManager.DBGorm.InsertSchedule(schedule)
		if err != nil {
			middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
			apiResponse.Result.ResultCode = define.ErrorDataBase
			return
		}
	}

	apiResponse.Result.ResultCode = define.Success
	return
}

func (uc structBsmgUsecase) DeleteReportReq(c echo.Context, rptIdx int32) (apiResponse define.OnlyResult) {
	if _, resultCode := uc.AuthorizationCheck(c); resultCode != define.Success {
		apiResponse.Result.ResultCode = define.ErrorNotAuthorizedUser
		return
	}

	err := uc.DeleteSchedule(rptIdx)
	// err = server.DBManager.DBGorm.DeleteSchedule(rptIdx)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return
	}

	err = uc.DeleteReport(rptIdx)
	// err = server.DBManager.DBGorm.DeleteReport(rptIdx)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return
	}

	apiResponse.Result.ResultCode = define.Success
	return
}

func (uc structBsmgUsecase) PutWeekRptReq(c echo.Context, report define.BsmgWeekRptInfo) (apiResponse define.OnlyResult) {
	// 세션으로 클라이언트 정보 Get
	session, err := session.Get(middleware.SessionKey, c)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorSession
		return
	}
	client := session.Values["Member"].(define.BsmgMemberInfo)

	// 본인만 수정 가능
	if client.Mem_ID != report.WRpt_Reporter {
		apiResponse.Result.ResultCode = define.ErrorNotAuthorizedUser
		return
	}

	err = uc.UpdateWeekReportInfo(report)
	// err = server.DBManager.DBGorm.UpdateWeekReportInfo(report)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return
	}
	apiResponse.Result.ResultCode = define.Success
	return
}

func (uc structBsmgUsecase) DeleteWeekRptReq(c echo.Context, wRptIdx int) (apiResponse define.OnlyResult) {
	if _, resultCode := uc.AuthorizationCheck(c); resultCode != define.Success {
		apiResponse.Result.ResultCode = define.ErrorNotAuthorizedUser
		return
	}

	err := uc.DeleteWeekReport(wRptIdx)
	// err = server.DBManager.DBGorm.DeleteWeekReport(wRptIdx)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerReport", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return
	}

	apiResponse.Result.ResultCode = define.Success
	return
}

func (uc structBsmgUsecase) GetUserListRequest() (result define.BsmgMemberListResponse) {
	userList, err := uc.SelectUserList()
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerUser", "err": err})
		result.Result.ResultCode = define.ErrorDataBase
		return
	}
	count := len(userList)
	result.TotalCount.Count = int32(count)
	if count > 0 {
		result.MemberList = userList
	}

	result.Result.ResultCode = define.Success
	return
}

func (uc structBsmgUsecase) GetIdCheckRequest(memID string) (apiResponse define.OnlyResult) {
	isExist, err := uc.CheckMemberIDDuplicate(memID)
	// isExist, err := server.DBManager.DBGorm.CheckMemberIDDuplicate(memID)
	if err != nil {
		// record not found는 nil로 오도록 처리함. 그외 문제만 DB에러로
		middleware.PrintE(middleware.LogArg{"pn": "handlerUser", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return
	}

	if isExist {
		apiResponse.Result.ResultCode = define.ErrorDuplicatedID
		return
	}

	apiResponse.Result.ResultCode = define.Success

	return
}

func (uc structBsmgUsecase) PostUserReq(c echo.Context, apiRequest define.BsmgMemberRequest) (apiResponse define.OnlyResult) {
	// 인가 체크
	if _, resultCode := uc.AuthorizationCheck(c); resultCode != define.Success {
		apiResponse.Result.ResultCode = int32(resultCode)
		return
	}
	member := apiRequest.Data.MemberInfo
	// argon2 사용하여 salting, hashing
	// Pass the plaintext password and parameters to our generateFromPassword
	encodedHash, err := middleware.GenerateFromPassword(member.Mem_Password)
	if err != nil {
		log.Printf("%v \n", err)
		// TODO : 암호화 전용 에러코드 생성 필요
		apiResponse.Result.ResultCode = define.ErrorInvalidParameter
		return
	}

	member.Mem_Password = encodedHash

	err = uc.InsertMember(member)
	// err = server.DBManager.DBGorm.InsertMember(member)
	if err != nil {
		middleware.PrintE(middleware.LogArg{"pn": "handlerUser", "err": err})
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return
	}

	apiResponse.Result.ResultCode = define.Success
	return

}

func (uc structBsmgUsecase) PutUserReq(c echo.Context, apiRequest define.BsmgPutMemberRequest) (apiResponse define.OnlyResult) {

	// 인가 체크
	if _, resultCode := uc.AuthorizationCheck(c); resultCode != define.Success {
		apiResponse.Result.ResultCode = int32(resultCode)
		return
	}
	var member define.BsmgMemberInfo

	for _, reqMember := range apiRequest.Data.MemberList {
		member = reqMember.ParseMember()
		uc.UpdateUser(member)
		// server.DBManager.DBGorm.UpdateUser(member)
	}

	apiResponse.Result.ResultCode = define.Success
	return
}
