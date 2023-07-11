package main

import "fmt"


func main(){
	fmt.Println("Welcome to loops in Go!")



	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)
	// looping over slice using for loop

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}


	// looping over slice using range keyword
	for i, d := range days {
		fmt.Printf("index is %v and value is %v\n", i, d)
	}

	newValue := 1
	for newValue < 10 {
		if newValue == 5 {
			goto lsp
		}
		fmt.Println("Value is: ", newValue)
		newValue++
	}

	lsp:
	 fmt.Println("Jumping to lsp portfolio") // lsp: label statement

	



}