package main

import (
	"BsmgRefactoring/define"
	"errors"
	"strconv"
	"strings"
)

// 쿼리를 파싱 (호출해주는 역할)
func ParseQuery(query string) (Values, error) {
	m := make(Values)
	err := parseQuery(m, query)
	return m, err
}

// 쿼리를 파싱 (실제 알고리즘)
func parseQuery(m Values, query string) (err error) {
	for query != "" {
		key := query
		if i := strings.IndexAny(key, "&;"); i >= 0 {
			key, query = key[:i], key[i+1:]
		} else {
			query = ""
		}
		if key == "" {
			continue
		}
		value := ""
		if i := strings.Index(key, "="); i >= 0 {
			key, value = key[:i], key[i+1:]
		}
		m[key] = append(m[key], value)
	}
	return err
}

type Values map[string][]string // 흠..?

func (v Values) Get(key string) string {
	if v == nil {
		return ""
	}
	vs := v[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

// -------------------------------------------------------페이징
func getPageInfo(query string) (PageInfo define.PageInfo, err error) {
	values, err := ParseQuery(query)
	if err != nil {
		return
	}

	var tempInt int64

	tempInt, err = strconv.ParseInt(values.Get("limit"), 10, 32)
	if err != nil {
		return
	}
	PageInfo.Limit = int32(tempInt)
	if PageInfo.Limit < 1 {
		err = errors.New("Limit must bigger than 0")
		return
	}

	tempInt, err = strconv.ParseInt(values.Get("offset"), 10, 32)
	if err != nil {
		return
	}
	PageInfo.Offset = int32(tempInt)

	return
}
