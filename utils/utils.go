package utils

import (
	"strconv"
	"strings"
)

// 트리에서 가져온 value를 다시 attr2_idx로
func GetAttr2Idx(rpt_attr2String string) int32 {
	idx := strings.Index(rpt_attr2String, "-")
	rpt_attr2String = rpt_attr2String[idx+1:]
	rpt_attr2, err := strconv.ParseInt(rpt_attr2String, 10, 32)
	if err != nil {
		return 0
	}

	return int32(rpt_attr2)
}
