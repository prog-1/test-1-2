package main

import "fmt"

func main() {
	fmt.Println("The program finds the capacity of three inductors connected in parallel")
	fmt.Println("Enter inductance of three inductors: 1 2 3")
	var L1, L2, L3 float64
	fmt.Scanln(&L1, &L2, &L3)
	Cn := 1/L1 + 1/L2 + 1/L3
	fmt.Println("The total inductance of the three inductors connected in parallel.")
}