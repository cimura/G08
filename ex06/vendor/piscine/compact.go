package piscine

func Size(str []string) int {
	count := 0
	for _, s := range str {
		if s != "" {
			count++
		}
	}
	return count
}

func Compact(ptr *[]string) int {
	count := 0
	compact := make([]string, Size(*ptr))
	for _, s := range *ptr {
		if s != "" {
			compact[count] += s
			count++
		}
	}
	*ptr = compact
	return count
}