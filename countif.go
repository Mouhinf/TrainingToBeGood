package piscine

func CountIf(f func(string) bool, tab []string) int {
	S := 0
	for i := 0; i < len(tab); i++ {
		if f(tab[i]) {
			S = S + 1
		}
	}
	return S
}
