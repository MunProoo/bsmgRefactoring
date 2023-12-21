package database

import (
	"BsmgRefactoring/define"
	"fmt"
	"log"
	"strconv"

	"github.com/blue1004jy/gorm"
	_ "github.com/go-sql-driver/mysql"
)

// DB연결 초기화
func (dbm *DBGormMaria) Release() {
	if dbm.DB != nil {
		dbm.DB.Close()
		dbm.DB = nil
	}
}

func (dbm *DBGormMaria) ConnectMariaDB() (err error) {
	dbm.Release()

	// config 파일로 받아오도록 수정
	id := dbm.DBConfig.DatabaseID
	pw := dbm.DBConfig.DatabasePW
	ip := dbm.DBConfig.DatabaseIP
	port := dbm.DBConfig.DatabasePort
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/", id, pw, ip, port)

	// dbm.DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	dbm.DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Printf("ConnectMariaDB %v \n", err)
		// 로그
		return err
	}

	// Open 실패해도 err nil이므로 Ping으로 연결 체크
	// if err = ito.exDbConn.Ping(); err != nil {
	if err = dbm.DB.DB().Ping(); err != nil {
		// out.Printe(out.LogArg{"pn": "ITO", "fn": "connectDB", "text": "connect fail", "err": err})
		dbm.DB.Close()
		dbm.DB = nil
		return err
	}
	// 로그
	// 연결 성공
	return nil
}

func (dbm *DBGormMaria) IsConnected() (err error) {
	err = dbm.DB.DB().Ping()
	if err != nil {
		// out.Printe(out.LogArg{"pn": "ITO", "fn": "connectDB", "text": "connect fail", "err": err})
		dbm.DB.Close()
		dbm.DB = nil
		return err
	}
	return nil
}

// DB 존재여부 확인
func (DBGorm *DBGormMaria) IsExistBSMG() error {
	var count int64
	err := DBGorm.DB.Table("INFORMATION_SCHEMA.SCHEMATA").Where("SCHEMA_NAME = ?", DBNAME).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		// 로그
		// 에러
		err = fmt.Errorf("'BSMG' is not exist")
		return err
	}
	return nil
}

// BSMG DB 생성
func (dbm *DBGormMaria) CreateDataBase() error {
	createQuery := fmt.Sprintf("CREATE DATABASE %s", DBNAME)
	err := dbm.DB.Exec(createQuery).Error
	if err != nil {
		// 로그
		// 연결 실패
		fmt.Println("CreateDataBase Failed . ExecuteQuery = ", createQuery)
		return err
	}

	return nil
}

