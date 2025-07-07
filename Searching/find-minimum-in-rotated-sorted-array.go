func findMin(nums []int) int {

	i, j := 0, len(nums)-1

	for j >= i {

		if nums[i] < nums[j] {
			return nums[i]
		}

		m := (i + j) / 2

		if nums[m] >= nums[0] {
			i = m + 1
		} else if nums[m] > nums[m-1] {
			j = m - 1
		} else {
			return nums[m]
		}
	}

	if i >= len(nums) {
		return nums[0]
	}
	return nums[i]
}