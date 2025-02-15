# Go Hello World

This is a simple Go program that prints a message to the console.

## Description

The program demonstrates basic usage of the `fmt` package to print output to the console.

## Code

```go
// package main is the entry point of the Go program
package main

// import fmt package to use its functions
import "fmt"

// main function is the entry point of the Go program
func main() {
    fmt.Printf("hello , its my first line code in go language \n")
    fmt.Printf("fmt.Print allow us to print something to the console")
}
```

## How to Build and Run

### Initialize the Go Module

Before building the application, you need to initialize a Go module:

```sh
go mod init <module_name>
```

### Build the Application

To compile the program, run:

```sh
go build
```

This will generate an executable file in the current directory.

### Run the Application

You can run the compiled program using:

```sh
./<module_name>
```

### Expected Output

```sh
hello , its my first line code in go language
fmt.Print allow us to print something to the console
```

### Run Without Building

If you want to run the application without compiling it first, use:

```sh
go run main.go
```

## Requirements

- Go installed on your system (check installation with `go version`).
- A terminal or command prompt to run commands.

## License

This project is open-source and available for learning purposes.

## Author

[CostaEp]