// BSMG 연결
func (dbm *DBGormMaria) ConnectBSMG() (err error) {
	dbm.Release()
	id := dbm.DBConfig.DatabaseID
	pw := dbm.DBConfig.DatabasePW
	ip := dbm.DBConfig.DatabaseIP
	port := dbm.DBConfig.DatabasePort
	connectionString := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`, id, pw, ip, port, DBNAME)
	dbm.DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	if err = dbm.DB.DB().Ping(); err != nil {
		// out.Printe(out.LogArg{"pn": "ITO", "fn": "connectDB", "text": "connect fail", "err": err})
		dbm.DB.Close()
		dbm.DB = nil
		return err
	}
	return err

}

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

// 사용자 등록
func (dbm *DBGormMaria) InsertMember(member define.BsmgMemberInfo) (err error) {
	nextIdx := dbm.FindMinIdx()
	if nextIdx == 0 {
		// 0은 INSERT 에러, doesn't have a default value
		nextIdx++
	}
	member.Mem_Idx = nextIdx

	err = dbm.DB.Debug().Create(&member).Error
	// queryString := fmt.Sprintf(`INSERT INTO bsmg_member_infos (mem_index, mem_id, mem_password, mem_name,
	// 	mem_rank, mem_part) VALUES (%d,%s,%s,%s,%s,%s)`, member.Mem_Index, member.Mem_ID, member.Mem_Password,
	// 	member.Mem_Name, member.Mem_Rank, member.Mem_Part)
	// err = dbm.DB.Debug().Exec(queryString).Error

	if err != nil {
		log.Printf("%v \n", err)
		return
	}
	return
}

// 가장 작은 user index 번호 return
func (dbm *DBGormMaria) FindMinIdx() int32 {
	queryString := `SELECT MIN(t1.mem_idx + 1) AS nextIdx
	FROM bsmg_member_infos t1
	LEFT JOIN bsmg_member_infos t2 ON t1.mem_idx + 1 = t2.mem_idx
	WHERE t2.mem_idx IS NULL;`

	var nextIdx int32
	err := dbm.DB.Debug().Raw(queryString).Row().Scan(&nextIdx)
	fmt.Printf("%v \n ", nextIdx)
	if err != nil {
		log.Printf("%v \n", err)
	}

	return nextIdx
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

func (dbm *DBGormMaria) InsertDefaultAttr1() {
	attr1List := make([]define.BsmgAttr1Info, 2)
	attr1List[0] = define.BsmgAttr1Info{Attr1_Idx: 1, Attr1_Category: "솔루션"}
	attr1List[1] = define.BsmgAttr1Info{Attr1_Idx: 2, Attr1_Category: "제품"}

	// 슬라이스 사용이 안되다니
	for _, val := range attr1List {
		err := dbm.DB.Create(val).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
}
func (dbm *DBGormMaria) InsertDefaultAttr2() {
	attr2List := make([]define.BsmgAttr2Info, 19)
	attr2List[0] = define.BsmgAttr2Info{Attr2_Idx: 1, Attr1_Idx: 1, Attr2_Name: "출입통제 솔루션"}
	attr2List[1] = define.BsmgAttr2Info{Attr2_Idx: 2, Attr1_Idx: 1, Attr2_Name: "발열감지 솔루션"}
	attr2List[2] = define.BsmgAttr2Info{Attr2_Idx: 3, Attr1_Idx: 1, Attr2_Name: "근태관리 솔루션"}
	attr2List[3] = define.BsmgAttr2Info{Attr2_Idx: 4, Attr1_Idx: 1, Attr2_Name: "식수관리 솔루션"}
	attr2List[4] = define.BsmgAttr2Info{Attr2_Idx: 5, Attr1_Idx: 1, Attr2_Name: "생체인증형 음주측정 솔루션"}
	attr2List[5] = define.BsmgAttr2Info{Attr2_Idx: 6, Attr1_Idx: 1, Attr2_Name: "비대면 방문자 및 행사관리 솔루션"}
	attr2List[6] = define.BsmgAttr2Info{Attr2_Idx: 7, Attr1_Idx: 1, Attr2_Name: "모바일 출입카드 시스템"}
	attr2List[7] = define.BsmgAttr2Info{Attr2_Idx: 8, Attr1_Idx: 1, Attr2_Name: "서버기반 생체인증 솔루션"}

	attr2List[8] = define.BsmgAttr2Info{Attr2_Idx: 9, Attr1_Idx: 2, Attr2_Name: "얼굴인식 장치"}
	attr2List[9] = define.BsmgAttr2Info{Attr2_Idx: 10, Attr1_Idx: 2, Attr2_Name: "홍채인식 장치"}
	attr2List[10] = define.BsmgAttr2Info{Attr2_Idx: 11, Attr1_Idx: 2, Attr2_Name: "지문인식 장치"}
	attr2List[11] = define.BsmgAttr2Info{Attr2_Idx: 12, Attr1_Idx: 2, Attr2_Name: "카드인식 장치"}
	attr2List[12] = define.BsmgAttr2Info{Attr2_Idx: 13, Attr1_Idx: 2, Attr2_Name: "라이브 스캐너"}
	attr2List[13] = define.BsmgAttr2Info{Attr2_Idx: 14, Attr1_Idx: 2, Attr2_Name: "지문 스캐너"}
	attr2List[14] = define.BsmgAttr2Info{Attr2_Idx: 15, Attr1_Idx: 2, Attr2_Name: "도장 스캐너"}
	attr2List[15] = define.BsmgAttr2Info{Attr2_Idx: 16, Attr1_Idx: 2, Attr2_Name: "지문인식 모듈"}
	attr2List[16] = define.BsmgAttr2Info{Attr2_Idx: 17, Attr1_Idx: 2, Attr2_Name: "컨트롤러"}
	attr2List[17] = define.BsmgAttr2Info{Attr2_Idx: 18, Attr1_Idx: 2, Attr2_Name: "발열감지 모듈"}
	attr2List[18] = define.BsmgAttr2Info{Attr2_Idx: 19, Attr1_Idx: 2, Attr2_Name: "단종 제품"}

	for _, val := range attr2List {
		err := dbm.DB.Create(val).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
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
