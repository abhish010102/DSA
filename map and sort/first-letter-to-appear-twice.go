func repeatedCharacter(s string) byte {
    c := make([]int, 26)

    for _, val := range s {
        c[int(val - 'a')]++
        if c[int(val - 'a')]==2{
            return byte(val)
        }
    }
    return 0
}