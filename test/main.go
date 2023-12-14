package main

import (
	"fmt"
)

// DBinterface 통해서 메서드를 할당하고 싶은데 어떻게 가야하나? 어댑터패턴으로 가야하나

const (
	DBNAME = "BSMG"
)

func main() {
	// dbManager init (db 연결)
	dbManager, err := InitDBManager()

	fmt.Printf("%v and %v", dbManager, err)

	dbExist := false
	err = dbManager.DBGorm.IsExistBSMG()
	if err == nil {
		dbExist = true
	}

	if dbExist {
		return
	}

	err = dbManager.DBGorm.CreateDataBase()
	if err != nil {
		// 로그
		fmt.Printf("CreateDataBase Failed . err = %v\n", err)
	}

	dbManager.DBGorm.CreateMemberTable()

}
