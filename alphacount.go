package piscine

func AlphaCount(s string) int {
	nb := 0
	p := []rune(s)
	for i := 0; i < len(p); i++ {
		if (p[i] >= 65 && p[i] <= 90) || (p[i] >= 97 && p[i] <= 122) {
			nb++
		}
	}
	return nb
}
