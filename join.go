package piscine

func Join(strs []string, sep string) string {
	c := ""
	for i, k := range strs {
		c += k
		if i != len(strs)-1 {
			c += sep
		}
	}
	return c
}
