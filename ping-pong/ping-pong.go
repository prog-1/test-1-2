package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program prints Ping, Pong or a number.")
	fmt.Print("Enter a number:")
	var a int
	fmt.Scan(&a)
	if a%5 != 0 && a%2 != 0 {
		fmt.Println("PingPong")
		return

	}
	if a%2 != 0 {
		fmt.Println("Ping")
		return

	}
	if a%5 != 0 {
		fmt.Println("Pong")
		return

	}
	fmt.Println(a)

}
