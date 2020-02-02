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

// func TestLower_bound(t *testing.T) {

// 	arr := []int{1, 2, 3, 4, 5, 8, 29, 99, 432, 432432}
// 	pos := My_lower_bound(arr, 0)
// 	fmt.Println("pos ========", pos, arr[pos])
// 	pos = My_lower_bound(arr, 6)
// 	fmt.Println("pos ========", pos, arr[pos])
// 	pos = My_lower_bound(arr, 100)
// 	fmt.Println("pos ========", pos, arr[pos])
// 	pos = My_lower_bound(arr, 1000)
// 	fmt.Println("pos ========", pos, arr[pos])
// 	pos = My_lower_bound(arr, 10000000)
// 	fmt.Println("pos ========", pos, arr[pos])
// }

// func BenchmarkMy_lower_bound(b *testing.B) {
// 	arr := []int{1, 2, 3, 4, 5, 8, 29, 99, 100, 102, 201, 202, 301, 302, 401, 402, 432, 432432}
// 	My_lower_bound(arr, 77)
// }

func TestRemove(t *testing.T) {
	fmt.Println("TestRemove")
	arr := []int{1, 2, 3, 4, 5}
	arr = arr[1:]
	arr = append(arr, 6)
	a := arr[len(arr)-1]
	fmt.Println(arr, a)
}
