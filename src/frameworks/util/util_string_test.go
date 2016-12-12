package util

import "testing"

func TestStrToInt64(t *testing.T) {
	iVal64 := int64(10987654321234)
	tVal64, _ := StrToInt64(string("10987654321234"))
	if iVal64 != tVal64 {
		t.Errorf("StrToInt64 error")
	}
}
