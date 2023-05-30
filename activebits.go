package piscine

func ActiveBits(n int) int {
	s := 1
	for n/2 != 0 {
		if n%2 == 1 {
			s++
		}
		n /= 2
	}
	return s
}
