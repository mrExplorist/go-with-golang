// Post request is similar to get request. The only difference is that we use http.Post() method instead of http.Get() method.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main(){
	fmt.Println("Welcome to GoLang Server ")

	PerformPostJsonRequest()
}


func PerformPostJsonRequest(){

	requestBody := strings.NewReader(`{"id": 1, "title": "foo", "body": "bar", "userId": 1}`)




	const url = "https://jsonplaceholder.typicode.com/posts"
	response, err := http.Post(url, "application/json",requestBody )
	if err != nil{
		panic(err)
	}


	defer response.Body.Close()

	content,_ := ioutil.ReadAll(response.Body)

	// Response Status Code
	fmt.Println("Status Code: ", response.StatusCode)
	// // Response Header
	// fmt.Println("Header: ", response.Header)

	// Response Body
	fmt.Println("Body: ", string(content))
	// Response Content Length
	fmt.Println("Content Length: ", response.ContentLength)
}