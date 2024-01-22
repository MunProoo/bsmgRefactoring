package repository

import (
	"BsmgRefactoring/database"
	"sync"
)

/*
	실제 애플리케이션에서 data를 가져오는 부분에 직/간접적으로 상호작용하는 layer
	데이터의 관점에서 생각하여 원본 데이터에 대한 근본적인 행위만 표현한다고 볼 수 있으므로
	행위로만 method 이름 설정
*/

type BsmgRepository interface {
	database.DBInterface
}

type structBsmgRepository struct {
	dm    database.DatabaseManagerInterface
	Mutex sync.RWMutex
}

func NewBsmgRepository(dm database.DatabaseManagerInterface) BsmgRepository {
	return &structBsmgRepository{dm, sync.RWMutex{}}
}
