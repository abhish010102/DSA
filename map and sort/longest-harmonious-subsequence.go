// use sort
func findLHS(nums []int) int {
	
	sort.Ints(nums)
    i, j, ret := 0, 1, 0 

	for j < len(nums) {
    	for nums[j]-nums[i] > 1 {
			i++;
		}
		if nums[j]-nums[i]==1 && ret < (j - i+1) {
			ret = j - i+1
		}	
    	j++
    }

	return ret
}

//use map
func findLHS(nums []int) int {
    l1 := map[int](int){}
    j, ret := 0, 0

    for j<len(nums){
        l1[nums[j]]++
        j++;
    }

    for key, value := range l1 {
        if _, check := l1[key+1]; check{
            if ret<(value + l1[key+1]){
                ret=(value + l1[key+1])
            }
        }
    }
    return ret
}