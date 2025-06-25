func findMaxK(nums []int) int {
    out := -1
    slices.Sort(nums)
    i := 0
    j := len(nums) - 1

    for i <= j {
        if nums[i] + nums[j] == 0 {
            return int(math.Abs(float64(nums[i])))
        }
        if -nums[i] > nums[j] {
            i++
        } else {
            j--
        }
    }

    return out
}