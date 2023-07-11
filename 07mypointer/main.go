package main

import "fmt"

// what is pointer?
// pointer is a variable that stores the memory address of another variable
// & is used to get the address of a variable
// * is used to get the value of a variable that the pointer points to
// * is also used to declare a pointer variable
// * is also used to dereference a pointer variable



func main (){


	fmt.Println("Welcome to pointers in Go!")




	// declaring a pointer
	var myIntPointer *int
	fmt.Println("My int pointer is ", myIntPointer)

	// returning nill -> is a special value in Go that means nothing or zero
	// nill is the default value of a pointer



	// declaring a variable
	myInt := 42
	var ptr = &myInt


	// pointer is a reference to the memory address of a variable
	fmt.Println("Memory address to which ptr is locating to :", ptr)

	// dereferencing a pointer
	fmt.Println("Value of actual pointer is ", *ptr)

	// changing the value of a variable through a pointer
	*ptr = *ptr + 2
	fmt.Println("New Value of myInt is ", myInt)
}