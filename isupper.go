package piscine

func IsUpper(s string) bool {
	p := []rune(s)
	for i := 0; i < len(p); i++ {
		if (p[i] >= 0 && p[i] <= 64) || (p[i] >= 91 && p[i] <= 127) {
			return false
		}
	}
	return true
}
