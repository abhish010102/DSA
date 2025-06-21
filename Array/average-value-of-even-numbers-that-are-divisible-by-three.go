func averageValue(nums []int) int {
    avg := 0
    n := 0

    for i:=0; i<len(nums); i++ {
        if nums[i]%2==0 && nums[i]%3==0{
            n++
            avg+=nums[i]
        }
    }
    if n!=0 {
        avg = avg/n
    }    
    return avg
}