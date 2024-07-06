package piscine

func Rot14(s string) string {
	sRune := []rune(s)
	for i, r := range sRune {
		switch {
		case 'a' <= r && r <= 'z':
			sRune[i] = 'a' + (r - 'a' + 14) % 26
		case 'A' <= r && r <= 'Z':
			sRune[i] = 'A' + (r - 'A' + 14) % 26
		}
	}
	return string(sRune)
}