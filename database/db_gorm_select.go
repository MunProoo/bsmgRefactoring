package database

import (
	"BsmgRefactoring/define"
	"log"
	"strconv"
)

func (dbm *DBGormMaria) Login(member *define.BsmgMemberInfo) (err error) {
	dbWhere := dbm.DB.Model(define.BsmgMemberInfo{}).
		Debug().Where("mem_id = ? and mem_password = ? ", member.Mem_ID, member.Mem_Password).
		Find(&member)
	err = dbWhere.Error
	if err != nil {
		log.Printf("Login Failed :  %v \n", err)
		return err
	}

	return
}

func (dbm *DBGormMaria) CheckMemberIDDuplicate(memID string) (isExist bool, err error) {
	var count int
	dbWhere := dbm.DB.Model(define.BsmgMemberInfo{}).
		Debug().Select("Count(0)").Where("mem_id = ?", memID).Count(&count)
	err = dbWhere.Error
	if err != nil {
		log.Printf("CheckMemberIDDuplicate :  %v \n", err)
		return true, err
	}

	if count > 0 {
		isExist = true
	}
	return
}

func (dbm *DBGormMaria) SelectRankList() (rankList []define.BsmgRankInfo, err error) {
	dbWhere := dbm.DB.Model(define.BsmgRankInfo{})
	err = dbWhere.Debug().Find(&rankList).Error
	if err != nil {
		return nil, err
	}
	return
}

func (dbm *DBGormMaria) SelectPartist() (partList []define.BsmgPartInfo, err error) {
	dbWhere := dbm.DB.Model(define.BsmgPartInfo{})
	err = dbWhere.Debug().Find(&partList).Error
	if err != nil {
		return nil, err
	}
	return
}

// 속성1, 속성2 로우 개수 체크
func (dbm *DBGormMaria) Attr1Count() int32 {
	var count int32
	dbWhere := dbm.DB
	dbWhere.Debug().Model(define.BsmgAttr1Info{}).Count(&count)

	return count
}

// 속성1, 속성2 로우 개수 체크
func (dbm *DBGormMaria) Attr2Count() int32 {
	var count int32
	dbWhere := dbm.DB
	dbWhere.Debug().Model(define.BsmgAttr2Info{}).Count(&count)

	return count
}

// 카테고리 Select (트리의 Parent)
func (dbm *DBGormMaria) SelectAttrCategory(count int32) (attrTrees []define.AttrTree, err error) {
	dbWhere := dbm.DB
	attrTrees = make([]define.AttrTree, int(count))

	rows, err := dbWhere.Debug().Model(&define.BsmgAttr1Info{}).Select("Attr1_Category").Rows()
	if err != nil {
		log.Printf("SelectAttrCategory : %v \n", err)
		return nil, err
	}
	defer rows.Close()

	idx := 0
	category := ""
	for rows.Next() {
		err = rows.Scan(&category)
		if err != nil {
			log.Printf("SelectAttrCategory : %v \n", err)
			return nil, err
		}
		attrTrees[idx].Label = category
		attrTrees[idx].Value = strconv.Itoa(idx + 1) // 1부터
		attrTrees[idx].Parent = "0"                  // 트리의 루트
		idx++
	}
	return
}

func (dbm *DBGormMaria) SelectUserList() (userList []define.BsmgMemberInfo, err error) {
	var count int32
	dbWhere := dbm.DB
	dbWhere = dbWhere.Debug().Model(define.BsmgMemberInfo{})
	dbWhere.Count(&count)

	userList = make([]define.BsmgMemberInfo, count)

	err = dbWhere.Select("*").Find(&userList).Error
	if err != nil {
		log.Printf("SelectUserList : %v \n", err)
		return nil, err
	}
	return
}

// 일정 등록을 위해 업무보고 기록 후 idx 송신
func (dbm *DBGormMaria) SelectLatestRptIdx(reporter string) (rptIdx int32, err error) {
	// Pluck 함수는 슬라이스에서만 받을 수 있음
	var rptIdxSlice []int32
	dbWhere := dbm.DB.Model(define.BsmgReportInfo{}).
		Debug().Where("rpt_reporter = ?", reporter).
		Order("rpt_idx DESC").Limit(1).
		Pluck("rpt_idx", &rptIdxSlice)
	err = dbWhere.Error
	if err != nil {
		log.Printf("SelectLatestRptIdx : %v \n", err)
		return 0, err
	}
	rptIdx = rptIdxSlice[0]
	return
}

