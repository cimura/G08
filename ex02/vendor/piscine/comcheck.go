package piscine

func Comcheck(s []string) bool {
	for _, Str := range s {
		if Str == "42" || Str == "piscine" || Str == "piscine 42" {
			return true
		}
	}
	return false
}