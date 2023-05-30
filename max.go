package piscine

func Max(a []int) int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] < a[j] {
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp

			}
		}
	}
	return a[len(a)-1]
}
