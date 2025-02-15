# Lab-03: IMT Calculator with Functions

## Overview
In **Lab-03**, we improve our IMT (Index of Mass of the Body) calculator by implementing **functions** in Go. By refactoring our code to use functions, we make it more **structured, modular, and reusable**.

## Learning Objectives
By the end of this lab, you will:
- Understand the importance of **functions** in Go.
- Learn how to **pass parameters** and **return values**.
- Use functions to **organize and improve** code structure.

## Folder Structure
```
lab-03-imt-app-func/
├── basics-func/           # Learning how to use functions in Go
│   ├── 01-main.go         # Introduces basic functions
│   ├── 02-main.go         # Expands function usage with return values
│   ├── README.md          # Documentation for basics-func
│
├── final-imt-app-func/    # The complete IMT calculator using functions
│   ├── main.go            # Final version of IMT calculator with modular functions
│   ├── README.md          # Documentation for final-imt-app-func
│
└── README.md              # Documentation for the entire lab
```

## What We Will Learn
### 1. Basics of Functions (Folder: `basics-func/`)
- **01-main.go**: Introduces basic functions to modularize the IMT calculation.
- **02-main.go**: Expands on function usage by returning values.

### 2. Building a Complete Application (Folder: `final-imt-app-func/`)
- Taking **user input** using `fmt.Scan()`.
- Implementing **modular functions** to structure the application.
- Returning multiple values from functions.

## Running the Programs
### Running Basics Code
To compile and run the basics programs:
```sh
# Navigate to the basics-func directory
cd basics-func

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
# Navigate to the final-imt-app-func directory
cd final-imt-app-func

# Initialize Go module (only needed once per project)
go mod init final-imt-app-func

# Build and run the application
go build -o finalimt_binary main.go
./finalimt_binary
```

## Summary
In **Lab-03**, we:
- Explored **functions** and their importance in Go.
- Refactored our **IMT calculator** using **modular functions**.
- Improved **code readability** and **maintainability**.

This lab provides a strong foundation for **structured programming** in Go. Next, we will explore **control structures** and **loops** to enhance our applications!
