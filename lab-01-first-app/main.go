// package main is the entry point of the go program
package main

// import fmt package to use its functions
import "fmt"

// main function is the entry point of the go program
func main() {
	fmt.Printf("hello , its my first line code in go language \n")
	fmt.Printf("fmt.Print allow us to print something to the console")
}

// if we whant to build this app we need to initialize the go module by run this command:
// go mod init <module_name>
// -------------------------------------------
// then we can build the app by run this command:
// go build
// -------------------------------------------
// lets try this at the command line
// -------------------------------------------
//now we can run the app this command:
// ./<module_name>
// -------------------------------------------
// output:
// hello , its my first line code in go language
// fmt.Print allow us to print something to the console
// -------------------------------------------
// if we want to run the app without build it we can run this command:
// go run main.go
// -------------------------------------------
