package util

func Arr_remove(arr []interface{}, index int) []interface{} {
	if index < 0 || index >= len(arr) {
		return nil
	}

	// 从数组切片中删除元素
	if index < len(arr)-1 { // 中间元素
		arr = append(arr[:index-1], arr[index+1:]...)
	} else if index == 0 {
		// 删除仅有的一个元素
		arr = arr[0:0]
	} else { // 删除的是最后一个元素

		arr = arr[:index-1]
	}
	return arr
}

func My_lower_bound(array []int, key int) int {
	size := len(array)
	if size <= 1 {
		return 0
	}
	first := 0
	last := size - 1
	middle := 0
	pos := 0 //需要用pos记录第一个大于等于key的元素位置

	for first < last {
		middle = (first + last) / 2
		if array[middle] < key { //若中位数的值小于key的值，我们要在右边子序列中查找，这时候pos可能是右边子序列的第一个
			first = middle + 1
			pos = first
		} else {
			last = middle //若中位数的值大于等于key，我们要在左边子序列查找，但有可能middle处就是最终位置，所以我们不移动last,
			pos = last    //而是让first不断逼近last。
		}
	}
	return pos
}
