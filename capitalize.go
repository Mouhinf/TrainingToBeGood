package piscine

func Capitalize(s string) string {
	d := []rune(s)
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if d[i] >= 97 && d[i] <= 122 {
				d[i] = d[i] - 32
			}
		} else {
			if d[i-1] < 48 || (d[i-1] > 57 && d[i-1] < 65) || (d[i-1] > 90 && d[i-1] < 97) || (d[i-1] > 122) {
				if d[i] >= 97 && d[i] <= 122 {
					d[i] = d[i] - 32
				}
			} else {
				if d[i] >= 65 && d[i] <= 90 {
					d[i] = d[i] + 32
				}
			}
		}
	}
	return string(d)
}
