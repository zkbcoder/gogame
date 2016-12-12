package util

import (
	"fmt"
	"testing"
)

//func TestRemove(t *testing.T) {

//	arr := []int{1, 2, 3, 4, 5}
//	fmt.Println("before ========", arr)
//	arr = Arr_remove([]interface{}(arr), 3)
//	fmt.Println("after =========", arr)
//}

func TestLower_bound(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 8, 29, 99, 432, 432432}
	pos := My_lower_bound(arr, 0)
	fmt.Println("pos ========", pos, arr[pos])
	pos = My_lower_bound(arr, 6)
	fmt.Println("pos ========", pos, arr[pos])
	pos = My_lower_bound(arr, 100)
	fmt.Println("pos ========", pos, arr[pos])
	pos = My_lower_bound(arr, 1000)
	fmt.Println("pos ========", pos, arr[pos])
	pos = My_lower_bound(arr, 10000000)
	fmt.Println("pos ========", pos, arr[pos])

}
