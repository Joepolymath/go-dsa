package leet

func containsDuplicate(nums []int) bool {
    hashMap := make(map[int]int)
    for _, element := range nums {
        if _, found := hashMap[element]; found {
            return true
        }
        hashMap[element] = 1
    }
    return false
}