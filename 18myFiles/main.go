package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// files in golang
// 1. Create a file
// 2. Open a file
// 3. Write to a file
// 4. Read from a file
// 5. Close a file
// 6. Delete a file
// 7. Rename a file
// 8. Append to a file
// 9. Truncate a file
// 10. Seek to a specific location in a file


func main(){
	fmt.Println("Welcome to files in golang")



	content := "This needs to go in a file - https://go.dev/doc/effective_go"

	// create a file
	file, err := os.Create("./myGoFile.txt")

	// error handling
	checkNilError(err)
	// writing file
	length, err := io.WriteString(file, content)

	// error handling
	checkNilError(err)

	// printing length of file
	fmt.Println("Length is: ", length)

	// closing file
	defer file.Close()

	// reading file
	readFile("./myGoFile.txt")

}


// Reading from a file



func readFile(filename string){
 // reading a file

  databyte , err := ioutil.ReadFile(filename)

  checkNilError(err)

//   fmt.Println("Text from file: ", databyte)

  fmt.Println("Text from file: ", string(databyte))
}



func checkNilError(err error){
	if err != nil {
		panic(err)
	}
}