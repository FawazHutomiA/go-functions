package main

import (
	"fmt"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello world")
}

func main() {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}
