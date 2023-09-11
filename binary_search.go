package main

func BinarySearch(dataList []int, key int) bool {
	low := 0
	high := len(dataList) - 1

	for low <= high {
		median := (low + high)/2

		if dataList[median] < key {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(dataList) || dataList[low] != key {
		return false
	}
	return true
}