package main

import "fmt"

func Min(values []int) (int, error) {
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
