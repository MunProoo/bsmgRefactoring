package main

import (
	"BsmgRefactoring/define"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 사용자 업무보고 속성트리
func getAttrTreeReq(c echo.Context) error {
	log.Println("getAttrTreeReq")

	var result define.BsmgTreeResult
	result.AttrTreeList = make([]*define.AttrTree, 0)
	result.PartTreeList = make([]*define.PartTree, 0)

	return c.JSON(http.StatusOK, result)
}

// 직급, 부서 정보
func getRankPartReq(c echo.Context) error {
	log.Println("getRankPartReq")

	var result define.BsmgRankPartResult
	result.PartList = make([]*define.BsmgPartInfo, 0)
	part := &define.BsmgPartInfo{
		Part_Idx:  1,
		Part_Name: "부서1",
	}
	result.PartList = append(result.PartList, part)
	result.RankList = make([]*define.BsmgRankInfo, 0)
	rank := &define.BsmgRankInfo{
		Rank_Idx:  1,
		Rank_Name: "랭크1",
	}
	result.RankList = append(result.RankList, rank)
	result.Result.ResultCode = define.Success
	return c.JSON(http.StatusOK, result)
}

// 주간 업무보고 카테고리 정보
func getWeekRptCategoryReq(c echo.Context) error {
	log.Println("getWeekRptCategoryReq")

	var result define.BsmgTreeResult
	result.AttrTreeList = make([]*define.AttrTree, 0)
	result.PartTreeList = make([]*define.PartTree, 0)

	result.Result.ResultCode = define.Success
	return c.JSON(http.StatusOK, result)
}

// 주간 업무보고 카테고리 정보
func getToRptReq(c echo.Context) error {
	log.Println("getToRptReq")

	var result define.BsmgTeamLeaderResult

	value, err := c.FormParams()
	if err != nil {
		log.Printf("%v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	parser := initFormParser(value)
	if parser == nil {
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}
	part_idx, err := parser.getInt32Value(0, "part_idx", 0)
	if err != nil {
		log.Printf("%v \n", err)
		result.Result.ResultCode = define.ErrorInvalidParameter
		return c.JSON(http.StatusOK, result)
	}

	result.Part.PartIdx = part_idx
	result.Result.ResultCode = define.Success

	return c.JSON(http.StatusOK, result)
}
