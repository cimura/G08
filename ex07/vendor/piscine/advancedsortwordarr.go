package	piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	count := 0
	for range a {
		count++
	}
	for i := 0; i < count - 1; i++ {
		for j := 0; j < count - i - 1; j++ {
			if f(a[j], a[j + 1]) > 0 {
				a[j], a[j + 1] = a[j + 1], a[j]
			}
		}
	}
}