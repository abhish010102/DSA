// Using Ex-or Gate
func singleNonDuplicate1(nums []int) int {
	num := 0

	for i := 0; i < len(nums); i++ {
		num = num ^ nums[i]
	}

	return num
}

//Using Binary Search 
func singleNonDuplicate(nums []int) int {
	i, j := 0, len(nums)-1

	for j > i {
		m := (i + j) / 2
        if m%2==1{
            m-=1
        }

        if nums[m] == nums[m+1]{
            i=m+2
        }else if m>0 && nums[m] == nums[m-1]{
            j=m-2
        }else{
            return nums[m]
        }
    
    }

	return nums[i]
}