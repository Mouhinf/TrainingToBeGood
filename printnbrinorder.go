package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n > 0 {
		var t []int
		V := 0
		var R int
		c := 0
		for n != 0 {
			V = n % 10
			n /= 10
			t = append(t, V)
		}
		for d := range t {
			c = d + 1
		}
		for i := 0; i < c-1; i++ {
			for j := 0; j < c-i-1; j++ {
				if t[j] > t[j+1] {
					R = t[j]
					t[j] = t[j+1]
					t[j+1] = R
				}
			}
		}
		for i := 0; i < c; i++ {
			c := 0
			var V int
			for n != 0 {
				R = n % 10
				n /= 10
				t = append(t, V)
			}
			for d := range t {
				c = d + 1
			}
			for i := 0; i < c-1; i++ {
				for j := 0; j < c-i-1; j++ {
					if t[j] > t[j+1] {
						R = t[j]
						t[j] = t[j+1]
						t[j+1] = V
					}
				}
			}
		}
		for i := 0; i < c; i++ {
			z01.PrintRune(rune(t[i] + 48))
		}
	}
}
