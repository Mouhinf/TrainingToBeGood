package piscine

func ListLast(l *List) interface{} {
	n := l.Head
	if n == nil {
		return nil
	}
	for n != nil {
		if n.Next == nil {
			return n.Data
		}
		n = n.Next
	}
	return n.Data
}
