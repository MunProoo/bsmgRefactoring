package database

import (
	"BsmgRefactoring/define"
	"log"
)

// BSMG 직급 테이블 생성
func (dbm *DBGormMaria) CreateRankTable() (err error) {
	log.Println("CreateRankTable")
	exist := dbm.DB.HasTable(&define.BsmgRankInfo{})
	if !exist {
		err := dbm.DB.CreateTable(&define.BsmgRankInfo{}).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
	return
}

// BSMG 부서 테이블 생성
func (dbm *DBGormMaria) CreatePartTable() (err error) {
	log.Println("CreatePartTable")
	exist := dbm.DB.HasTable(&define.BsmgPartInfo{})
	if !exist {
		err := dbm.DB.CreateTable(&define.BsmgPartInfo{}).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
	return
}

// BSMG 멤버 테이블 생성
func (dbm *DBGormMaria) CreateMemberTable() (err error) {
	log.Println("CreateMemberTable")
	exist := dbm.DB.HasTable(&define.BsmgMemberInfo{})
	if !exist {
		err = dbm.DB.CreateTable(&define.BsmgMemberInfo{}).Error
		if err != nil {
			log.Printf("%v \n", err)
			return
		}
		err = dbm.DB.Model(&define.BsmgMemberInfo{}).AddForeignKey("mem_rank", "bsmg_rank_infos(rank_idx)", "NO ACTION", "CASCADE").Error
		if err != nil {
			log.Printf("%v \n", err)
		}
		err = dbm.DB.Model(&define.BsmgMemberInfo{}).AddForeignKey("mem_part", "bsmg_part_infos(part_idx)", "NO ACTION", "CASCADE").Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
	return
}

// BSMG 업무속성(카테고리) 테이블 생성
func (dbm *DBGormMaria) CreateAttr1Table() (err error) {
	log.Println("CreateAttr1Table")
	exist := dbm.DB.HasTable(&define.BsmgAttr1Info{})
	if !exist {
		err := dbm.DB.CreateTable(&define.BsmgAttr1Info{}).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
	return
}

// BSMG 업무속성(이름) 테이블 생성
func (dbm *DBGormMaria) CreateAttr2Table() (err error) {
	log.Println("CreateAttr2Table")
	exist := dbm.DB.HasTable(&define.BsmgAttr2Info{})
	if !exist {
		err := dbm.DB.CreateTable(&define.BsmgAttr2Info{}).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
		err = dbm.DB.Model(&define.BsmgAttr2Info{}).AddForeignKey("attr1_idx", "bsmg_attr1_infos(attr1_idx)", "NO ACTION", "CASCADE").Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
	return
}

// BSMG 일일보고 테이블 생성
func (dbm *DBGormMaria) CreateDailyReportTable() (err error) {
	log.Println("CreateDailyReportTable")
	exist := dbm.DB.HasTable(&define.BsmgReportInfo{})
	if !exist {
		err := dbm.DB.CreateTable(&define.BsmgReportInfo{}).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
		err = dbm.DB.Model(&define.BsmgReportInfo{}).AddForeignKey("rpt_attr1", "bsmg_attr1_infos(attr1_idx)", "NO ACTION", "CASCADE").Error
		if err != nil {
			log.Printf("%v \n", err)
		}
		err = dbm.DB.Model(&define.BsmgReportInfo{}).AddForeignKey("rpt_attr2", "bsmg_attr2_infos(attr2_idx)", "NO ACTION", "CASCADE").Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
	return
}

// BSMG 일정 테이블 생성
func (dbm *DBGormMaria) CreateScheduleTable() (err error) {
	log.Println("CreateScheduleTable")
	exist := dbm.DB.HasTable(&define.BsmgScheduleInfo{})
	if !exist {
		err := dbm.DB.CreateTable(&define.BsmgScheduleInfo{}).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
	return
}

// BSMG 주간보고 테이블 생성
func (dbm *DBGormMaria) CreateWeekReportTable() (err error) {
	log.Println("CreateWeekReportTable")
	exist := dbm.DB.HasTable(&define.BsmgWeekRptInfo{})
	if !exist {
		err := dbm.DB.CreateTable(&define.BsmgWeekRptInfo{}).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
		err = dbm.DB.Model(&define.BsmgWeekRptInfo{}).AddForeignKey("w_rpt_part", "bsmg_part_infos(part_idx)", "NO ACTION", "CASCADE").Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
	return
}
