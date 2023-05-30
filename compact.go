package piscine

func Compact(ptr *[]string) int {
	T := 0
	f := make([]string, T)
	for _, i := range *ptr {
		if i != "" {
			f = append(f, i)
			T++
		}
	}
	*ptr = f
	return T
}
