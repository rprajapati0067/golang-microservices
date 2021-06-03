package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 3)

	c <- "1"
	fmt.Println("1")
	c <- "2"
	fmt.Println("2")
	c <- "3"
	fmt.Println("3")

}
func helloWorld() {
	fmt.Print("Hello World!")
}
