package piscine

func TrimAtoi(s string) int {
	r := 0
	n := 0
	c := 0
	d := 1
	for i := 0; i <= len(s)-1; i++ {
		c = int(s[i]) - 48
		if c == -3 {
			if n == 0 {
				d = -1
				r++
			}
		} else if c >= 0 && c <= 9 {
			n = (10 * n) + c
			r++
		}
	}
	n = n * d
	return n
}
