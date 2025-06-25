func checkIfPangram(sentence string) bool {
    ele := map[int]int{}
    for _, val := range sentence {
        ele[int(val - 'a')]++
        if len(ele)==26{
            return true
        }
    }
    return false
}