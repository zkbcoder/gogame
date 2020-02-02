package util

import (
	"encoding/json"
	"strconv"
	"strings"
	"unicode"
	"regexp"
	"fmt"
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

// 字符串json转成map
// result 是函数内的临时变量，作为返回值可以直接返回调用层
func Json2map(str []byte) (s map[string]interface{}, err error) {
	var result map[string]interface{}
	if err := json.Unmarshal(str, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// 结构体转json异常返回空字符串(类型是:[]byte)
func Class2Json(v interface{}) []byte {
	res, err := json.Marshal(v)
	if err != nil {
		return []byte("")
	}
	return res
}


// 字符串转大数
func Str2BigNum(str string) *BigNum {
	str = strings.ToUpper(str)
	reg := regexp.MustCompile("[A-Z]+")
	strs := reg.FindAllString(str, -1)
	unit := ""
	if len(strs) > 0 {
		unit = strs[0] // 单位
	}
	unitPost := GetPost(unit)

	str = reg.ReplaceAllString(str, "") // 去掉单位
	strNums := strings.Split(str, ".")
	high,_ := strconv.Atoi(strNums[0])
	low := 0
	if len(strNums) > 1 {
		lowStr := fmt.Sprintf(".%s", strNums[1])
		fTemp, _ := strconv.ParseFloat(lowStr, 64)
		low = int(fTemp * 1000)
	}

	bigNum := BigNum{}
	bigNum.AddVal(unitPost, int16(high))
	if unitPost > 0 {
		bigNum.AddVal(unitPost-1, int16(low))
	}
	return &bigNum
}


//  int数组转字符串
func SilceIntToString(sil []int) (ret string) {
	ret = ""
	for i := 0; i < len(sil); i++ {
		if i >= 1 {
			ret += "_"
		}
		ret += fmt.Sprintf("%d", sil[i])
	}
	return
}

// 字符串转int数组
func StringToSilceInt(str string) (sil []int) {
	sil = make([]int, 0)
	if str == "" {
		return
	}
	strs := strings.Split(str, "_")
	for i := 0; i < len(strs); i++ {
		num, _ := strconv.Atoi(strs[i])
		sil = append(sil, num)
	}
	return
}

// 字符串转string数组
func StringToSilce(str string) (sil []string) {
	sil = make([]string, 0)
	if str == "" {
		return
	}
	sil = strings.Split(str, "_")
	return
}
