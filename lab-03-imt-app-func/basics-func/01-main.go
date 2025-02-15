package main

import (
	"fmt"
	"math"
)

func main() {

	const IMTPower = 2
	var userHeight float64
	var userWeight float64

	fmt.Println(`		---This program calculates---
	  ---the IMT (Index of Mass of the Body)---`)
	fmt.Print("Enter your height in meters (for example 1.6): ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight in kilograms (for example 70): ")
	fmt.Scan(&userWeight)
	IMT := userWeight / math.Pow(userHeight, IMTPower)
	// call the function resultIMT and pass the IMT value
	resultIMT(IMT)

}

// resultIMT function takes a float64 value and prints the result
func resultIMT(imt float64) {
	result := fmt.Sprintf("your IMT is : %.2f", imt)
	fmt.Print(result)
}

//this is a primitive example of a function that takes a float64 value and prints the result
