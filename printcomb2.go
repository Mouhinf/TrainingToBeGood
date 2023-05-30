package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			l := j + 1
			for k := i; k <= '9'; k++ {
				for ; l <= '9'; l++ {
					if i == '0' && j == '0' && k == '0' && l == '1' {

						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(l)
						z01.PrintRune(',')

					} else if i == '9' && j == '8' && k == '9' && l == '9' {

						z01.PrintRune(' ')
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(l)

					} else {

						z01.PrintRune(' ')
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(l)
						z01.PrintRune(',')
					}
				}
				l = '0'
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintComb2()
}
