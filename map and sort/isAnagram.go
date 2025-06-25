func isAnagram(s string, t string) bool {
    c := make([]int, 26)

    for _, val := range s {
        c[int(val - 'a')]++
    }
    for _, val := range t {
        c[int(val - 'a')]--
    }

    for _, val := range c {
        if val != 0 {
            return false
        }
    }

    return true
}