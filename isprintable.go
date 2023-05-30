package piscine

func IsPrintable(s string) bool {
	P := []rune(s)
	for i := 0; i < len(P); i++ {
		if P[i] >= 0 && P[i] < 32 {
			return false
		}
	}
	return true
}
