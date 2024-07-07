package piscine

func Len(a []int) int {
	count := 0
	for range a {
		count++
	}
	return count
}

func Unmatch(a []int) int {
	for i := 0; i < Len(a); i++ {
		found := false
		for j := 0; j < Len(a); j++ {
			if i != j && a[i] == a[j] {
				found = true
				break
			}
		}
		if !found {
			return a[i]
		}
	}
	return -1
}
