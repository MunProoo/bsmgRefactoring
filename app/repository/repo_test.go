package repository

import (
	"BsmgRefactoring/define"
	mocks "BsmgRefactoring/mocks/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 유닛 테스트 코드 작성

// 정상적인 상황에서 GetBsmgUserList 예상대로 동작하는지 확인
func TestGetUserListHappyPath(t *testing.T) {
	dm := &mocks.DatabaseManagerInterface{} // mockup DB 구현
	mr := NewBsmgRepository(dm)             // 레포지토리 선언 및 mockup 주입

	dm.On("SelectUserList").Return([]define.BsmgMemberInfo{}, nil) // mockup에 기대결과 구현
	userList, err := mr.SelectUserList()
	expectedUserList := make([]define.BsmgMemberInfo, 0)

	assert.Nil(t, err, "GetBsmgUserList is not working")
	assert.Equal(t, userList, expectedUserList, "Member is not equal")
	dm.AssertExpectations(t) // mockup이 기대대로 동작했는지 확인

}
