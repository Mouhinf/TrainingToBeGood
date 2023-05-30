package piscine

func IterativeFactorial(nb int) int {
	f := 1
	if nb < 0 || nb <= -32767 || nb >= 32767 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		f = f * i
	}
	return f
}
