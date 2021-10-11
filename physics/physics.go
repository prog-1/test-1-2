package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program finds the capacity of three inductors connected in parallel.")
	fmt.Print("Enter inductance of three inductors:")
	var L1, L2, L3 float64
	fmt.Scan(&L1, &L2, &L3)
	var L float64
	L = 1 / ((1 / L1) + (1 / L2) + (1 / L3))
	fmt.Println("The total inductance of the three inductors connected in parallel is:", L)

}
