package main

import "fmt"

// Defer in Golang?
// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.



func main(){
	// ~ --------->  Creating a defer **********************

	defer fmt.Println("This is first defer statement")
	defer fmt.Println("This is second defer statement")
	defer fmt.Println("This is third defer statement")

	fmt.Println("Welcome to defer in golang")
	myDefer()


	// executed like LIFO
}



func myDefer(){
	for i := 0; i < 5; i++ {

		defer fmt.Println(i)
		
	}

}
