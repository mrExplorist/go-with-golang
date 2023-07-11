package main

import (
	"encoding/json"
	"fmt"
)

// type User struct{
// 	Id int `json:"userid"`
// 	Name string `json:"name"`
// 	Email string `json:"email"`
// 	Phone string `json:"phone"`
// 	// Password string `json:"-"`
// 	Website string `json:"website"`
// 	Tags []string `json:"tags"`
// }


func main(){
	fmt.Println("Welcome! We are dealing with JSON in GoLang")

	//  EncodeJson()
	DecodeJson()


}



// Encode and Decode JSON in GoLang 

// func EncodeJson(){
// 	// encode json data in

// 	curiousUsers := []User{
// 		{1 , "foo", "foo@example.com" , "1234567890" , "www.example.com" , []string{"foo", "bar"}},
// 		{2 , "foo2", "foo@example2.com" , "1234567890" , "www.example2.com" , []string{"foo2", "bar2"}},
// 		{3 , "foo3", "foo@example3.com" , "1234564590" , "www.example3.com" , []string{"foo3", "bar3"}},
// 		{4 , "foo4", "foo@example4.com" , "123456234590" , "www.example4.com" , []string{"foo4", "bar4"}},	
// 	}


// 	// package the data as JSON data

// 	// finalJsonData, err := json.Marshal(curiousUsers)
// 	// json.MarshalIndent() is used to indent the JSON data for better readability and debugging purpose in the browser or terminal 


// 	finalJsonData, err := json.MarshalIndent(curiousUsers, "", " ")

// 	if err != nil{
// 		panic(err)
// 	}

// 	// print the JSON data
// 	fmt.Println(string(finalJsonData))

// }


type Person struct {
	ID    int    `json:"userid"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}


// consume json data then decode JSON data in GoLang

func DecodeJson(){
	// JSON data to be consumed

	// get json data from web jsonplaceholder.typicode.com


	// data from web is the form of bytes that will be converted to string and then to JSON data in GoLang format and then unmarshalled to GoLang struct type
	var jsonDataFromWeb = []byte(`
[
  {
    "id": 1,
    "name": "John Doe",
    "age": 30,
    "email": "johndoe@example.com"
  },
  {
    "id": 2,
    "name": "Jane Smith",
    "age": 25,
    "email": "janesmith@example.com"
  },
  {
    "id": 3,
    "name": "Bob Johnson",
    "age": 35,
    "email": "bobjohnson@example.com"
  }
]
`)



	// validating the JSON data before unmarshalling it to GoLang struct type 

    var people []Person

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid{

		fmt.Println("JSON data is valid")
		// unmarshal the JSON data

		json.Unmarshal(jsonDataFromWeb, &people)

		fmt.Printf("%#v\n", people)



	}else{
		fmt.Println("JSON data is not valid")
	}




	// some cases where we just need to add data to key value pair 

	// var person map[string]interface{}

	// json.Unmarshal(jsonDataFromWeb, &person)

	// fmt.Printf("%#v\n", person)

	

	// Print the parsed data
	for _, person := range people {
		fmt.Printf("ID: %d, Name: %s, Age: %d, Email: %s\n", person.ID, person.Name, person.Age, person.Email)
	}


	}
	





