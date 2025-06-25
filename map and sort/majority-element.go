func majorityElement(nums []int) int {
    // t := 0 
    // c := 1

    // for i:=1; i<len(nums);i++{
    //     if nums[i]==nums[t]{
    //         c++
    //     }else{
    //         c--
    //     }
    //     if(c==0){
    //         t=i;
    //         c=1;
    //     }
    // }   
    // return nums[t]
    ret := 0
    mc := 0
    ec := map[int]int{}

    for i:=0; i<len(nums);i++{
        ec[nums[i]]++;
        if(ec[nums[i]]>mc){
            mc=ec[nums[i]]
            ret=nums[i]
        }
    }

    return ret
}