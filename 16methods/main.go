package main

import (
	"fmt"
)

// Difference between methods and functions in golang ?
// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// In this example, the Abs method has a receiver of type Vertex named v.





func main() {
	fmt.Println("Welcome to methods in Go!") 

	// ~ --------->  Creating a struct **********************

	
	
	lalit := User{"Lalit", "djf@gmail.com", true, 21}



	fmt.Println(lalit)
	fmt.Printf("lalit details are: %+v\n", lalit)
	fmt.Printf("Name is %v and Email is %v\n", lalit.Name, lalit.Email)


	// ~ --------->  Creating a method **********************
	lalit.GetStatus()
	lalit.NewMail()

}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}


// ~ --------->  Creating a method **********************
// *General syntax in golang for methods
// func (t Type) methodName(parameter list) {
	// code
// }

func (u User) GetStatus() {
	fmt.Println("GetStatus method is called on", u.Status)
}
func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email is changed to", u.Email)
}






