package piscine

func Join(strs []string, sep string) string {
	Str := ""
	for i, s := range strs {
		if i > 0 {
			Str += sep
		}
		Str += s
	}
	return Str
}
