package piscine

func ListSize(l *List) int {
	s := 0
	c := l.Head
	for c != nil {
		s++
		c = c.Next
	}
	return s
}
