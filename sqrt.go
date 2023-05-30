package piscine

func Sqrt(nb int) int {
	var l int
	var m int
	l = nb
	m = 1

	if nb < 0 {
		return 0
	}
	for i := 1; i <= l-m; i++ {
		l = (l + m) / 2
		m = nb / l
	}
	if nb == l*l {
		return l
	} else {
		return 0
	}
}
