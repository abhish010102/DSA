func numberOfPairs(nums []int) []int {
    
    ec := map[int]int{}
    np := 0

    for i:=0;i<len(nums);i++ {
        ec[nums[i]]++;
        if ec[nums[i]]==2{
            np++
            ec[nums[i]]=0
        }
    }

    return []int{np,len(nums)-2*np}
}