package util

import (
	"testing"

	"github.com/bmizerany/assert"
	"strconv"
	"fmt"
	"strings"
)

func TestStrToSlice(t *testing.T) {
	str := "aaa_bbb_ccc"
	sli := StringToSilce(str)
	fmt.Println(sli)
}

func TestStrToInt64(t *testing.T) {
	iVal64 := int64(10987654321234)
	tVal64, _ := StrToInt64(string("10987654321234"))
	if iVal64 != tVal64 {
		t.Errorf("StrToInt64 error")
	}

	s := ".009"
	v1, _ := strconv.ParseFloat(s, 32)
	v2, _ := strconv.ParseFloat(s, 64)
	fmt.Println(v1, v2)
	fmt.Println(int(v2 * 1000))
}

type A struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func TestClass2Json(t *testing.T) {
	var a A
	str := Class2Json(a)
	assert.Equal(t, string(str), "{\"x\":0,\"y\":0}", "class 2 json")
}

func BenchmarkJson(b *testing.B) {

}

func TestStr2BigNum(t *testing.T) {
	assert.Equal(t, Str2BigNum("123").String(), "123", "TestStr2BigNum")
	assert.Equal(t, Str2BigNum("346KK").String(), "346KK", "TestStr2BigNum")
	assert.Equal(t, Str2BigNum("346.9KK").String(), "346.900KK", "TestStr2BigNum")
}

func TestFirstToUpper1(t *testing.T) {
	fmt.Println(strings.ToUpper("hello world 1232PP.f"))
}