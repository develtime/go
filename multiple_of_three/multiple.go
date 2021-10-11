package main

import "fmt"

func Multiple(start int, end int) {
	for i := start; i <= end; i++ {
		if i%3 == 0 {
			fmt.Printf("Number %d is multiple of three\n", i)
		}
	}
}
