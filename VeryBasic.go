// In this exercise this exercices, I learn the loop et the instructions
package main

import "fmt"

func main() {
	var nb int
	f := 1
	fmt.Print("Give the number that you want say the multiplication table : ")
	fmt.Scan(&nb)
	if nb <= 0 {
		fmt.Println("Don't use the negative or nil number ")
		fmt.Print("Give a correct number : ")
		fmt.Scan(&nb)
	}
	for i := 1; i <= 10; i++ {
		f = nb * i
		fmt.Println(nb, " multiply for ", i, "give", f)
	}
}
