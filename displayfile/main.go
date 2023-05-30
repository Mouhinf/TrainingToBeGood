package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	nA := len(os.Args)
	if nA < 2 {
		fmt.Println("File name missing")
	} else if nA > 2 {
		fmt.Println("Too many arguments")
	} else if nA == 2 {
		fn := os.Args[1]
		f, e := os.Open(fn)
		if e != nil {
			fmt.Println("Error : ", e)
		}
		fc, e := ioutil.ReadAll(f)
		if e != nil {
			fmt.Println("Error cant read file:", e)
		}
		fmt.Print(string(fc))
		f.Close()
	}
}
