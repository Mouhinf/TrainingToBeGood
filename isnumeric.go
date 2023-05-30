package piscine

func IsNumeric(s string) bool {
	N := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if N[i] > 0 && N[i] < 48 || N[i] > 57 && N[i] <= 127 {
			return false
		}
	}
	return true
}
