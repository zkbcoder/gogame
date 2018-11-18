package util

import (
	"fmt"
	"testing"
)

type T1 struct {
	A    int
	Arry []int
}

// ！！！深拷贝 结构体里面的数组拷贝和c＋＋指针一样，需要做深拷贝不然只是指针赋值而已
func TestDeepCopy(t *testing.T) {
	a := T1{1, make([]int, 2)}
	a.Arry[0] = 9

	var b T1
	DeepCopy(&b, &a)
	b.A = 3
	b.Arry[0] = 77
	fmt.Println(a)
	fmt.Println(b)
}
