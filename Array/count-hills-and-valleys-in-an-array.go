func countHillValley(nums []int) int {
	i, ret, pre, cur := 1, 0, nums[0], nums[0]

	for i < len(nums) {
		cur = nums[i]
		if pre != cur {
			break
		}
		i++
	}

	for i < len(nums) {
		nxt := nums[i]
		
		if cur == nxt {
			i++
			continue
		}

		if pre < cur && nxt < cur {
			ret++
		} else if pre > cur && nxt > cur {
			ret++
		}

		i++
		pre = cur
		cur = nxt
	}

	return ret
}