package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	choumi := os.Args
	rev := ""
	for _, a := range choumi[1:] {
		rev = a + "\n" + rev
	}
	for _, b := range rev {
		z01.PrintRune(b)
	}
}
