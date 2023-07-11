package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)


func main(){
	fmt.Println("Welcome to GoLang Server")
	PerformGetRequest()
}

// function to handle the requests


func PerformGetRequest(){
	const url = "https://jsonplaceholder.typicode.com/posts/1"
	response, err := http.Get(url)
	if err != nil{
		panic(err)
	}

	defer response.Body.Close()

	//status code
	fmt.Println("Status Code: ", response.StatusCode)
	// header
	fmt.Println("Header: ", response.Header)
	// content length
	fmt.Println("Content Length: ", response.ContentLength)
	// length of json data
	fmt.Println("Length: ", len(response.Header))
	// // proto
	// fmt.Println("Proto: ", response.Proto)
	// request
	fmt.Println("Request: ", response.Request)


	// create builder to build the request body and header fields 
	responseString := strings.Builder{}

	//read the response body 
	content, err := ioutil.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}

	byteCount,_ := responseString.Write(content)
	fmt.Println("Byte Count: ", byteCount)
	fmt.Println("Response String: ", responseString.String())




	// fmt.Println("Content: ", string(content))
	// fmt.Printf("Content: %s\n", content)


}

