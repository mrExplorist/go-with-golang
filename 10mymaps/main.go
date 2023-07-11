package main

import "fmt"




func main(){
	fmt.Println("Welcome to maps in Go!")


	// ~-------------------> creating maps in Go *****************

	// syntax
	// var mapName map[keyType]valueType
	// mapName = make(map[keyType]valueType)

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println(languages)
	// map[GO:Golang JS:Javascript PY:Python RB:Ruby]
    // Get the value of a particular key
	fmt.Println(languages["RB"])
	// Ruby


	// delete a key value pair from a map
	delete(languages, "JS")
	fmt.Println(languages)




	// Looping over a map
	// Some things about loops in Go
	// 1. Go only has one looping construct - the for loop
	// 2. Go does not have while or do while loops
	// 3. However, we can emulate while and do while loops using for loops
	// 4. for loop syntax
	// for initialisation; condition; post {}
	// 5. for is also used to iterate over collections like arrays, slices, maps, etc
	// 6. for range syntax
	// for key, value := range collection {}


	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// For key GO, value is Golang
	// For key PY, value is Python
	// For key RB, value is Ruby
	



// Flow of above for loop
// 1. key, value := range languages -> key = "GO", value = "Golang" (first iteration) 
// 2. key, value := range languages -> key = "JS", value = "Javascript" (second iteration)
// 3. key, value := range languages -> key = "PY", value = "Python" (third iteration)
// 4. key, value := range languages -> key = "RB", value = "Ruby" (fourth iteration)







	
	
}