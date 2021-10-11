package main

import "fmt"

func main() {
	fmt.Println("The program prints Ping, Pong or a number.")
	fmt.Println("Enter a number:")
	var number int
	fmt.Scanln(&number)
	if number%2 != 0 && number%5 != 0 {
	} else if number%2 != 0 {
		fmt.Println("Ping")
	} else if number%5 != 0 {
		fmt.Println("Pong")
	} else {
		fmt.Println(number)
	}
}
