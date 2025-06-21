func findGCD(nums []int) int {
    min :=nums[0]
    max :=nums[0]

    ret := 1

    for _,no := range nums{
        if min > no{
            min=no
        }else if max<no{
            max=no
        }
    }

    for i:=2;i<=min;i++{
        if min%i==0 && max%i==0{
            ret=i
        }
    }   

    return ret
}