# IMT Calculator in Go

## Overview
This is a simple Go program that calculates the **Body Mass Index (IMT)** based on a user's height and weight. It demonstrates variable declaration, type conversion, and basic mathematical operations using Go's built-in `math` package.

## Features
- Declares and initializes variables using the `var` keyword.
- Demonstrates type conversion.
- Uses the `math.Pow()` function to calculate the square of a number.
- Outputs the calculated IMT using `fmt.Printf()`.

## Code
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// Declare and initialize variables:
	var userHeight float64 = 1.7
	var userWeight float64 = 100

	// Calculate IMT using the formula: weight / (height^2)
	var IMT = userWeight / math.Pow(userHeight, 2)

	// Print the result
	fmt.Printf("IMT is : %f", IMT)
}
```

## Installation and Running
### Prerequisites
- Go installed on your system ([Download Go](https://go.dev/dl/))

### Running the program
1. Save the code in a file named `main.go`.
2. Open a terminal and navigate to the directory containing `main.go`.
3. Run the command:
   ```sh
   go run main.go
   ```
4. The program will output the IMT value.

## Go Variable Types
The program also provides a reference for Go's variable types:

| Type        | Description |
|------------|-------------|
| `bool`     | `true`, `false` |
| `string`   | String values like `"Hello, World!"` |
| `int`      | Integer values like `42` |
| `int8`     | -128 to 127 |
| `int16`    | -32,768 to 32,767 |
| `int32`    | -2,147,483,648 to 2,147,483,647 |
| `int64`    | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 |
| `uint`     | 0 to 18,446,744,073,709,551,615 |
| `uint8`    | 0 to 255 |
| `uint16`   | 0 to 65,535 |
| `uint32`   | 0 to 4,294,967,295 |
| `uint64`   | 0 to 18,446,744,073,709,551,615 |
| `uintptr`  | Unsigned integer storing pointer values |
| `byte`     | Alias for `uint8` |
| `rune`     | Alias for `int32` (Unicode character representation) |
| `float32`  | 32-bit floating-point numbers |
| `float64`  | 64-bit floating-point numbers |
| `complex64`| Complex numbers with `float32` real and imaginary parts |
| `complex128` | Complex numbers with `float64` real and imaginary parts |

## License
This project is open-source and available under the MIT License.

