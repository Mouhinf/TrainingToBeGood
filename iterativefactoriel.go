package main

import "fmt"

func IterativeFactorial(nb int) int {
	f := 1
	for i := 1; i <= nb; i++ {
		f = f * i
	}
	return f
}

func main() {
	var nb int
	fmt.Print("Give the number that you want know its factorial : ")
	fmt.Scan(&nb)
	fmt.Println(IterativeFactorial(nb))
}
