package main

import (
	"BsmgRefactoring/define"
)

func parseLoginRequest(parser *FormParser) (result *define.BsmgMemberResult) {
	var err error
	result = &define.BsmgMemberResult{}
	result.MemberInfo = &define.BsmgMemberInfo{}
	result.MemberInfo.Mem_ID, err = parser.getStringValue(0, "mem_id", 0)
	if err != nil {
		result.Result.ResultCode = define.ErrorInvalidParameter
		return result
	}

	result.MemberInfo.Mem_Password, err = parser.getStringValue(0, "mem_pw", 0)
	if err != nil {
		result.Result.ResultCode = define.ErrorInvalidParameter
		return result
	}

	result.Result.ResultCode = 0
	return
}

func parseSearchRequest(parser *FormParser) (search *define.SearchData) {
	var err error
	search = &define.SearchData{}
	search.SearchCombo, err = parser.getStringValue(0, "search_combo", 0)
	if err != nil {
		log.Printf("%v \n ", err)
		return
	}

	search.SearchInput, err = parser.getStringValue(0, "search_input", 0)
	if err != nil {
		log.Printf("%v \n ", err)
		return
	}
	return
}
func parseUserRegistRequest(parser *FormParser) (member *define.BsmgMemberInfo, err error) {
	member = &define.BsmgMemberInfo{}
	member.Mem_ID, err = parser.getStringValue(0, "mem_id", 0)
	if err != nil {
		log.Printf("%v \n ", err)
		return nil, err
	}

	member.Mem_Password, err = parser.getStringValue(0, "mem_pw", 0)
	if err != nil {
		log.Printf("%v \n ", err)
		return nil, err
	}

	member.Mem_Name, err = parser.getStringValue(0, "mem_name", 0)
	if err != nil {
		log.Printf("%v \n ", err)
		return nil, err
	}
	member.Mem_Rank, err = parser.getStringValue(0, "mem_rank", 0)
	if err != nil {
		log.Printf("%v \n ", err)
		return nil, err
	}
	member.Mem_Part, err = parser.getStringValue(0, "mem_part", 0)
	if err != nil {
		log.Printf("%v \n ", err)
		return nil, err
	}
	return member, nil
}
