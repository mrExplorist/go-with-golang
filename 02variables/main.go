package main

import "fmt"

// why i should learn go?
// 1. Go is a statically typed language
// 2. Go is a compiled language
// 3. Go is a strongly typed language
// 4. Go is a garbage collected language
// 5. Go is a concurrent language
// 6. Go is a simple language
// 7. Go is a fast language
// 8. Go is a production ready language

//   var jwtToken := 3000030 // this is not allowed outside functions

const LoginToken string = "asdkjfl" // this is a package level variable
// package level variables are available to all the files inside the package
// package level variables are available to all the functions inside the package
func main() {
	var username string = "root"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)


	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)


	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)


	var smallFloat float32 = 258.4554453464384875
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)


	var smallFloat2 float64 = 258.4554453464384875
	fmt.Println(smallFloat2)
	fmt.Printf("Variable is of type: %T \n", smallFloat2)

	// Default values and some aliases
	var anotherVariable int // when no value is assigned, it is assigned the default value of 0 not some garbage value
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable) 


	// Type inference - Go can infer the type of the variable from the value assigned to it


	var implicitVar = "I am implicitly declared" // type is inferred from the value assigned
	


	fmt.Println(implicitVar)
	fmt.Printf("Variable is of type: %T \n", implicitVar)
	
//  implicitVar = 10 // this will throw an error because implicitVar is of type string and we are trying to assign an int value to it





// no var style
// This style is used only inside functions
// This style is not allowed outside functions
// This style is not allowed for package level variables
// This style is not allowed for global variables
// This style is not allowed for variables which are not initialize
// This style is not allowed for variables which are not used
// This style is not allowed for variables which are not declared
// This style is not allowed for variables which are not declared and initialized


	noVarStyle := "I am no Var styled" // type is inferred from the value assigned
	fmt.Println(noVarStyle)

	noVarStyle = "I am no Var styled and I am reassigned" // type is inferred from the value assigned

	fmt.Println(noVarStyle)


	// noVarStyle = 10 // this will throw an error because noVarStyle is of type string and we are trying to assign an int value to it


	numberOfUsers := 30098743
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type: %T \n", numberOfUsers)




	

}