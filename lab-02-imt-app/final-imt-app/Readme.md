# Final IMT App - A Practical Implementation

## Overview
In this section, we build a **fully functional IMT (Index of Mass of the Body) calculator**. This project will demonstrate:
- User input handling.
- Performing calculations based on user input.
- Using **constants** and **variables** efficiently.

## Learning Objectives
By the end of this section, you will:
- Understand how to use **`fmt.Scan()`** to take user input.
- Apply the **IMT formula** dynamically.
- Build and run a complete Go application.

## Folder Structure
```
final-imt-app/
├── main.go       # The main Go application
├── README.md     # Documentation
├── final-imt-app # final-imt-app is a binary file (can be run on any machine)
└── go.mod        # Go module file (if needed)
```

## Task - Building the Final IMT Application
### Step 1: Understanding the Code
The final application prompts the user for their height and weight, then calculates and displays their IMT.

#### Code Breakdown
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    const IMTPower = 2
    var userHeight float64
    var userWeight float64

    // Using Println to print and go to the next line
    // We can use backticks to print multiple lines
    fmt.Println(`		---This program calculates---
	  ---the IMT (Index of Mass of the Body)---`)
    fmt.Print("Enter your height in meters (for example 1.6): ")
    fmt.Scan(&userHeight)

    fmt.Print("Enter your weight in kilograms (for example 70): ")
    fmt.Scan(&userWeight)

    IMT := userWeight / math.Pow(userHeight, IMTPower)
    // printf and %.2f to print only 2 decimal places
    // Using format you can go to Go documentation https://pkg.go.dev/fmt
    // If we want only to save the value, we can use Sprintf
    result := fmt.Sprintf("your IMT is : %.2f", IMT)
    fmt.Print(result)
}
```

### Step 2: How the Code Works
- **Takes user input** using `fmt.Scan(&variable)`.
- **Uses `math.Pow(userHeight, IMTPower)`** to calculate `height²`.
- **Performs IMT calculation** using `weight / height²`.
- **Formats the output** using `fmt.Sprintf("your IMT is : %.2f", IMT)` to limit decimal places.
- **Displays the result** to the user.

### Step 3: Running the Application
To run the final IMT application:
```sh
# Initialize a Go module (only needed once per project)
go mod init final-imt-app

# Build the application
go build

# Run the application
./final-imt-app
```

## Summary
In this section, we:
- Created a **fully functional IMT calculator**.
- Handled **user input** and performed **dynamic calculations**.
- Learned how to **build and execute** a Go program efficiently.

This marks the completion of our practical implementation. 🚀 Next, let's summarize our learnings from the entire lab!
