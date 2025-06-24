func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}

    rev := 0
    num := x
    
    for num > 0 {
        rev *=10
        l := num%10
        rev += l
        num /= 10
    }

    if rev == x{
        return true
    }

    return false
}