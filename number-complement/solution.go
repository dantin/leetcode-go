package p476

func FindComplement(num int) int {
	mask := ^0
	for (mask & num) > 0 {
		mask <<= 1
	}
	return ^mask & ^num
}
