package leet

func twoSum(nums []int, target int) []int {
    memoryMap := make(map[int]int)

    for i, elem := range nums {
        if val, found := memoryMap[target - elem]; found {
            return []int{val, i}
        }
        memoryMap[elem] = i
    }
    return nil
}