package database

import (
	"BsmgRefactoring/define"
	"log"
)

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

func (dbm *DBGormMaria) InsertDefaultRank() {
	rankList := make([]define.BsmgRankInfo, 4)
	rankList[0] = define.BsmgRankInfo{Rank_Idx: 1, Rank_Name: "연구소장"}
	rankList[1] = define.BsmgRankInfo{Rank_Idx: 2, Rank_Name: "부소장"}
	rankList[2] = define.BsmgRankInfo{Rank_Idx: 3, Rank_Name: "팀장"}
	rankList[3] = define.BsmgRankInfo{Rank_Idx: 4, Rank_Name: "Pro"}

	// 슬라이스 사용이 안되다니
	for _, val := range rankList {
		err := dbm.DB.Create(val).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
}
func (dbm *DBGormMaria) InsertDefaultPart() {
	partList := make([]define.BsmgPartInfo, 11)
	partList[0] = define.BsmgPartInfo{Part_Idx: 1, Part_Name: "연구소"}
	partList[1] = define.BsmgPartInfo{Part_Idx: 2, Part_Name: "SW1팀"}
	partList[2] = define.BsmgPartInfo{Part_Idx: 3, Part_Name: "SW2팀"}
	partList[3] = define.BsmgPartInfo{Part_Idx: 4, Part_Name: "FW1팀"}
	partList[4] = define.BsmgPartInfo{Part_Idx: 5, Part_Name: "FW2팀"}
	partList[5] = define.BsmgPartInfo{Part_Idx: 6, Part_Name: "HW1팀"}
	partList[6] = define.BsmgPartInfo{Part_Idx: 7, Part_Name: "HW2팀"}
	partList[7] = define.BsmgPartInfo{Part_Idx: 8, Part_Name: "Mobile팀"}
	partList[8] = define.BsmgPartInfo{Part_Idx: 9, Part_Name: "디자인팀"}
	partList[9] = define.BsmgPartInfo{Part_Idx: 10, Part_Name: "광학기구팀"}
	partList[10] = define.BsmgPartInfo{Part_Idx: 11, Part_Name: "연구관리팀"}

	// 슬라이스 사용이 안되다니
	for _, val := range partList {
		err := dbm.DB.Create(val).Error
		if err != nil {
			log.Printf("%v \n", err)
		}
	}
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
	if err != nil {
		log.Printf("%v \n", err)
		return
	}
	return
}

func (dbm *DBGormMaria) InsertDailyReport(report define.BsmgReportInfo) (err error) {
	dbWhere := dbm.DB.Model(define.BsmgReportInfo{})
	err = dbWhere.Debug().Create(&report).Error
	if err != nil {
		log.Printf("CreateDailyReport : %v \n", err)
		return
	}

	return
}

func (dbm *DBGormMaria) InsertSchedule(schedule define.BsmgScheduleInfo) (err error) {
	dbWhere := dbm.DB.Model(define.BsmgScheduleInfo{})
	err = dbWhere.Debug().Create(&schedule).Error
	if err != nil {
		log.Printf("InsertSchedule : %v \n", err)
		return
	}

	return
}
