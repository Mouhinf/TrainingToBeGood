package piscine

func FindNextPrime(nb int) int {
	i := nb
	for {
		if IsPrime(i) {
			return i
		}
		i++
	}
}
