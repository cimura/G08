package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	count := 0
	for range a {
		count++
	}
	for i := 0; i < count - 1; i++ {
		if f(a[i], a[i + 1]) > 0 {
			return false
		}
	}
	return true
}