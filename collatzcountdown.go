package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	s := 0
	for {
		if start%2 == 0 {
			start = start / 2
			s = s + 1
		} else {
			start = (start * 3) + 1
			s = s + 1
		}
		if start == 1 {
			return s
		}
	}
	return s
}
