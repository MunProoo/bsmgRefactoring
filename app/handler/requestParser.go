package handler

import (
	"errors"
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
