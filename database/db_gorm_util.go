package database

import (
	"BsmgRefactoring/define"
	"log"
	"strconv"
)

// 가장 작은 user index 번호 return
func (dbm *DBGormMaria) FindMinIdx() int32 {
	queryString := `SELECT MIN(t1.mem_idx + 1) AS nextIdx
	FROM bsmg_member_infos t1
	LEFT JOIN bsmg_member_infos t2 ON t1.mem_idx + 1 = t2.mem_idx
	WHERE t2.mem_idx IS NULL;`

	var nextIdx int32
	err := dbm.DB.Debug().Raw(queryString).Row().Scan(&nextIdx)
	if err != nil {
		log.Printf("%v \n", err)
	}

	return nextIdx
}

// 클라이언트 단에서 작업하는게 서버의 부하를 줄이겠지만, 클라이언트 수정이 여건상 힘드므로 부득이하게..
// 웹앱의 트리구조 포맷에 맞춰 업무속성을 서버에서 트리로 제작
func (dbm *DBGormMaria) MakeAttrTree() (attrTrees []define.AttrTree, err error) {
	count1 := dbm.Attr1Count()
	attrTrees, err = dbm.SelectAttrCategory(count1)
	if attrTrees == nil || err != nil {
		log.Printf("SelectAttrTree : %v \n", err)
		return nil, err
	}

	attr2Trees, err := dbm.AttrTreeSetNameAndValue(dbm.Attr2Count())

	attrTrees = append(attrTrees, attr2Trees...)
	return attrTrees, err
}

// 업무속성을 parent에 맞게 트리구조 작성
func (dbm *DBGormMaria) AttrTreeSetNameAndValue(count int32) (attrTrees []define.AttrTree, err error) {
	attrTrees = make([]define.AttrTree, count)

	dbWhere := dbm.DB
	// 트리구조를 위해 솔루션이면 1- , 제품이면 2-로 value 작성
	queryString := `SELECT ba2.attr2_name , 
	CASE WHEN ba2.attr1_idx = 1 THEN '1-' 
		 WHEN ba2.attr1_idx = 2 THEN '2-' 
		 ELSE 'N/A' 
	END AS value, 
	ba2.attr1_idx, ba2.attr2_idx 
	FROM bsmg_attr2_infos ba2;`
	rows, err := dbWhere.Debug().Raw(queryString).Rows()
	if err != nil {
		log.Printf("AttrTreeSetNameAndValue : %v \n", err)
		return nil, err
	}
	defer rows.Close()

	// var attr2Infos struct {
	var (
		attr2_name string
		value      string
		attr1_idx  int32
		attr2_idx  int32
	)
	// }

	var idx int
	for rows.Next() {
		err = rows.Scan(&attr2_name, &value, &attr1_idx, &attr2_idx)
		if err != nil {
			log.Printf("AttrTreeSetNameAndValue : %v \n", err)
			return nil, err
		}
		attrTrees[idx].Label = attr2_name
		attrTrees[idx].Value = value + strconv.Itoa(int(attr2_idx))
		if attr1_idx == 1 { // 업무 속성이 솔루션이면 parent를 1로
			attrTrees[idx].Parent = "1"
		} else { // 업무 속성이 제품이면 parent 2
			attrTrees[idx].Parent = "2"
		}
		idx++
	}

	return
}

// 클라이언트 단에서 작업하는게 서버의 부하를 줄이겠지만, 클라이언트 수정이 여건상 힘드므로 부득이하게..
// 웹앱의 트리구조 포맷에 맞춰 업무속성을 서버에서 트리로 제작
func (dbm *DBGormMaria) MakePartTree() (partTreeList []define.PartTree, err error) {
	partList, err := dbm.SelectPartist()
	partTreeList = make([]define.PartTree, len(partList)+1)

	// 트리 최상위 부모 설정
	partTreeList[0].Label = define.WeekCategoryName
	partTreeList[0].Value = "1"
	partTreeList[0].Parent = "0"

	for idx, part := range partList {
		partTreeList[idx+1].Label = part.Part_Name
		partTreeList[idx+1].Value = "1-" + strconv.Itoa(int(part.Part_Idx))
		partTreeList[idx+1].Parent = "1"
	}

	return

}
