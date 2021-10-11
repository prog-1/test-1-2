package main

import "fmt"

// Replace the comments below with your explanations. Try to provide a
// "high-level" idea, instead of writing literally what each line does.
//
// TODO: What does the program do?
func main() {
	var num uint
	fmt.Scan(&num)

	// TODO: What do the variables a, b and c store?
	a, b, c := num%10, num/10%10, num/100%10
	// a = remainder of the num divided by 10
	// b = remainder of the number divided by 10, which was divided by 10
	// c = remainder of the number divided by 10, which was divided by 100

	// TODO: What do the three `if ... { ... }` statements below do?
	if a < b {
		a, b = b, a
	}
	// if a < b, then a = b and b = a
	if a < c {
		a, c = c, a
	}
	// if a < c, then a = c and c = a
	if b < c {
		b, c = c, b
	}
	// if b < c, then b = c and c = b

	// TODO: What do the three lines below do?
	num = a
	num = num*10 + b
	num = num*10 + c
	// num becomes a, num multiplies by 10 + b, num multiplies by 10 + c

	fmt.Println(num)
}
