func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	ele := make([]int,n*n+1)
	sum := 0
	r := -1

	for _, row := range grid {
        for _, val := range row {
			ele[val]++
			sum += val
			if ele[val] == 2 {
				r = val
			}
		}
	}

	sum = n*n*(n*n+1)/2 - sum + r
	return []int{r, sum}
}