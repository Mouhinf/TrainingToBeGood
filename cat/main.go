package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func res(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}

func main() {
	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			f, e := ioutil.ReadFile(os.Args[i])
			if e != nil {
				err := "ERROR: " + e.Error() + "\n"
				res(err)
				os.Exit(1)
			}
			res(string(f))
		}
	} else {
		f, _ := ioutil.ReadAll(os.Stdin)
		res(string(f))
	}
}
