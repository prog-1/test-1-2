package main

import "fmt"

// Replace the comments below with your explanations. Try to provide a
// "high-level" idea, instead of writing literally what each line does.
//
//This function takes 3-digit number and sorted in descending order
func main() {
	var num uint
	fmt.Scan(&num)
	// Tthis variables splits number in 3 parts
	a, b, c := num%10, num/10%10, num/100%10

	// Sorts varibles a b c, so a is the largest and c is the lowest
	if a < b {
		a, b = b, a
	}
	if a < c {
		a, c = c, a
	}
	if b < c {
		b, c = c, b
	}

	// Combines a b c into a number
	num = a
	num = num*10 + b
	num = num*10 + c

	fmt.Println(num)
}
