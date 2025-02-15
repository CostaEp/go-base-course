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
├── main.go     # The main Go application
├── README.md   # Documentation
└── go.mod      # Go module file (if needed)
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

    fmt.Printf("This program calculates the IMT (Index of Mass of the Body)\n")
    fmt.Printf("Enter your height in meters (e.g., 1.6): ")
    fmt.Scan(&userHeight)

    fmt.Printf("Enter your weight in kilograms (e.g., 70): ")
    fmt.Scan(&userWeight)

    IMT := userWeight / math.Pow(userHeight, IMTPower)
    fmt.Printf("Your IMT is: %f\n", IMT)
}
```

### Step 2: How the Code Works
- **Takes user input** using `fmt.Scan(&variable)`.
- **Uses `math.Pow(userHeight, IMTPower)`** to calculate `height²`.
- **Performs IMT calculation** using `weight / height²`.
- **Displays the result** to the user.

### Step 3: Running the Application
To run the final IMT application:
```sh
# Initialize a Go module (only needed once per project)
go mod init final-imt-app

# Build the application
go build -o final_binary main.go

# Run the application
./final_binary
```

## Summary
In this section, we:
- Created a **fully functional IMT calculator**.
- Handled **user input** and performed **dynamic calculations**.
- Learned how to **build and execute** a Go program efficiently.

This marks the completion of our practical implementation. 🚀 Next, let's summarize our learnings from the entire lab!
