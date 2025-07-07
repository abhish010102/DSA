func searchRange(nums []int, target int) []int {
	i, j := 0, len(nums) - 1

	for j >= i {
		m := (i + j) / 2
		if nums[m] == target {
			j = m
			break
		} else if nums[m] > target {
			j = m - 1
		} else {
			i = m + 1
		}
	}

	if j < i {
		return []int{-1, -1}
	}

	i = j
	for i>0 && nums[i-1] == nums[j] {
		i--
	}

	for j<len(nums)-1 && nums[j+1] == nums[i] {
		j++
	}
    
	return []int{i, j}
}