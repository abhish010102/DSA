func findLucky(arr []int) int {
	mp := map[int]int{}
	ret := -1

	for i := 0; i < len(arr); i++ {
		mp[arr[i]]++
	}

	for k, v := range mp {
		if k == v && ret < k {
			ret = k
		}
	}

	return ret
}