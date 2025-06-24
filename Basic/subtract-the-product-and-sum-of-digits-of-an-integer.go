func subtractProductAndSum(n int) int {
    p := 1
    s := 0

    for n>0 {
        l := n%10
        p *= l
        s += l
        n /= 10
    }

    return (p-s)
}