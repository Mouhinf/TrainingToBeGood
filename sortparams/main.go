package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args
	for i := 1; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				tempo := a[i]
				a[i] = a[j]
				a[j] = tempo
			}
		}
	}
	for i := 1; i < len(a); i++ {
		arg := os.Args[i]
		for _, r := range arg {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
