package main

import (
	"fmt"
	"math"
)

func main() {
	// Declare and initialize variables:
	//var is an operator that declares a variable
	// userHeight is the name of the variable we use camelCase for naming variables
	var userHeight float64 = 1.7
	var userWeight float64 = 100

	//we can not give a variable type and convert later
	// var userWeight = 100

	// math.Pow is a function that takes two arguments and returns the first argument raised to the power of the second argument
	// example : math.Pow(2, 3) = 8
	var IMT = userWeight / math.Pow(userHeight, 2)

	// we can also use float64() to convert the type of the variable
	// var IMT = float64(userWeight) / userHeight

	fmt.Printf("IMT is : %f", IMT)
}

// go var type:
//-------------------------------
// bool : true, false
// string : "Hello, World!"
// int : 42
// int8 : -128 to 127
// int16 : -32768 to 32767
// int32 : -2147483648 to 2147483647
// int64 : -9223372036854775808 to 9223372036854775807
// uint : 0 to 18446744073709551615
// uint8 : 0 to 255
// uint16 : 0 to 65535
// uint32 : 0 to 4294967295
// uint64 : 0 to 18446744073709551615
// uintptr : an unsigned integer to store the uninterpreted bits of a pointer value
// byte // alias for uint8
// rune // alias for int32
// float32 : 32-bit floating-point numbers
// float64 : 64-bit floating-point numbers
// complex64 : complex numbers which have float32 real and imaginary parts
// complex128 : complex numbers with float64 real and imaginary parts
//-------------------------------
