package main

import "fmt"

// Replace the comments below with your explanations. Try to provide a
// "high-level" idea, instead of writing literally what each line does.
//
// TODO: What does the program do?
func main() {
	var num uint
	fmt.Scan(&num) // person enters number
	// this program is similar to program in homework wich places numbers in order
	// TODO: What do the variables a, b and c store?
	a, b, c := num%10, num/10%10, num/100%10
	// % as the compiler knows that %num is the placeholder for int.
	// TODO: What do the three `if ... { ... }` statements below do?
	if a < b {
		a, b = b, a //all if chek if number is bigger or smaller and changes the order of numbers
	}
	if a < c {
		a, c = c, a
	}
	if b < c {
		b, c = c, b
	}

	// TODO: What do the three lines below do?
	num = a // it combines all number (a*num10+b )*10+c
	num = num*10 + b
	num = num*10 + c

	fmt.Println(num) // prints result
}
