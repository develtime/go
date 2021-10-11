package main

import "fmt"

func min(values []int) (int, error) {
	if values == nil || len(values) == 0 {
		return 0, fmt.Errorf("Array can't be nil or empty")
	}

	min := values[0]

	for _, value := range values {
		if value < min {
			min = value
		}
	}

	return min, nil
}

func main() {
	numbers := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}

	min, err := min(numbers)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Smalles item in %v is %d\n", numbers, min)
	}
}
