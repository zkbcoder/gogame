package util

import (
	"strconv"
	"strings"
	"unicode"
)

// 字符串首字母变大写
func FirstToUpper(str string) string {
	if len(str) == 0 {
		return ""
	}
	r := []rune(str)
	r[0] = unicode.ToUpper(r[0])
	str = string(r)
	return str
}

// 将字符串小写蛇体转驼峰 例如: user_level--> UserLevel
func StringToClass(str string) string {
	strs := strings.Split(str, "_")
	out := ""
	for _, v := range strs {
		out += FirstToUpper(v)
	}
	return out
}

// 按规则转化成redis key 例如: game_attr--> AttrEntity
// [如果没有game_的情况] 例如: attr-->AttrEntity
func StringToEntity(str string) string {
	strNew := str
	if (len(str) > 5) && (str[0:5] == "game_") {
		strNew = str[5:]
	}
	strs := strings.Split(strNew, "_")
	out := ""
	for _, v := range strs {
		out += FirstToUpper(v)
	}

	out += "Entity"
	return out
}

func StrToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 0)
}

// 字符串数组转化成map 并且过滤重复key
func StringArrToMap(arr []string) map[string]bool {
	out := make(map[string]bool)
	for _, key := range arr {
		_, ok := out[key]
		if !ok {
			out[key] = true
		}
	}
	return out
}
