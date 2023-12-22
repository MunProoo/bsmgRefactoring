package database

import (
	"BsmgRefactoring/define"
	"log"
)

func (dbm *DBGormMaria) UpdateUser(setVal map[string]interface{}, memID string) error {
	dbWhere := dbm.DB.Model(define.BsmgMemberInfo{})

	err := dbWhere.Debug().Where("mem_id = ?", memID).Update(setVal).Error
	if err != nil {
		log.Printf("UpdateUser : %v \n", err)
		return err
	}
	return err
}
