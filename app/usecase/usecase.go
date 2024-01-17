package usecase

import (
	"BsmgRefactoring/define"
	"BsmgRefactoring/repository"
)

// 애플리케이션/비즈니스 로직이 담겨있는 레이어

// 각 상황에 따라 이럴 땐 이렇게, 저럴 땐 저렇게 하는 등의 애플리케이션 내에서 어떤 가치를 만들어내는 자동화된 코드의 흐름이다.

type BsmgUsecase interface {
	SelectUserList() (userList []define.BsmgMemberInfo, err error)
}

type structBsmgUsecase struct {
	br repository.BsmgRepository
}

func NewBsmgUsecase(br repository.BsmgRepository) BsmgUsecase {
	return &structBsmgUsecase{br}
}
