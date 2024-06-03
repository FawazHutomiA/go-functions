package main

import "fmt"

func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {
	panggil := getGoodBye

	fmt.Println(panggil("Fawaz"))
}
