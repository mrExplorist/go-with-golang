package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


const url = "https://www.lco.dev"



func main(){
	fmt.Println("Welcome to web requests in golang")

	// making a request to a website
	resp, err := http.Get(url)

	// error handling
	if(err != nil){
		panic(err)
	}

	// printing response
	fmt.Printf("Response is of type: %T\n", resp)

	// closing response body
	defer resp.Body.Close()



	// printing response status
	fmt.Println("Response status is: ", resp.Status)

	// reading response body data as bytes and sending it to a function to convert it to string and print it out to console 

	databytes, err := ioutil.ReadAll(resp.Body)

	// error handling
	if(err != nil){
		panic(err)
	}


	content := string(databytes)

	fmt.Println("Content is: ", content)














	
}