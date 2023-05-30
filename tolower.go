package piscine

func ToLower(s string) string {
	P := []rune(s)
	result := ""
	for i := 0; i < len(P); i++ {
		if (P[i] >= 'A') && P[i] <= 'Z' {
			P[i] = P[i] + 32
		}
		result += string(P[i])
	}
	return result
}
