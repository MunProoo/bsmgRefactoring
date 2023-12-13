package main

import "gorm.io/gorm"

type DBGormMaria struct {
	*gorm.DB
}

func (dbm *DBGormMaria) ConnectDB() (err error) {
	return nil
}
