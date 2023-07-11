package main

import (
	"fmt"
	"time"
)


func main() {
	fmt.Println("Welcome to the time in Go!")



	presentTime := time.Now()
	fmt.Println("Present Time is ", presentTime)
	// Formatted time
	fmt.Println("Formatted time is ", presentTime.Format("01-02-2006 15:04:05 Monday"))




	// Creating a date manually 
	createdDate := time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC)
	
	fmt.Println("Created date is ", createdDate)

	fmt.Println("Created date is ", createdDate.Format("01-02-2006 Monday"))





}