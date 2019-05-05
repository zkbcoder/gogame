package util

import (
	"fmt"
	"strconv"
	"strings"
)

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
