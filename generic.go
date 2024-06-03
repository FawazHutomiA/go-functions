package main

import "fmt"

func MultipleTypeParameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func main() {
	MultipleTypeParameter[string, int]("Fawaz", 20)
	MultipleTypeParameter[int, string](30, "Fawaz")
}
