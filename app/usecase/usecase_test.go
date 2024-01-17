package usecase

import (
	"BsmgRefactoring/define"
	mocks "BsmgRefactoring/mocks/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 유닛 테스트 코드 작성

// 정상적인 상황에서 GetBsmgUserList 예상대로 동작하는지 확인
func TestGetUserListHappyPath(t *testing.T) {
	// mockUser := define.BsmgMemberInfo{Mem_Idx: 0, Mem_ID: "admin", Mem_Password: "admin", Mem_Name: "admin", Mem_Rank: 1, Mem_Part: 1}
	mr := &mocks.BsmgRepository{}                                  // 레포지토리 선언
	uc := NewBsmgUsecase(mr)                                       // usecase에 레포지토리 주입
	mr.On("SelectUserList").Return([]define.BsmgMemberInfo{}, nil) // mockRepository에 기대조건 설정 (return이 nil일 것이다)

	userList, err := uc.SelectUserList()
	expectedUserList := make([]define.BsmgMemberInfo, 0)
	// expectedUserList[0] = define.BsmgMemberInfo{
	// 	Mem_Idx: 0, Mem_ID: "admin", Mem_Password: "admin", Mem_Name: "admin", Mem_Rank: 1, Mem_Part: 1,
	// }

	assert.Nil(t, err, "GetBsmgUserList is not working")
	assert.Equal(t, userList, expectedUserList, "Member is not equal")

	mr.AssertExpectations(t) // 기대대로 동작했는지 확인
}
