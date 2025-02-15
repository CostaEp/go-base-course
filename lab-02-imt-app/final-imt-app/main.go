package main

import (
	"fmt"
	"math"
)

func main() {

	const IMTPower = 2
	var userHeight float64
	var userWeight float64
	fmt.Printf("This program calculates the IMT (Index of Mass of the Body)\n")
	fmt.Printf("Enter your height in meters (for example 1.6): ")
	fmt.Scan(&userHeight)
	fmt.Printf("Enter your weight in kilograms (for example 70): ")
	fmt.Scan(&userWeight)

	IMT := userWeight / math.Pow(userHeight, IMTPower)

	fmt.Printf("your IMT is : %f", IMT)
}

//final lets build uor code
//go mod init final-imt-app
//go build -o final_binary main.go
// run the binary file by typing ./final_binary
