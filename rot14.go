package piscine

func Rot14(s string) string {
	help := []rune(s)
	for i := 0; i < len(help); i++ {
		if help[i] != 32 || help[i] != 33 {
			if (help[i] >= 97 && help[i] < 109) || (help[i] >= 65 && help[i] < 77) {
				help[i] = help[i] + 14
			} else if (help[i] >= 109 && help[i] <= 122) || (help[i] > 76 && help[i] <= 90) {
				help[i] = help[i] - 12
			}
		}
	}
	return string(help)
}
