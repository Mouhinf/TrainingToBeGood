package piscine

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func IsSorted(f func(a, b int) int, a []int) bool {
	s := true
	c := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			c = false
		}
		if f(a[i], a[i+1]) < 0 {
			s = false
		}
	}
	return c || s
}
