package main

func MeterToFuts(meters float64) float64 {
	const coef = 0.3048 //Fut coefficient
	return meters / coef
}
