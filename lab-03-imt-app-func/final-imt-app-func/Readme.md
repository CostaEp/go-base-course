# Final IMT App Func - A Structured Approach

## Overview
In this section, we build a **fully modular IMT (Index of Mass of the Body) calculator** using **functions** in Go. This project introduces better **code organization**, **parameter handling**, and **function return values**.

## Learning Objectives
By the end of this section, you will:
- Use **functions** to make code modular and reusable.
- Handle **multiple return values** in Go.
- Improve **code readability** by structuring logic efficiently.

## Folder Structure
```
final-imt-app-func/
├── main.go     # The main Go application using functions
├── README.md   # Documentation
└── go.mod      # Go module file (if needed)
```

## Task - Building the Final IMT Application with Functions
### Step 1: Understanding the Code
The final application prompts the user for their height, weight, and name, then calculates and displays their IMT using **functions**.

#### Code Breakdown
```go
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
```
- This function **prompts the user** for input and returns **three values**.
- The return order follows the function signature.

```go
// calcIMT function takes two float64 values and returns the result
func calcIMT(userHeight float64, userWeight float64) float64 {
    IMT := userWeight / math.Pow(userHeight, IMTPower)
    return IMT
}
```
- This function **calculates the IMT** and **returns the result**.
- **IMTPower** is used as a **constant** to maintain consistency.

```go
// resultIMT function now takes both IMT value and userName
func resultIMT(imt float64, userName string) {
    result := fmt.Sprintf("%s, your calculated IMT is: %.2f", userName, imt)
    fmt.Print(result)
}
```
- The **IMT result** is now personalized with the user’s **name**.

### Step 2: Running the Application
To execute the final IMT application:
```sh
# Initialize a Go module (only needed once per project)
go mod init final-imt-app-func

# Build the application
go build -o finalimt_binary main.go

# Run the application
./finalimt_binary
```

## Summary
In this section, we:
- Used **functions** to organize our IMT application better.
- Implemented **multiple return values** for better input handling.
- **Improved readability** and **maintainability** of our Go program.

This structured approach makes our **IMT calculator scalable and reusable**. 🚀 Next, let’s summarize the entire lab!
