package piscine

func Index(s string, toFind string) int {
	a := []rune(s)
	b := []rune(toFind)
	n := len(a)
	m := len(b)
	for i := 0; i <= n-m; i++ {
		if toFind == s[i:i+m] {
			return i
		}
	}
	return -1
}
