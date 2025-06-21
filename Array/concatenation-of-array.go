func getConcatenation(nums []int) []int {
    // var arr [5]int
    l := len(nums)
    arr:= make([]int, l*2)

    for i:=0; i<l; i++{
        arr[i] = nums[i]
        arr[i+l] = nums[i]
    }

    return arr;
}

func getConcatenation1(nums []int) []int {
	return append(nums, nums...)
}