package main

import "testing"

func TestMin(t *testing.T) {
	numbers := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}

	min, err := Min(numbers)

	if err != nil || min != 9 {
		t.Error("Expected 5, got ", min)
	}
}
