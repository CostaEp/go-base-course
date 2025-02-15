# Basics-Func - Learning Functions in Go

## Overview
In this section, we explore how to use **functions** in Go while implementing a basic **IMT (Index of Mass of the Body) calculator**. We progressively enhance our code to make it more modular and reusable.

## Learning Objectives
By the end of this section, you will:
- Learn how to define and use **functions** in Go.
- Understand function **parameters** and **return values**.
- Implement modular and reusable code.

## Folder Structure
```
basics-func/
├── 01-main.go    # Introduces functions with a simple example
├── 02-main.go    # Expands function usage with return values
└── README.md     # Documentation for this section
```

## 01-main.go - Introducing Functions
This file demonstrates how to define and use a simple function in Go.

### Key Concepts:
- Declaring a function that accepts parameters.
- Using `fmt.Sprintf()` to format output.
- Calling a function from `main()`.

### Code Explanation:
```go
// resultIMT function takes a float64 value and prints the result
func resultIMT(imt float64) {
    result := fmt.Sprintf("your IMT is : %.2f", imt)
    fmt.Print(result)
}
```
- The function `resultIMT(imt float64)` takes a **float64** parameter and prints a formatted result.
- It is called in `main()` to display the calculated IMT.

### Running the Program:
```sh
go build -o 01_binary 01-main.go
./01_binary
```

## 02-main.go - Expanding Functions with Return Values
This file introduces functions that **return values**, making our code more modular and reusable.

### Key Concepts:
- Defining a function that **returns a value**.
- Using multiple parameters.
- Separating logic into **functions**.

### Code Explanation:
```go
// calcIMT function takes two float64 values and returns the IMT result
func calcIMT(userHeight float64, userWeight float64) float64 {
    const IMTPower = 2
    IMT := userWeight / math.Pow(userHeight, IMTPower)
    return IMT
}
```
- `calcIMT(userHeight, userWeight)` calculates IMT and returns the result.
- `resultIMT(IMT)` is called separately to display the result.

### Running the Program:
```sh
go build -o 02_binary 02-main.go
./02_binary
```

## Summary
In this section, we:
- **Introduced functions** in Go.
- Created a **resultIMT** function to display results.
- Developed a **calcIMT** function that returns a value.
- Learned how to **structure code better** using functions.

Next, we apply these concepts in the **final-imt-app-func** folder!