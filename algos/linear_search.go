package algos

func LinearSearch(dataList []int, key int) bool {
	for i := range dataList {
		if dataList[i] == key {
			return true
		}
	}
	return false
}