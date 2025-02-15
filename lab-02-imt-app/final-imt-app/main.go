package main

import (
	"fmt"
	"math"
)

func main() {

	const IMTPower = 2
	var userHeight float64
	var userWeight float64

	//using Println print and go to the next line
	//we can use bacticks to print multiple lines
	fmt.Println(`		---This program calculates---
	  ---the IMT (Index of Mass of the Body)---`)
	fmt.Print("Enter your height in meters (for example 1.6): ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight in kilograms (for example 70): ")
	fmt.Scan(&userWeight)

	IMT := userWeight / math.Pow(userHeight, IMTPower)
	// printf and %.2f to print only 2 decimal places
	//using format you cant go to go documentaion https://pkg.go.dev/fmt
	//if we whant only to save the value we can use Sprintf
	result := fmt.Sprintf("your IMT is : %.2f", IMT)
	fmt.Print(result)
}

//final lets build uor code
//o mod init final-imt-app
//o build
// run the next command to run the app
// ./final-imt-app
