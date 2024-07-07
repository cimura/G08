package piscine

func Max(a []int) int {
	max := -9223372036854775808
	for _, value := range a {
		if max < value {
			max = value
		}
	}
	return max
}
