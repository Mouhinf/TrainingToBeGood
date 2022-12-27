/*Write a recursive function that returns the value of nb to the power of power.

Negative powers will return 0. Overflows do not have to be dealt with.

for is forbidden for this exercise.
*/

package main

import "fmt"

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 1 {
		return nb
	}
	if power == 0 {
		return 1
	}
	return (nb * RecursivePower(nb, power-1))

}

func main() {
	var nb int
	var power int
	fmt.Print("Give the number that you want know the power : ")
	fmt.Scan(&nb)
	fmt.Print("Give the number the power : ")
	fmt.Scan(&power)
	fmt.Println(RecursivePower(nb, power))
}
