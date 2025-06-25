func numJewelsInStones(jewels string, stones string) int {
	ele := map[int]int{}
	ret := 0

	for _, val := range stones {
		ele[int(val-'a')]++
	}

	for _, val := range jewels {
		if _, check := ele[int(val-'a')]; check {
			ret += ele[int(val-'a')]
		}
	}

	return ret
}

