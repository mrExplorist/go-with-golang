package main

import (
	"fmt"
	"net/url"
)

// Handling URLs in Golang
// 1. Parse a URL
// 2. Build a URL
// 3. Query a URL
// 4. Encode a URL
// 5. Decode a URL
// 6. Parse a URL with Query Parameters
// 7. Parse a URL Path
// 8. Parse a URL Fragment
// 9. Parse a URL Host
// 10. Parse a URL Port
// 11. Parse a URL User


const myURL = "https://lco.dev:3000/learn?course=reactjs&paymentid=1234#learn"


func main(){
	fmt.Println("Welcome to handling URLs in golang")

	fmt.Println("URL is: ", myURL)


	// parsing a URL

	// function takes a string as input and returns a pointer to a URL struct
	result, _ := url.Parse(myURL) 

	// fmt.Println("Scheme is: ", result.Scheme)
	// fmt.Println("Host is: ", result.Host)
	// fmt.Println("Path is: ", result.Path)
	// fmt.Println("Port is: ", result.Port())
	fmt.Println("RawQuery is: ", result.RawQuery)

	// Querying a URL	
	queryParams := result.Query()

	fmt.Printf("Query params are: %T\n", queryParams)
	fmt.Println("Query params are: ", queryParams)
	fmt.Println("Query params are: ", queryParams["course"])
	// using range to print out all the query params

	for key, value := range queryParams{
		fmt.Println("Key is: ", key, "Value is: ", value)
	}

	// building a URL

	// creating a URL struct
	partsOfURL := &url.URL{ // pointer to a URL struct that will be returned by the function 
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=abc",
	}

	// CONSTRUCTING A URL FROM THE URL STRUCT USING THE STRING METHOD

	anotherURL := partsOfURL.String()

	fmt.Println("Another URL is: ", anotherURL)


}