package main

import "fmt"

// Replace the comments below with your explanations. Try to provide a
// "high-level" idea, instead of writing literally what each line does.
//
// TODO: What does the program do?
func main() {
	var num uint // Here the user must enter an integer and a posititve number.
	fmt.Scan(&num)

	// TODO: What do the variables a, b and c store?
	a, b, c := num%10, num/10%10, num/100%10 // a contains a number with a reminder of 10, b contains a number divivded by 10 with a reminder of 10 and c contains a number divivded by 100 with a reminder of 10.

	// TODO: What do the three `if ... { ... }` statements below do?
	if a < b { // These three if statements put a, b and c in descending order.
		a, b = b, a
	}
	if a < c {
		a, c = c, a
	}
	if b < c {
		b, c = c, b
	}

	// TODO: What do the three lines below do?
	num = a          // Contains variable a.
	num = num*10 + b // Multiplies num by 10 and adds b.
	num = num*10 + c // Multiplies num by 10 and adds c.

	fmt.Println(num)
}
