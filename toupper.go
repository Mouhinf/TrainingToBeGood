package piscine

func ToUpper(s string) string {
	T := []rune(s)
	result := ""
	for i := 0; i < len(T); i++ {
		if T[i] >= 'a' && T[i] <= 'z' {
			T[i] = T[i] - 32
		}
		result += string(T[i])
	}
	return result
}
