package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	welcome := "Welcome to the world of Go!"
	fmt.Println(welcome)


	// Taking input from user
	// var name string
	// fmt.Println("What is your name?")
	// fmt.Scanf("%s", &name)
	// fmt.Printf("Hi %s, I am Go! You are awesome!\n", name)
	
	// * using bufio package
	// what is bufio?
	// bufio stands for buffered input and output
	// bufio is a package in Go which allows us to take input from the user
	// bufio is a package in Go which allows us to write output to the terminal
	

	
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our Pizza: ")
	// comma ok syntax || err is the error object



	// COMMA ERROR SYNTAX
	 input , err := reader.ReadString('\n')
	 fmt.Println("Thanks for rating, ", input)

	 fmt.Printf("Type of this rating is %T", input)
	//~ why comma ok?
	// because we are returning two values from ReadString
	// first value is the actual string
	// second value is the error object
	// if there is no error, then the second value will be nil
	// if there is an error, then the second value will be an error object
	// if there is no error, then the first value will be the actual string
	// if there is an error, then the first value will be an empty string
	// if there is no error, then the comma ok will be true
	// if there is an error, then the comma ok will be false




	

}