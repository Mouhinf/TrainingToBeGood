package main

import "fmt"

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	if index >= 3 {
		return (Fibonacci(index-1) + Fibonacci(index-2))
	} else {
		return 1
	}

}

func main() {
	var nb int
	fmt.Print("Give your number : ")
	fmt.Scan(&nb)
	fmt.Println(Fibonacci(nb))
}

// For look the order of exercise watch the ReadMe
