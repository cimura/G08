package piscine

func SortWordArr(a []string) {
	count := 0
	for range a {
		count++
	}
	for i := 0; i < count - 1; i++ {
		for j := 0; j < count - i - 1; j++ {
			if a[j] > a[j + 1] {
				a[j], a[j + 1] = a[j + 1], a[j]
			}
		}
	}
}