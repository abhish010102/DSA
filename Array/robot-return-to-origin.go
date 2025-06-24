func judgeCircle(moves string) bool {
    x:=0
    y:=0

    for _,ch := range moves{
        if ch=='R'{
            x+=1
        }else if ch=='L'{
            x-=1
        }else if ch=='U'{
            y+=1
        }else if ch=='D'{
            y-=1
        }
    }

    if x==0 && y==0{
        return true
    }
    return false
}