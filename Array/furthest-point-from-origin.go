func furthestDistanceFromOrigin(moves string) int {
    d:=0
    l:=0
    r:=0

    for _,ch := range moves{
        if ch=='R'{
            r++
        }else if ch=='L'{
            l++
        }else if ch=='_'{
            d++
        }
    }

    return max(l,r)-min(l,r)+d
}