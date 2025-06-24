func numberGame(nums []int) []int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i += 2 {
		t := nums[i-1]
		nums[i-1] = nums[i]
		nums[i] = t
	}
	return nums
}