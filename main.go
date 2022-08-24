package main

import "fmt"

func main() {
	num := 38
	fmt.Println(addDigits(num))
}

func addDigits(num int) int {
	var r int
	sum := 0
	for num > 0 {

		r = num % 10
		sum = r + sum
		num = num / 10

	}
	if sum > 9 {
		sum = addDigits(sum)
	}
	return sum
}
