package main

import (
	"fmt"
	"net/http"
	"sync"
)

////go greeter("Hello World!") // go routine is created here and it will run in the background and the main function will continue to run and exit the program without waiting for the go routine to finish execution and hence the go routine will not be executed at all as the main function will exit before the go routine can execute the function call to greeter function and hence the output will be empty and the program will exit without printing anything on the console screen

var signals = []string{"test"}



var wg sync.WaitGroup // create a wait group variable to wait for the go routines to finish execution before the main function exits the program 

func main() {

	// Create a slice of urls
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.microsoft.com",
		"https://www.apple.com",
		"https://www.stackoverflow.com",
		"https://www.udemy.com",
		"https://www.udacity.com",
		"https://www.coursera.com",

	}
	// Iterate over the slice of urls

	for _, url := range urls {
		// Call getStatusCode function
		go getStatusCode(url)
		wg.Add(1) // increment the wait group counter by 1 for each go routine that is created here to avoid the main function from exiting before the go routines finish execution
	}

	wg.Wait() // wait for all the go routines to finish execution before the main function exits the program
	fmt.Println(signals)



	// create a goRoutine

	// Call greeter function
	// go greeter("Hello World!") 
	// greeter("Hello Again!")


}

// func greeter(s string) {
// 	// Print string 5 times
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond) // sleep for 100 milliseconds before printing the string again to the console screen 
// 		fmt.Println(s)
// 	}

// }


func getStatusCode(url string) {
	defer wg.Done() // decrement the wait group counter by 1 for each go routine that finishes execution to avoid the main function from waiting for the go routines to finish execution before exiting the program 

	res , err := http.Get(url)

	if err != nil {
		fmt.Println("OOPS! Something went wrong with the connection to the server and the connection to the server could not be established properly")


	} else {
		signals = append(signals , url)
		fmt.Printf("%d status code returned from the server for endpoint %s\n" , res.StatusCode , url)
	}
	

	

	
}