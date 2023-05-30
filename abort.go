package piscine

func Abort(a, b, c, d, e int) int {
	t := make([]int, 5)
	t[0] = a
	t[1] = b
	t[2] = c
	t[3] = d
	t[4] = e
	for i := 0; i < len(t); i++ {
		for j := i + 1; j < len(t); j++ {
			if t[i] > t[j] {
				tmp := t[i]
				t[i] = t[j]
				t[j] = tmp
			}
		}
	}
	return t[2]
}
