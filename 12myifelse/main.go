package main

import "fmt"

func main(){
	fmt.Println("Welcome to if else in Go!")


	// ~ --------->  if else **********************
	loginCount := 21
	
  var result string



	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}
    fmt.Println(result)


	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}


	// ~ --------->  if else with assignment **********************


	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is not less than 10")
	}

	//if num := 3; what is this? this is called initialization statement
	// num < 10 is called condition statement
	



}