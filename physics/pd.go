package main

import "fmt"

func main() {
fmt.Println("Thr program finds the capacity  of three inductors connected in parallel.")
fmt.Println("Enter inductance of three inductors:")
var a, b, c float64
fmt.Scan(&a, &b, &c)
result := 1 / (1/a + 1/b + 1/c)
fmt.Println("The total inductance of the three inductors connected in parallel is",result)
}