package piscine

func IsAlpha(s string) bool {
	T := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if (T[i] >= 0) && (T[i] <= 47) || (T[i] >= 58) && (T[i] <= 64) && (T[i] >= 91 && T[i] <= 96) || (T[i] >= 123) && (T[i] <= 127) {
			return false
		}
	}
	return true
}
