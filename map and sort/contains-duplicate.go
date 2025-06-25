// func containsDuplicate(nums []int) bool {
//     sort.Ints(nums)

//	    for i:=1;i<len(nums);i++ {
//	        if nums[i-1]==nums[i]{
//	            return true
//	        }
//	    }
//	    return false
//	}
func containsDuplicate(nums []int) bool {
	ele := map[int]int{}
	for i := 0; i < len(nums); i++ {
		ele[nums[i]]++
		if ele[nums[i]] == 2 {
			return true
		}
	}
	return false
}