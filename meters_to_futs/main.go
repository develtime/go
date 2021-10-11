package main

import (
	"fmt"
	"futs"
)

// import "meters_to_futs.go"

func main() {
	fmt.Print("Enter a number in meters: ")

	var meters float64

	fmt.Scanf("%f", &meters)

	fmt.Printf("Your number in futs is: %f\n", futs.MeterToFuts(meters))
}
