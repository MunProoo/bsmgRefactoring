package main

import (
	"BsmgRefactoring/database"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// func initConfig() (config *define.DBConfig, err error) {
// 	curPath, err := filepath.Abs(filepath.Dir(os.Args[0]))

// 	if _, err := os.Stat(curPath + "/dbConfig.json"); os.IsNotExist(err) {
// 		file, err := os.OpenFile(curPath+"/dbConfig.json", os.O_CREATE|os.O_RDWR, 0666)
// 		if err != nil {

// 		}
// 	}
// 	return
// }

type ServerProcessor struct {
	dbManager database.DatabaseManager
}

func (server *ServerProcessor) ConnectDataBase() (err error) {
	err = server.dbManager.InitDBManager()
	if err != nil {
		// 로그
		log.Printf("InitDBManager Failed . err = %v\n", err)
		return
	}
	return
}
