package main

import "fmt"




func main(){
	fmt.Println("Welcome to arrays in Go!")
	// declaring an array
	var fruitList [4]string





	// declaring an array with initial values
	// another method to store in array
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Grape"

	fmt.Println("Fruit List is ", fruitList)
	fmt.Println("Fruit List is ", len(fruitList))


	// declaring and initializing an array in one line
	veggieList := [3]string{"Potato", "Tomato", "Brinjal"}
	fmt.Println("Veggie List is ", veggieList)
	fmt.Println("Veggie List is ", len(veggieList))
}