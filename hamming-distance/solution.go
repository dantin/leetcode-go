package p461

func HammingDistance(x int, y int) int {
	d := 0
	xor := x ^ y

	for xor != 0 {
		d++
		xor &= (xor - 1)
	}

	return d
}
