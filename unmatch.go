package piscine

func Unmatch(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				s++
			}
		}
		if s%2 != 0 {
			return a[i]
		} else {
			s = 0
		}
	}
	return -1
}
