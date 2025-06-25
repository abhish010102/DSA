func mostFrequent(nums []int, key int) int {
	ret := 0
	maxcount := 0
	elementcount := map[int]int{}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			elementcount[nums[i+1]]++
			if elementcount[nums[i+1]] > maxcount {
				maxcount = elementcount[nums[i+1]]
				ret = nums[i+1]
			}
		}
	}
	return ret
}