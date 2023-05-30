package piscine

func Map(f func(int) bool, a []int) []bool {
	ar := make([]bool, len(a))
	for k, m := range a {
		ar[k] = f(m)
	}
	return ar
}
