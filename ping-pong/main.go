package main

import "fmt"

func main() {
	fmt.Println("The program prints Ping, Pong or a number.")
	fmt.Print("Enter a number: ")
	var number int
	fmt.Scan(&number)
	if number%5 != 0 && number%2 != 0 {
		fmt.Println("PingPong")
		return
	}
	if number%2 != 0 {
		fmt.Println("Ping")
		return
	}
	if number%5 != 0 {
		fmt.Println("Pong")
		return
	}
	fmt.Println(number)
}
