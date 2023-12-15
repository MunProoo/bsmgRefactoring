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
