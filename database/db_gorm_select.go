package database

import (
	"BsmgRefactoring/define"
	"log"
	"strconv"
)

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
