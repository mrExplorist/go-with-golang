package main

import "fmt"

// *--------->  Inheritance in golang? **********************

// Go does not have inheritance like other object oriented languages like Java, C++, etc ... so how do we achieve inheritance in Go?

// We use composition instead of inheritance in Go
// Composition is a way to combine different types together

// syntax
// type Parent struct {}
func main() {
	fmt.Println("Welcome to structs in Go!") 

	// ~ --------->  Creating a struct **********************
	
	lalit := User{"Lalit", "djf@gmail.com", true, 21}
	fmt.Println(lalit)
	fmt.Printf("lalit details are: %+v\n", lalit)
	fmt.Printf("Name is %v and Email is %v\n", lalit.Name, lalit.Email)

}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}


	




