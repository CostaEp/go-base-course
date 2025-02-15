package main

import (
	"fmt"
	"math"
)

func main() {

	var userHeight float64
	var userWeight float64

	fmt.Println(`		---This program calculates---
	  ---the IMT (Index of Mass of the Body)---`)
	fmt.Print("Enter your height in meters (for example 1.6): ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight in kilograms (for example 70): ")
	fmt.Scan(&userWeight)
	// call the function calcIMT and pass the userHeight and userWeight values
	IMT := calcIMT(userHeight, userWeight)
	resultIMT(IMT)
}

// resultIMT function takes a float64 value and prints the result
func resultIMT(imt float64) {
	result := fmt.Sprintf("your IMT is : %.2f", imt)
	fmt.Print(result)
}

// calcIMT function takes two float64 values and returns the result
func calcIMT(userHeight float64, userWeight float64) float64 {
	const IMTPower = 2
	IMT := userWeight / math.Pow(userHeight, IMTPower)
	// call the function resultIMT and pass the IMT value
	return IMT
}
