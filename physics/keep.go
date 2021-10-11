package main

import "fmt"

func main() {
	fmt.Println("The program finds the capacity of three inductors connected in parallel.")
	fmt.Println("Enter numbers")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	var l = 1/a + 1/b + 1/c
	fmt.Println("result", l)
}
