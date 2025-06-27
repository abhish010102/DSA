func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {  
            if nums[nums[i]-1]>0{
			    nums[nums[i]-1] *= -1
            }
		} else if nums[-nums[i]-1]>0{
			nums[-nums[i]-1] *= -1
		}
	}

	ret := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ret = append(ret, i+1)
		}
	}
	return ret
}