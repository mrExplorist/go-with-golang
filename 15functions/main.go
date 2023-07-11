package main

import "fmt"



func main(){
	fmt.Println("Welcome to functions in Go!")
    greeting := greeter("John")
	fmt.Println(greeting)

	sum := adder(3, 5)
	fmt.Println(sum)


	proSum := proAdder(1, 2, 3, 4, 5)
	fmt.Println(proSum)


}


// Syntax of defining functions in Golang

// func funcName(param1 type, param2 type) returnType {
//        logic
// 	      return value
// }



func greeter(name string) string {
	return "Hello " + name
}


func proAdder(values ...int) int { 
	// (values ...int)/// //all the values are of type int
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
func adder(x int, y int) int {
	return x + y
}




