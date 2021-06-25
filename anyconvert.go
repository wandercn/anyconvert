/* ====================================================
#   Copyright (C) 2021  All rights reserved.
#
#   Author        : wander
#   Email         : wander@email.cn
#   File Name     : anyconvert.go
#   Last Modified : 2021-06-25 14:56
#   Describe      :
#
# ====================================================*/

package anyconvert

import (
	"log"
	"strconv"
	"strings"
)

/* bson解析同一字段有多种类型数据的转换函数，比如同一个字段即是string类型，也是float64类型*/
func AnyToFloat64(any interface{}) float64 {
	var (
		err error
		f   float64
		ok  bool
		str string
		i32 int32
	)

	filter := func(s string) string {
		s = strings.TrimSpace(s)
		s = strings.ReplaceAll(str, "天", "")
		if strings.HasSuffix(s, "月") {
			switch s {
			case "1个月":
				return "30"
			case "2个月":
				return "60"
			case "3个月":
				return "30"
			default:
				return s
			}
		}
		return s
	}

	if any == nil {
		return 0
	}

	if f, ok = any.(float64); ok {
		return f
	}
	if i32, ok = any.(int32); ok {
		return float64(i32)
	}
	if str, ok = any.(string); ok {
		str = filter(str) // 转换文字到数字字符串
		f, err = strconv.ParseFloat(str, 64)
		if err != nil {
			log.Printf("string: %s convert to float64 failed: %v", str, err)
			return 0
		}
		return f
	}
	log.Printf("any: %v type: %T convert to float64 failed", any, any)
	return 0
}

/* bson解析同一字段有多种类型数据的转换函数，比如同一个字段即是string类型，也是float64类型*/
func AnyToInt(any interface{}) int {
	var (
		err error
		i   int
		ok  bool
		str string
	)
	if any == nil {
		return 0
	}
	if i, ok := any.(int); ok {
		return i
	}
	if str, ok = any.(string); ok {
		str = strings.TrimSpace(str)
		i, err = strconv.Atoi(str)
		if err != nil {
			log.Printf("Atoi failed: %v", err)
			return 0
		}
		return i
	}
	if f, ok := any.(float64); ok {
		return int(f)
	}
	log.Printf("any: %v type: %T convert to int failed", any, any)
	return 0
}
