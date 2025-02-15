package main

import (
	"fmt"
	"math"
)

func main() {
	// Declare and initialize variables:
	//const is an operator that declares a constant
	const IMTPower = 2

	// we can declear by using := operator
	// when we use := operator we do not need to specify the type of the variable
	// go will automatically detect the type of the variable
	userHeight := 1.7
	var userWeight float64 = 100

	IMT := userWeight / math.Pow(userHeight, IMTPower)

	fmt.Printf("IMT is : %f", IMT)
}

//for build the code run the command below:
//go build -o 02_binary 02-main.go
//becouse we few main fail in the same directory we need to specify the name of the file we want to build
//we cant mod init because we use at the same project directory few main files
