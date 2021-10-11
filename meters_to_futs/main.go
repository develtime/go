package main

import "fmt"

func main() {
	fmt.Print("Enter a number in meters: ")

	var meters float64

	fmt.Scanf("%f", &meters)

	fmt.Printf("Your number in futs is: %f\n", MeterToFuts(meters))
}
