package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)


func main(){

	fmt.Println("Welcome to GoLang Server")

	PerformPostFormRequest()


}


func PerformPostFormRequest(){

	// formdata to be sent in the request body

	formData := url.Values{}
	formData.Set("name", "foo")
	formData.Set("surname", "bar")
	formData.Set("id", "1")
	formData.Set("email", "foo@bar.com")
	







	const url = "https://jsonplaceholder.typicode.com/posts"
	response, err := http.PostForm(url,formData )
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


