package handler

import (
	"BsmgRefactoring/define"
	"errors"
	"log"
	"net/url"
	"strconv"
)

// eXBuilder에서 받은 Form 데이터를 parsing 하기 위한 파일
// API도 가능하게 하려면 application/json으로 보내는 것도 작업 필요

// application/x-www-form-urlencoded

type FormParser struct {
	Form      url.Values
	prefix    []string
	prefixLen int
}

// http request data를 서버에서 받기 편하게 eXBuilder format에 맞춰 parser 제작
func InitFormParser(form url.Values) (parser *FormParser) {
	parser = &FormParser{
		Form: form,
	}

	parser.prefix = form["@d#"]
	parser.prefixLen = len(parser.prefix)
	return
}

// keyword에 맞는 request data parsing -------------------------------------------------------------------
func (parser *FormParser) GetInt32Value(index int, keyword string, subIndex int) (value int32, err error) {
	if parser.prefixLen != 0 && index >= parser.prefixLen {
		err = errors.New("prefix index out of range")
		return
	}

	var key string
	if parser.prefixLen > 0 {
		key = parser.prefix[index] + keyword
	} else {
		key = keyword
	}

	if len(parser.Form[key]) > 0 {
		var data int64
		data, err = strconv.ParseInt(parser.Form[key][subIndex], 10, 32)
		if err != nil {
			return
		}
		value = int32(data)
	}
	return
}

func (parser *FormParser) GetInt64Value(index int, keyword string, subIndex int) (value int64, err error) {
	if parser.prefixLen != 0 && index >= parser.prefixLen {
		err = errors.New("prefix index out of range")
		return
	}
	var key string
	if parser.prefixLen > 0 {
		key = parser.prefix[index] + keyword
	} else {
		key = keyword
	}

	if len(parser.Form[key]) > 0 {
		value, err = strconv.ParseInt(parser.Form[key][subIndex], 10, 64)
		if err != nil {
			return
		}
	}
	return
}

func (parser *FormParser) getStringValue(index int, keyword string, subIndex int) (value string, err error) {
	if parser.prefixLen != 0 && index >= parser.prefixLen {
		err = errors.New("prefix index out of range")
		return
	}
	var key string
	if parser.prefixLen > 0 {
		key = parser.prefix[index] + keyword
	} else {
		key = keyword
	}

	if len(parser.Form[key]) > 0 {
		value = parser.Form[key][subIndex]
	}
	return
}

func (parser *FormParser) getStringArray(index int, keyword string) (values []string, err error) {
	if parser.prefixLen != 0 && index >= parser.prefixLen {
		err = errors.New("prefix index out of range")
		return
	}
	var key string
	if parser.prefixLen > 0 {
		key = parser.prefix[index] + keyword
	} else {
		key = keyword
	}

	if len(parser.Form[key]) > 0 {
		return parser.Form[key], nil
	}
	return
}

func (parser *FormParser) getValueCount(index int, keyword string) (count int, err error) {
	if parser.prefixLen != 0 && index >= parser.prefixLen {
		err = errors.New("prefix index out of range")
		return
	}
	var key string
	if parser.prefixLen > 0 {
		key = parser.prefix[index] + keyword
	} else {
		key = keyword
	}

	count = len(parser.Form[key])
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
	member.Mem_Rank, err = parser.GetInt32Value(0, "mem_rank", 0)
	if err != nil {
		log.Printf("%v \n ", err)
		return nil, err
	}
	member.Mem_Part, err = parser.GetInt32Value(0, "mem_part", 0)
	if err != nil {
		log.Printf("%v \n ", err)
		return nil, err
	}
	return member, nil
}
