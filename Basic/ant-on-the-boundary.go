
func returnToBoundaryCount(nums []int) int {
    
    ret := 0

    for i:=1; i<len(nums); i++ {
        nums[i]+=nums[i-1]
        if nums[i]==0 {
            ret++;
        }
    }

    return ret;
}