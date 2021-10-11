package main

import "fmt"

func main() {
	for value := 1; value <= 100; value++ {
		if value%3 == 0 {
			fmt.Printf("Number %d is multiple of three\n", value)
		}
	}
}
