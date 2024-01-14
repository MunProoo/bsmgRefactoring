package database

import (
	"BsmgRefactoring/define"
)

func (dbm *DBGormMaria) DeleteSchedule(rptIdx int32) (err error) {

	dbWhere := dbm.DB.Debug().
		Where("rpt_idx = ?", rptIdx).Delete(define.BsmgScheduleInfo{})
	err = dbWhere.Error
	if err != nil {
		return err
	}
	return
}

func (dbm *DBGormMaria) DeleteReport(rptIdx int32) (err error) {
	dbWhere := dbm.DB.Debug().
		Where("rpt_idx = ? ", rptIdx).Delete(define.BsmgReportInfo{})
	err = dbWhere.Error
	if err != nil {
		return err
	}
	return
}

func (dbm *DBGormMaria) DeleteMember(memID string) (err error) {
	dbWhere := dbm.DB.Debug().
		Where("mem_id = ?", memID).Delete(define.BsmgMemberInfo{})
	err = dbWhere.Error
	if err != nil {
		return err
	}
	return
}

func (dbm *DBGormMaria) DeleteWeekReport(wRptIdx int) (err error) {
	dbWhere := dbm.DB.Debug().
		Where("w_rpt_idx = ?", wRptIdx).Delete(define.BsmgWeekRptInfo{})
	err = dbWhere.Error
	if err != nil {
		return err
	}
	return
}
