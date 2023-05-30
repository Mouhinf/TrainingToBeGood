package piscine

func NRune(s string, n int) rune {
	f := []rune(s)
	for index, letter := range f {
		if index == n-1 {
			return letter
		}
	}
	return 0
}
