# Lab-02: Introduction to Variables, Constants, and IMT Calculation in Go

## Overview
In **Lab-02**, we explore how to declare and use **variables**, **constants**, and perform **calculations** in Go. The goal of this lab is to gain a deep understanding of Go's variable types while building a simple **IMT (Index of Mass of the Body) calculator**.

## Learning Objectives
By the end of this lab, you will:
- Learn how to declare **variables** using `var` and `:=`.
- Understand how to use **constants** with `const`.
- Explore different **data types** in Go.
- Perform calculations using **math functions**.
- Accept **user input** and process it in a Go program.

## Folder Structure
```
lab-02/
├── basics/           # Learning how to use variables and constants
│   ├── 01-main.go    # Declaring variables explicitly
│   ├── 02-main.go    # Using constants and shorthand variable declaration
│   ├── README.md     # Documentation for basics folder
│
├── final-imt-app/    # The complete IMT calculator
│   ├── main.go       # Final version of IMT calculator
│   ├── README.md     # Documentation for final project
│
└── README.md         # Documentation for the entire lab
```

## What We Will Learn
### 1. Basics of Variables and Constants (Folder: `basics/`)
- **01-main.go**: Learn how to declare variables explicitly and perform calculations.
- **02-main.go**: Learn how to declare constants and use shorthand notation (`:=`).
- Understanding **data types** and **type conversions**.

### 2. Building a Complete Application (Folder: `final-imt-app/`)
- Taking **user input** using `fmt.Scan()`.
- Implementing the **IMT formula** (`weight / height²`).
- Displaying results with `fmt.Printf()`.
- Compiling and running a **standalone Go application**.

## Running the Programs
### Running Basics Code
To compile and run the basics programs:
```sh
# Navigate to the basics directory
cd basics

# Build and run 01-main.go
go build -o 01_binary 01-main.go
./01_binary

# Build and run 02-main.go
go build -o 02_binary 02-main.go
./02_binary
```

### Running Final IMT Application
To run the full IMT application:
```sh
# Navigate to the final-imt-app directory
cd final-imt-app

# Initialize Go module (only needed once per project)
go mod init final-imt-app

# Build and run the application
go build -o final_binary main.go
./final_binary
```

## Summary
In **Lab-02**, we:
- Learned **variable and constant declaration** in Go.
- Explored **data types** and **type conversion**.
- Built a real-world **IMT calculator**.
- Used **user input** and **math functions** in Go.

This lab provides a solid foundation for working with **variables, constants, and basic operations** in Go. Next, we will explore **control structures** (if-else, loops) to add more logic to our programs!
