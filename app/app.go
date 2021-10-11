package main

import "fmt"

// Replace the comments below with your explanations. Try to provide a
// "high-level" idea, instead of writing literally what each line does.
//
// TODO: This function takes 3-digit number and sorted in descending order.
func main() {
	var num uint
	fmt.Scan(&num)

	// TODO: These variables splits number in 3 parts.
	a, b, c := num%10, num/10%10, num/100%10

	// TODO: Sorts variables a, b, c.
	if a < b {
		a, b = b, a
	}
	if a < c {
		a, c = c, a
	}
	if b < c {
		b, c = c, b
	}

	// TODO: Combines a, b, c into number
	num = a
	num = num*10 + b
	num = num*10 + c

	fmt.Println(num)
}
