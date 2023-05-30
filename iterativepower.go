package piscine

func IterativePower(nb int, power int) int {
	var i int
	var P int
	P = nb
	if power < 0 {
		return 0
	}
	if power == 1 {
		return nb
	}
	if power == 0 {
		return 1
	}
	for i = 1; i < power; i++ {
		P = P * nb
	}
	return P
}
