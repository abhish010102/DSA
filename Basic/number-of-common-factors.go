func commonFactors(a int, b int) int {
    c := 1

    for i:=2 ; i<= min(a,b); i++{
        if a%i==0 && b%i==0 {
            c++
        }
    } 

    return c
}