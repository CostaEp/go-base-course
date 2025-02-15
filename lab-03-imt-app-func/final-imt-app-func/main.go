package main

import (
	"fmt"
	"math"
)

// define the constant IMTPower
// IMTPower is const out of the scope and we can use it in the whole program
const IMTPower = 2

func main() {

	fmt.Println(`		---This program calculates---
	  ---the IMT (Index of Mass of the Body)---`)
	// call the function getUserInput and get the userHeight, userWeight and userName values
	//the order of the values returned by the function is the same
	// as the order of the values in the function signature
	userHeight, userWeight, userName := getUserInput()
	// call the function calcIMT and pass the userHeight and userWeight values
	IMT := calcIMT(userHeight, userWeight)

	// call the function resultIMT and pass the IMT and userName values
	resultIMT(IMT, userName)
}

// now resultIMT function takes a float64 value and a string value and prints the result
func resultIMT(imt float64, userName string) {
	//format the result string to include the userName and the imt value
	result := fmt.Sprintf("%s your calculate IMT is : %.2f", userName, imt)
	fmt.Print(result)
}

// calcIMT function takes two float64 values and returns the result
func calcIMT(userHeight float64, userWeight float64) float64 {
	IMT := userWeight / math.Pow(userHeight, IMTPower)
	// call the function resultIMT and pass the IMT value
	return IMT
}

// getUserInput function returns three values: float64, float64 and string
func getUserInput() (float64, float64, string) {
	var userName string
	var userHeight float64
	var userWeight float64
	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)
	fmt.Print("Enter your height in meters (for example 1.6): ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight in kilograms (for example 70): ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight, userName
}
