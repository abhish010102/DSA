func twoSum(nums []int, target int) []int {
    // total := target

    // for i, val := range nums{
    //     total = target - val
    //     for j:=i+1; j<len(nums); j++ {
    //         if total-nums[j] == 0 {
    //             return []int{i,j}
    //         }
    //     } 
    // }
    // return []int{0,0}

    m := make(map[int]int)

    for i, val := range nums{
        j, bool1 := m[val]
        if bool1 {
            return []int{i,j}
        }
        m[(target-val)] = i
    }
    
    return []int{0,0}
}