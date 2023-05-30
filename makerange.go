package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	tab := make([]int, max-min)
	tab[0] = min
	for i := 1; i < max-min; i++ {
		tab[i] = min + i
	}
	return tab
}
