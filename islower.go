package piscine

func IsLower(s string) bool {
	N := []rune(s)
	for i := 0; i < len(s); i++ {
		if (N[i] >= 0 && N[i] < 97) || (N[i] > 122 && N[i] <= 127) {
			return false
		}
	}
	return true
}
