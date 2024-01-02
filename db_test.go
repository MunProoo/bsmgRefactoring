package main

import (
	"BsmgRefactoring/define"
	"BsmgRefactoring/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinQuery(t *testing.T) {
	server := ServerProcessor{}
	server.ConnectDataBase()
	defer server.dbManager.DBGorm.Release()
	int2 := server.dbManager.DBGorm.FindMinIdx()

	// assert.err

	// t, expected, actual
	assert.Equal(t, int32(2), int2)
}

func TestTableSturctureUpdate(t *testing.T) {
	server := ServerProcessor{}
	server.ConnectDataBase()
	defer server.dbManager.DBGorm.Release()

	err := server.dbManager.DBGorm.CreateDailyReportTable()
	fmt.Printf("%v \n ", err)
	assert.NoError(t, err, "err있나본데")

}

func TestInsertDefaultDBData(t *testing.T) {
	server := ServerProcessor{}
	server.ConnectDataBase()
	defer server.dbManager.DBGorm.Release()

	// server.dbManager.DBGorm.InsertDefaultAttr1()
	// server.dbManager.DBGorm.InsertDefaultAttr2()
	server.dbManager.DBGorm.InsertDefaultRank()
	server.dbManager.DBGorm.InsertDefaultPart()

}

func TestMakeAttrTree(t *testing.T) {
	server := ServerProcessor{}
	server.ConnectDataBase()
	defer server.dbManager.DBGorm.Release()

	expectAttrTrees := make([]define.AttrTree, 3)
	expectAttrTrees[0].Label = "솔루션"
	expectAttrTrees[0].Value = "1"
	expectAttrTrees[0].Parent = "0"
	expectAttrTrees[1].Label = "제품"
	expectAttrTrees[1].Value = "2"
	expectAttrTrees[1].Parent = "0"
	expectAttrTrees[2].Label = "출입통제 솔루션"
	expectAttrTrees[2].Value = "1-1"
	expectAttrTrees[2].Parent = "1"

	attrTrees, err := server.dbManager.DBGorm.MakeAttrTree()
	assert.NoError(t, err, "뭐야 에러있어")
	fmt.Printf("%v\n", attrTrees)

	assert.Equal(t, expectAttrTrees[2], attrTrees[2], "만들어야하는거와 다름")
}

func TestSelectUser(t *testing.T) {
	server := ServerProcessor{}
	server.ConnectDataBase()
	defer server.dbManager.DBGorm.Release()

	userList, err := server.dbManager.DBGorm.SelectUserList()
	assert.NoError(t, err, "에러있네")
	fmt.Printf("%v \n", userList)

	assert.Equal(t, 1, len(userList), "틀리네")
}

func TestGetAWeekRpt(t *testing.T) {
	server := ServerProcessor{}
	server.ConnectDataBase()
	defer server.dbManager.DBGorm.Release()

	bef7d, bef1d, now, tms := utils.GetDate()
	err := server.dbManager.MakeWeekRpt(bef7d, bef1d, now, tms)
	assert.NoError(t, err, "WeekRpt는 실패했다")
}

func TestGetWeekRptList(t *testing.T) {
	server := ServerProcessor{}
	err := server.ConnectDataBase()
	assert.NoError(t, err, "DB 연결 실패")
	defer server.dbManager.DBGorm.Release()

	pageInfo := define.PageInfo{}
	pageInfo.Offset, pageInfo.Limit = int32(0), int32(6)

	searchData := define.SearchData{
		SearchCombo: 0,
		SearchInput: "",
	}

	rptList, _, err := server.dbManager.DBGorm.SelectWeekReportList(pageInfo, searchData)
	fmt.Printf("%v \n", rptList)
	assert.NoError(t, err, "WeekRpt 가져오기 실패했다")
}

func TestAllMemberEncryptToArgon2(t *testing.T) {
	// 모든 사용자 암호화 적용
	server := ServerProcessor{}
	err := server.ConnectDataBase()
	assert.NoError(t, err, "DB 연결 실패")
	defer server.dbManager.DBGorm.Release()

	userList, err := server.dbManager.DBGorm.SelectUserList()
	assert.NoError(t, err, "User Select 실패")
	for _, user := range userList {
		if user.Mem_ID == "argon2" {
			continue
		}
		encodedHash, err := generateFromPassword(user.Mem_Password)
		assert.NoError(t, err, "암호화 실패")

		user.Mem_Password = encodedHash
		server.dbManager.DBGorm.UpdateUser(user)

	}
}
