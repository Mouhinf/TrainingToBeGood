package piscine

func ListClear(l *List) {
	for l.Head != nil {
		n := l.Head
		l.Head = n.Next
	}
}
