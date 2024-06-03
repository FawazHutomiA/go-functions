package main

import (
	"errors"
	"fmt"
)

const showExpectedResult = false
const showHints = false

func CalculateMean(numbers []int) (*float64, error) {
	if numbers == nil || len(numbers) == 0 {
		return nil, errors.New("Empty Numbers")
	}

	var sum int
	for _, num := range numbers {
		sum += num
	}

	mean := float64(sum) / float64(len(numbers))
	return &mean, nil
}

func main() {
	numbers := []int{4, 6, 9, 45, 8, 17, 3}
	learnerResult, learnerError := CalculateMean(numbers)
	fmt.Println("Your code returned: ", *learnerResult, learnerError)
}
