package database

import (
	"BsmgRefactoring/define"
	"log"
)

func (dbm *DBGormMaria) UpdateUser(member define.BsmgMemberInfo) error {
	memID := member.Mem_ID

	setVal := make(map[string]interface{})
	setVal["mem_name"] = member.Mem_Name
	setVal["mem_rank"] = member.Mem_Rank
	setVal["mem_part"] = member.Mem_Part

	dbWhere := dbm.DB.Model(define.BsmgMemberInfo{})

	err := dbWhere.Debug().Where("mem_id = ?", memID).Update(setVal).Error
	if err != nil {
		log.Printf("UpdateUser : %v \n", err)
		return err
	}
	return err
}

func (dbm *DBGormMaria) UpdateReportInfo(report define.BsmgReportInfo) (err error) {
	rptIdx := report.Rpt_Idx
	setVal := make(map[string]interface{})
	setVal["rpt_title"] = report.Rpt_title
	setVal["rpt_content"] = report.Rpt_content
	setVal["rpt_etc"] = report.Rpt_etc
	setVal["rpt_attr1"] = report.Rpt_attr1
	setVal["rpt_attr2"] = report.Rpt_attr2

	dbWhere := dbm.DB.Model(define.BsmgReportInfo{}).
		Debug().Where("rpt_idx = ?", rptIdx).Update(setVal)
	err = dbWhere.Error
	if err != nil {
		log.Printf("UpdateReportInfo : %v \n", err)
		return err
	}

	return
}
