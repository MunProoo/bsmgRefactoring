package main

import (
	"BsmgRefactoring/define"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

// func connectToMariaDB() (*gorm.DB, error) {
// 	// out.Printi(out.LogArg{"pn": "ITO", "fn": "connectDB", "text": "connectDB start"})

// 	var DBConfig define.DBConfig
// 	DBConfig.DatabaseIP = "127.0.0.1"
// 	DBConfig.DatabaseID = "root"
// 	DBConfig.DatabasePW = "0000"
// 	DBConfig.DatabasePort = "3306"
// 	DBConfig.DatabaseName = ""

// 	// var connectionString string
// 	// connectionString = fmt.Sprintf("%s@%s:%s@tcp(%s:%s)")

// }

func createDataBase() (*gorm.DB, error) {

	// 메모리에 저장하자 AES 256해서
	var DBConfig define.DBConfig
	DBConfig.DatabaseIP = "127.0.0.1"
	DBConfig.DatabaseID = "root"
	DBConfig.DatabasePW = "0000"
	DBConfig.DatabasePort = "3306"
	DBConfig.DatabaseName = ""

	connectionString := "root:12345@tcp(127.0.0.1:3306)/"

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		// 로그남기기
		return nil, err
	}

	return db, nil

}