func (dbm *DBGormMaria) SelectReportList() (rptList []define.BsmgReportInfo, err error) {
	dbWhere := dbm.DB.Model(define.BsmgReportInfo{}).
		Debug().Select("*").
		Joins("INNER JOIN bsmg_member_infos m ON m.mem_id = bsmg_report_infos.rpt_reporter").
		Joins("INNER JOIN bsmg_attr1_infos a ON a.attr1_idx = bsmg_report_infos.rpt_attr1").
		Order("rpt_idx DESC").Limit(6).Offset(0).Find(&rptList)
	err = dbWhere.Error
	if err != nil {
		log.Printf("SelectReportList : %v \n", err)
		return nil, err
	}
	return
}
func (dbm *DBGormMaria) SelectReportInfo(idx int) (rptInfo define.BsmgReportInfo, err error) {
	dbWhere := dbm.DB.Model(define.BsmgReportInfo{})
	dbWhere = dbWhere.Debug().Where("rpt_idx = ?", idx).Find(&rptInfo)
	err = dbWhere.Error
	if err != nil {
		log.Printf("SelectReportInfo : %v \n", err)
		return define.BsmgReportInfo{}, err
	}
	return

}

func (dbm *DBGormMaria) SelectScheduleList(rptIdx int32) (scheduleList []define.BsmgScheduleInfo, err error) {
	dbWhere := dbm.DB.Model(define.BsmgScheduleInfo{}).
		Debug().Select("sc_content").Where("rpt_idx = ?", rptIdx).
		Find(&scheduleList)
	err = dbWhere.Error
	if err != nil {
		log.Printf("SelectScheduleList : %v \n", err)
		return nil, err
	}
	return
}

func (dbm *DBGormMaria) SelectMemberListSearch(searchData define.SearchData) (memberList []define.BsmgMemberInfo, err error) {
	dbWhere := dbm.DB

	ipb := searchData.SearchInput
	switch searchData.SearchCombo {
	case define.SearchAll:
		/*
			SELECT m.mem_id, m.mem_name, m.mem_rank, m.mem_part FROM bsmg_member_infos m
				LEFT OUTER JOIN bsmg_rank_infos r ON r.rank_idx = m.mem_rank
				LEFT OUTER JOIN bsmg_part_infos p ON p.part_idx = m.mem_part
				WHERE m.mem_name like '%%%s%%' or
				  m.mem_part IN (SELECT part_idx FROM bsmg_part_infos p WHERE p.part_name like '%%%s%%') or
				  m.mem_rank IN (SELECT rank_idx FROM bsmg_rank_infos r WHERE r.rank_name like '%%%s%%')
		*/
		dbWhere = dbWhere.Model(define.BsmgMemberInfo{}).Debug().
			Joins("LEFT JOIN bsmg_rank_infos r ON r.rank_idx = mem_rank").
			Joins("LEFT JOIN bsmg_part_infos p ON p.part_idx = mem_part").
			Where("mem_name LIKE ? OR r.rank_name LIKE ? OR p.part_name LIKE ?", "%"+ipb+"%", "%"+ipb+"%", "%"+ipb+"%").
			Find(&memberList)
		err = dbWhere.Error
		if err != nil {
			if err.Error() == "record not found" {
				return nil, nil
			}
			return nil, err
		}
	case define.SearchName:
		dbWhere = dbWhere.Model(define.BsmgMemberInfo{}).Debug().
			Where("mem_name LIKE ?", "%"+ipb+"%").Find(&memberList)
		err = dbWhere.Error
		if err != nil {
			if err.Error() == "record not found" {
				return nil, nil
			}
			return nil, err
		}
	case define.SearchRank:
		dbWhere = dbWhere.Model(define.BsmgMemberInfo{}).Debug().
			Joins("LEFT JOIN bsmg_rank_infos r ON r.rank_idx = mem_rank").
			Where("r.rank_name LIKE ? ", "%"+ipb+"%").
			Find(&memberList)
		err = dbWhere.Error
		if err != nil {
			if err.Error() == "record not found" {
				return nil, nil
			}
			return nil, err
		}
	case define.SearchPart:
		dbWhere = dbWhere.Model(define.BsmgMemberInfo{}).Debug().
			Joins("LEFT JOIN bsmg_part_infos p ON p.part_idx = mem_part").
			Where("p.part_name LIKE ?", "%"+ipb+"%").
			Find(&memberList)
		err = dbWhere.Error
		if err != nil {
			if err.Error() == "record not found" {
				return nil, nil
			}
			return nil, err
		}
	}

	return
}
