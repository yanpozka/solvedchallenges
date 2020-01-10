package bitwise

func rangeBitwiseAnd(m int, n int) int {
	res := uint(m)

	for c, lim := uint(m+1), uint(n); c <= lim; c++ {
		// fmt.Printf("%b\n", res)
		res &= c
		if res == 0 {
			break
		}
	}

	return int(res)
}
