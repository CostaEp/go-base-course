# Basics - Learning Variables and Constants in Go

## Overview
In this folder, we focus on learning how to declare and use **variables** and **constants** in Go. We will explore different ways to initialize variables and understand their types using two Go programs.

## Learning Objectives
By the end of this section, you will:
- Understand how to declare variables using `var` and `:=`.
- Learn about **constants** and their role in Go.
- Explore various **data types** in Go.
- Learn how to perform basic calculations using the `math` package.

## Folder Structure
```
basics/
├── 01-main.go   # Demonstrates variable declaration and type conversion
├── 02-main.go   # Demonstrates constants and shorthand variable declaration
└── README.md    # Documentation for this section
```

## 01-main.go - Variable Declaration and Type Conversion
This file introduces basic **variable declaration** and **type conversion**.

### Key Concepts:
- Using `var` to declare variables explicitly with types.
- Performing basic calculations using `math.Pow()`.
- Printing output using `fmt.Printf()`.

### Code Explanation:
```go
var userHeight float64 = 1.7
var userWeight float64 = 100
var IMT = userWeight / math.Pow(userHeight, 2)
```
- `userHeight` and `userWeight` are declared as `float64` explicitly.
- `math.Pow(userHeight, 2)` calculates the square of the height.
- The **IMT formula** (`weight / height²`) is applied.
- The result is printed using `fmt.Printf("IMT is : %f", IMT)`.

### Running the Program:
```sh
go build -o 01_binary 01-main.go
./01_binary
```

## 02-main.go - Using Constants and Short Variable Declaration
This file introduces **constants** and the shorthand `:=` notation for variable declaration.

### Key Concepts:
- Using `const` to declare a constant.
- Using `:=` for concise variable declaration.
- Performing calculations using a constant.

### Code Explanation:
```go
const IMTPower = 2
userHeight := 1.7
var userWeight float64 = 100
IMT := userWeight / math.Pow(userHeight, IMTPower)
```
- `IMTPower` is a **constant** that holds the power value.
- `userHeight` is declared using `:=`, so Go automatically assigns its type.
- The **IMT calculation** is performed the same way as in `01-main.go`.
- The result is printed using `fmt.Printf("IMT is : %f", IMT)`.

### Running the Program:
```sh
go build -o 02_binary 02-main.go
./02_binary
```

## Summary
In this section, we explored:
- **Variable declaration** using `var` and `:=`.
- **Constants** using `const`.
- **Basic arithmetic operations** in Go.
- **Building and running** Go programs using `go build` and `./binary_name`.

Next, we will apply these concepts in a real-world example in the **final-imt-app** folder!