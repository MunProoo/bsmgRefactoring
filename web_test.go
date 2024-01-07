package main

import (
	"BsmgRefactoring/define"
	"BsmgRefactoring/server"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMakePartTree(t *testing.T) {
	// 서버 연결
	server := server.ServerProcessor{}
	server.ConnectDataBase()
	defer server.DBManager.DBGorm.Release()

	assert := assert.New(t)

	// echo instance 생성
	e := echo.New()

	// 가짜 http 요청 생성
	req := httptest.NewRequest(http.MethodGet, "/bsmg/setting/weekRptCategory", nil)
	res := httptest.NewRecorder()

	// 에코에 context 추가
	c := e.NewContext(req, res)

	// 핸들러 호출
	err := getWeekRptCategoryHandler(&server, c)

	// 에러 확인
	assert.NoError(err, "TestMakePartTree : 쿼리 , 통신간의 오류가 있다.")

	// http 응답 상태 확인
	assert.Equal(http.StatusOK, res.Code, "TestMakePartTree : http 상태 오류")

	// 응답 바디 확인
	expectedJSON := `{"ds_List":null,"ds_partTree":[{"label":"부서별 주간 업무보고","value":"1","parent":"0"},
	{"label":"연구소","value":"1-1","parent":"1"},{"label":"SW1팀","value":"1-2","parent":"1"},
	{"label":"SW2팀","value":"1-3","parent":"1"},{"label":"FW1팀","value":"1-4","parent":"1"},
	{"label":"FW2팀","value":"1-5","parent":"1"},{"label":"HW1팀","value":"1-6","parent":"1"},
	{"label":"HW2팀","value":"1-7","parent":"1"},{"label":"Mobile팀","value":"1-8","parent":"1"},
	{"label":"디자인팀","value":"1-9","parent":"1"},{"label":"광학기구팀","value":"1-10","parent":"1"},
	{"label":"연구관리팀","value":"1-11","parent":"1"}],"Result":{"ResultCode":0}}`
	var expectedMap map[string]interface{}
	err = json.Unmarshal([]byte(expectedJSON), &expectedMap)
	assert.NoError(err, "json Unmarshal Failed")

	var actualMap map[string]interface{}
	err = json.Unmarshal(res.Body.Bytes(), &actualMap)
	assert.NoError(err, "json Unmarshal Failed")

	assert.Equal(expectedMap["ds_partTree"], actualMap["ds_partTree"], "TestMakePartTree : 예상과 다름")

}

func getWeekRptCategoryHandler(server *server.ServerProcessor, c echo.Context) (err error) {
	log.Println("getWeekRptCategoryReq")

	var apiResponse define.BsmgTreeResult
	apiResponse.PartTreeList, err = server.DBManager.DBGorm.MakePartTree()
	if err != nil {
		log.Printf("%v \n", err)
		apiResponse.Result.ResultCode = define.ErrorDataBase
		return c.JSON(http.StatusOK, apiResponse)
	}

	apiResponse.Result.ResultCode = define.Success
	return c.JSON(http.StatusOK, apiResponse)
}
