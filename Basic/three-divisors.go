func isThree(n int) bool {
    c := 1;

    for i:=2 ; i<=n ; i++ {
        if n%i==0 {
            c++;
            if c>3 {
                return false
            }
        }
    }
    if c==3 {
        return true
    }

    return false
}