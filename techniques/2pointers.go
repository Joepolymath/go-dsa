package techniques

func TwoSumWith2Pointers(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	for nums[left] + nums[right] != target {
		if nums[left] + nums[right] < target {
			left++
			continue
		}
		if nums[left] + nums[right] > target {
			right--
		}

		if left >= right && nums[left] + nums[right] != target {
			return nil
		}
	}
	return []int{left, right}
}