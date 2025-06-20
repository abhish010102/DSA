func maxDepth(s string) int {
    max := 0
    b := 0

    for _, chr:= range s{
        if chr == '('{
            b++
            if b>max {
                max=b
            }
        }else if chr == ')'{
            b--
        }
    }

    return max
}