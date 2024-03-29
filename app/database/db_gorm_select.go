package database

import (
	"BsmgRefactoring/define"
	"errors"
	"strconv"
)

func (dbm *DBGormMaria) SelectMemberInfo(member *define.BsmgMemberInfo) (err error) {
	dbWhere := dbm.DB.Model(define.BsmgMemberInfo{}).
		Debug().Where("mem_id = ?", member.Mem_ID).
		Find(&member)
	err = dbWhere.Error
	if err != nil {
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
		return nil, err
	}
	defer rows.Close()

	idx := 0
	category := ""
	for rows.Next() {
		err = rows.Scan(&category)
		if err != nil {
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
		return 0, err
	}
	rptIdx = rptIdxSlice[0]
	return
}

func (dbm *DBGormMaria) SelectReportList(pageInfo define.PageInfo, searchData define.SearchData) (rptList []define.BsmgReportInfoForWeb, totalCount int32, err error) {
	ipb := searchData.SearchInput
	offset, limit := pageInfo.Offset, pageInfo.Limit

	var reportIncludeName []define.BsmgIncludeNameReport
	dbm.DB.AutoMigrate(define.BsmgIncludeNameReport{})

	dbWhere := dbm.DB.Model(define.BsmgReportInfo{}).Debug().
		Select(`*, m.mem_name as reporter_name`)

	switch searchData.SearchCombo {
	case define.SearchReportAll:
		dbWhere = dbWhere.
			Joins("INNER JOIN bsmg_member_infos m ON m.mem_id = bsmg_report_infos.rpt_reporter").
			Joins("INNER JOIN bsmg_attr1_infos a ON a.attr1_idx = bsmg_report_infos.rpt_attr1").
			Where("rpt_title LIKE ? OR rpt_content LIKE ? OR m.mem_name LIKE ?", "%"+ipb+"%", "%"+ipb+"%", "%"+ipb+"%").
			Order("rpt_idx DESC")
		dbWhere.Count(&totalCount)
		dbWhere.Limit(limit).Offset(offset).Scan(&reportIncludeName)

	case define.SearchReportTitle:
		dbWhere = dbWhere.
			Joins("INNER JOIN bsmg_member_infos m ON m.mem_id = bsmg_report_infos.rpt_reporter").
			Joins("INNER JOIN bsmg_attr1_infos a ON a.attr1_idx = bsmg_report_infos.rpt_attr1").
			Where("rpt_title LIKE ?", "%"+ipb+"%").
			Order("rpt_idx DESC")
		dbWhere.Count(&totalCount)
		dbWhere.Limit(limit).Offset(offset).Scan(&reportIncludeName)
	case define.SearchReportContent:
		dbWhere = dbWhere.
			Joins("INNER JOIN bsmg_member_infos m ON m.mem_id = bsmg_report_infos.rpt_reporter").
			Joins("INNER JOIN bsmg_attr1_infos a ON a.attr1_idx = bsmg_report_infos.rpt_attr1").
			Where("rpt_content LIKE ?", "%"+ipb+"%").
			Order("rpt_idx DESC")
		dbWhere.Count(&totalCount)
		dbWhere.Limit(limit).Offset(offset).Scan(&reportIncludeName)
	case define.SearchReportReporter:
		dbWhere = dbWhere.
			Joins("INNER JOIN bsmg_member_infos m ON m.mem_id = bsmg_report_infos.rpt_reporter").
			Joins("INNER JOIN bsmg_attr1_infos a ON a.attr1_idx = bsmg_report_infos.rpt_attr1").
			Where("m.mem_name LIKE ?", "%"+ipb+"%").
			Order("rpt_idx DESC")
		dbWhere.Count(&totalCount)
		dbWhere.Limit(limit).Offset(offset).Scan(&reportIncludeName)
	default:
		return nil, 0, errors.New("invalid Condition")
	}
	err = dbWhere.Error
	if err != nil {
		return nil, 0, err
	}

	// 보고의 reporter는 ID고 웹에선 이름으로 보여주기 위해..
	// report쪽에 사용자 이름도 추가할까????
	for _, report := range reportIncludeName {
		rptList = append(rptList, report.AllocationInfo())
	}
	return
}

func (dbm *DBGormMaria) SelecAttrSearchReportList(pageInfo define.PageInfo, attrData define.AttrSearchData) (rptList []define.BsmgReportInfoForWeb, totalCount int32, err error) {
	attrValue, attrCategory := attrData.AttrValue, attrData.AttrCategory

	var reportIncludeName []define.BsmgIncludeNameReport

	dbWhere := dbm.DB.Model(define.BsmgReportInfo{}).Debug().Select("*, m.mem_name as reporter_name")

	switch attrCategory {
	case 0: // 솔루션, 제품 등 카테고리로만 검색
		dbWhere = dbWhere.Joins("INNER JOIN bsmg_member_infos m ON m.mem_id = rpt_reporter").
			Where("rpt_attr1 = ?", attrValue)
		dbWhere.Count(&totalCount)
		dbWhere.Limit(pageInfo.Limit).Offset(pageInfo.Offset).Scan(&reportIncludeName)
	case 1: // 솔루션 or 제품의 이름으로 검색
		dbWhere = dbWhere.Joins("INNER JOIN bsmg_member_infos m ON m.mem_id = rpt_reporter").
			Where("rpt_attr2 = ?", attrValue)
		dbWhere.Count(&totalCount)
		dbWhere.Limit(pageInfo.Limit).Offset(pageInfo.Offset).Scan(&reportIncludeName)
	default:
		return nil, 0, errors.New("invalid Condition")
	}

	err = dbWhere.Error
	if err != nil {
		return nil, 0, err
	}

	for _, report := range reportIncludeName {
		rptList = append(rptList, report.AllocationInfo())
	}
	return
}
func (dbm *DBGormMaria) SelectReportInfo(idx int) (reportInfoForWeb define.BsmgReportInfoForWeb, err error) {
	/*
		SELECT r.rpt_idx, m1.mem_name, r.rpt_date, m2.mem_name, r.rpt_ref, r.rpt_title,
			r.rpt_content, r.rpt_etc, ba1.attr1_category, ba2.attr2_name, r.rpt_confirm
			FROM bsmgReport r
			INNER JOIN bsmgMembers m1 ON m1.mem_id = r.rpt_reporter
			INNER JOIN bsmgMembers m2 ON m2.mem_id = r.rpt_toRpt
			WHERE rpt_idx = %d`, rpt_idx
	*/
	var reportIncludeName define.BsmgIncludeNameReport
	dbWhere := dbm.DB.Model(define.BsmgReportInfo{})
	dbWhere = dbWhere.Debug().Select(`*, m1.mem_name as reporter_name, m2.mem_name as to_rpt_name`).
		Joins("INNER JOIN bsmg_member_infos m1 ON m1.mem_id = rpt_reporter").
		Joins("INNER JOIN bsmg_member_infos m2 ON m2.mem_id = rpt_to_rpt").
		// Joins("INNER JOIN bsmg_member_infos m3 ON m3.mem_id = rpt_ref").
		Where("rpt_idx = ?", idx).Scan(&reportIncludeName)
	err = dbWhere.Error
	if err != nil {
		return define.BsmgReportInfoForWeb{}, err
	}
	reportInfoForWeb = reportIncludeName.AllocationInfo()
	return
}

func (dbm *DBGormMaria) SelectScheduleList(rptIdx int32) (scheduleList []define.BsmgScheduleInfo, err error) {
	dbWhere := dbm.DB.Model(define.BsmgScheduleInfo{}).
		Debug().Select("sc_content").Where("rpt_idx = ?", rptIdx).
		Find(&scheduleList)
	err = dbWhere.Error
	if err != nil {
		return nil, err
	}
	return
}

func (dbm *DBGormMaria) SelectMemberListSearch(searchData define.SearchData) (memberList []define.BsmgMemberInfo, err error) {
	dbWhere := dbm.DB

	ipb := searchData.SearchInput
	switch searchData.SearchCombo {
	case define.SearchUserAll:
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
	case define.SearchUserName:
		dbWhere = dbWhere.Model(define.BsmgMemberInfo{}).Debug().
			Where("mem_name LIKE ?", "%"+ipb+"%").Find(&memberList)
		err = dbWhere.Error
		if err != nil {
			if err.Error() == "record not found" {
				return nil, nil
			}
			return nil, err
		}
	case define.SearchUserRank:
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
	case define.SearchUserPart:
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

// 1주일 간의 업무보고 Select
func (dbm *DBGormMaria) SelectReportListAWeek(Mem_ID, bef7d, bef1d string) (reportList []define.BsmgReportInfo, err error) {
	dbWhere := dbm.DB.Model(define.BsmgReportInfo{}).Debug().
		Where("rpt_reporter = ?", Mem_ID).
		Where("CAST(rpt_date as DATETIME) >= ? ", bef7d).
		Where("CAST(rpt_date as DATETIME) <= ?", bef1d).
		Find(&reportList)
	err = dbWhere.Error
	if err != nil {
		return nil, err
	}
	return
}

func (dbm *DBGormMaria) SelectPartLeader(Mem_Part int32) (partLeader string, err error) {
	memberInfo := define.BsmgMemberInfo{}
	dbWhere := dbm.DB.Model(define.BsmgMemberInfo{}).Debug().
		Where("mem_rank = ?", define.PartLeader). // 팀장은 3으로 고정해야할듯
		Where("mem_part = ?", Mem_Part).
		Find(&memberInfo)
	err = dbWhere.Error

	if err != nil {
		if err.Error() == "record not found" {
			partLeader, err = "JJaturi", nil // 팀장 데이터를 안넣었다면 admin으로 고정
			return
		}
		return "", err
	}

	partLeader = memberInfo.Mem_ID
	return
}

func (dbm *DBGormMaria) SelectWeekReportList(pageInfo define.PageInfo, searchData define.SearchData) (rptList []define.BsmgWeekRptInfoForWeb, totalCount int32, err error) {
	ipb := searchData.SearchInput
	offset, limit := pageInfo.Offset, pageInfo.Limit

	var reportIncludeName []define.BsmgIncludeNameWeekReport
	dbm.DB.AutoMigrate(define.BsmgIncludeNameWeekReport{})

	dbWhere := dbm.DB.Model(define.BsmgWeekRptInfo{}).Debug()
	switch searchData.SearchCombo {
	case define.SearchReportAll: // 전체
		dbWhere = dbWhere.Select("*, m1.mem_name as reporter_name, m2.mem_name as to_rpt_name").
			Joins("LEFT JOIN bsmg_member_infos m1 ON m1.mem_id = w_rpt_reporter").
			Joins("LEFT JOIN bsmg_member_infos m2 ON m2.mem_id = w_rpt_to_rpt").
			Where("w_rpt_title LIKE ? OR m1.mem_name LIKE ? OR m2.mem_name LIKE ?", "%"+ipb+"%", "%"+ipb+"%", "%"+ipb+"%")

	case define.SearchReportTitle: // 제목
		dbWhere = dbWhere.Select("*, m1.mem_name as reporter_name, m2.mem_name as to_rpt_name").
			Joins("LEFT JOIN bsmg_member_infos m1 ON m1.mem_id = w_rpt_reporter").
			Joins("LEFT JOIN bsmg_member_infos m2 ON m2.mem_id = w_rpt_to_rpt").
			Where("w_rpt_title LIKE ? ", "%"+ipb+"%")

	case define.SearchWeekReportContent: // 내용
		dbWhere = dbWhere.Select("*, m1.mem_name as reporter_name, m2.mem_name as to_rpt_name").
			Joins("LEFT JOIN bsmg_member_infos m1 ON m1.mem_id = w_rpt_reporter").
			Joins("LEFT JOIN bsmg_member_infos m2 ON m2.mem_id = w_rpt_to_rpt").
			Where("w_rpt_content LIKE ? ", "%"+ipb+"%")

	case define.SearchReportReporter: // 보고대상
		dbWhere = dbWhere.Select("*, m1.mem_name as reporter_name, m2.mem_name as to_rpt_name").
			Joins("LEFT JOIN bsmg_member_infos m1 ON m1.mem_id = w_rpt_reporter").
			Joins("LEFT JOIN bsmg_member_infos m2 ON m2.mem_id = w_rpt_to_rpt").
			Where("m2.mem_name LIKE ? ", "%"+ipb+"%")

	default:
		return nil, 0, errors.New("invalid Condition")
	}

	dbWhere = dbWhere.Order("w_rpt_idx DESC")
	dbWhere.Count(&totalCount)
	dbWhere.Limit(limit).Offset(offset).Scan(&reportIncludeName)
	err = dbWhere.Error
	if err != nil {
		return nil, 0, err
	}

	// 보고의 reporter는 ID고 웹에선 이름으로 보여주기 위해..
	// report쪽에 사용자 이름도 추가할까????
	for _, report := range reportIncludeName {
		rptList = append(rptList, report.AllocationInfo())
	}

	return
}

func (dbm *DBGormMaria) SelectWeekReportCategorySearch(pageInfo define.PageInfo, partIdx int) (rptList []define.BsmgWeekRptInfoForWeb, totalCount int32, err error) {
	offset, limit := pageInfo.Offset, pageInfo.Limit

	var reportIncludeName []define.BsmgIncludeNameWeekReport
	dbm.DB.AutoMigrate(define.BsmgIncludeNameWeekReport{})

	dbWhere := dbm.DB.Model(define.BsmgWeekRptInfo{}).Debug().
		Select("*, m1.mem_name as reporter_name, m2.mem_name as to_rpt_name").
		Joins("LEFT JOIN bsmg_member_infos m1 ON m1.mem_id = w_rpt_reporter").
		Joins("LEFT JOIN bsmg_member_infos m2 ON m2.mem_id = w_rpt_to_rpt")

	if partIdx != define.AllCategory { // 1개의 부서를 특정한 경우
		dbWhere = dbWhere.Where("w_rpt_part = ?", partIdx)
	}

	dbWhere.Count(&totalCount)
	dbWhere.Order("w_rpt_idx DESC").Limit(limit).Offset(offset).Scan(&reportIncludeName)
	err = dbWhere.Error
	if err != nil {
		return nil, 0, err
	}

	// 보고의 reporter는 ID고 웹에선 이름으로 보여주기 위해..
	// report쪽에 사용자 이름도 추가할까????
	for _, report := range reportIncludeName {
		rptList = append(rptList, report.AllocationInfo())
	}

	return
}

func (dbm *DBGormMaria) SelectWeekReportInfo(wRptIdx int) (rptInfo define.BsmgWeekRptInfoForWeb, err error) {
	var reportIncludeName define.BsmgIncludeNameWeekReport
	dbm.DB.AutoMigrate(define.BsmgIncludeNameWeekReport{})

	dbWhere := dbm.DB.Model(define.BsmgWeekRptInfo{}).Debug().
		Select("*, m1.mem_name as reporter_name, m2.mem_name as to_rpt_name").
		Joins("LEFT JOIN bsmg_member_infos m1 ON m1.mem_id = w_rpt_reporter").
		Joins("LEFT JOIN bsmg_member_infos m2 ON m2.mem_id = w_rpt_to_rpt").
		Where("w_rpt_idx = ?", wRptIdx).
		Scan(&reportIncludeName)
	err = dbWhere.Error
	if err != nil {
		return define.BsmgWeekRptInfoForWeb{}, err
	}

	rptInfo = reportIncludeName.AllocationInfo()

	return
}

func (dbm *DBGormMaria) SelectAttr1List() (attr1List []define.BsmgAttr1Info, err error) {
	dbWhere := dbm.DB.Model(define.BsmgAttr1Info{}).Debug().
		Find(&attr1List)
	err = dbWhere.Error
	if err != nil {
		return nil, err
	}

	return
}
