package main

import (
	"fmt"
	"sort"
)

// ~---------------> Slices in Go ! <----------------- //

// Slices are dynamic arrays in Go
// Slices are reference types
// Slices are like a window to the underlying array
// Slices are created using the make keyword
// Slices are used more often than arrays in Go




func main(){


	fmt.Println("Welcome to slices in Go!")


	// ----------> General slices syntax <----------- //

	// var sliceName []type = make([]type, length, capacity)


	// declaring and initializing a slice
	
	
	var fruitList = []string{"Apple", "Tomato", "Peach", "Banana"}
	// type of fruitlist
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fmt.Println("Fruit List is ", fruitList)
	fmt.Println("Fruit List is of size: ", len(fruitList))
	
	// adding more elements to a slice using append function
	fruitList = append(fruitList, "Mango", "Kiwi")

	fmt.Println("Fruit List after adding elements using append is ", fruitList)

	// slicing a slice
	// syntax: sliceName[startIndexIncluding : uptoNotIncluding]

	fmt.Println("Sliced FruitList is ", fruitList[1:3])
	fmt.Println("Sliced FruitList is with default initial startIndex", fruitList[:3])
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)
	

    // ~-------------------> declaring slices using make keyword *****************

	// var sliceName []type = make([]type, length, capacity)

	
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 // this will give an error as the capacity is 4 
	highScores = append(highScores, 555, 666, 777)



	fmt.Println("High Scores is ", highScores)
	fmt.Println("High Scores is of length ", len(highScores))
	

	// ~---------->  sort method -------------------- ****************
	// about sorting of slices in Go
	// sort package is used for sorting slices in Go
	// sort package implements sorting for int, float64, string slices
	// sort package uses quick sort for sorting
	// sort package implements 2 methods for sorting: sort and stable sort
	// sort method is used to sort slices of primitive types
	// stable sort method is used to sort slices of structs



	fmt.Println("Is sorted ? ", sort.IntsAreSorted(highScores))
	 sort.Ints(highScores)
	 fmt.Println("High Scores after sorting is ", highScores)
	 fmt.Println("Is sorted ? ", sort.IntsAreSorted(highScores))



	//  Removing elements from a slice using index

	// syntax: sliceName = append(sliceName[:index], sliceName[index+1:]...)


	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println("Courses are ", courses)


	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses after removing index 2 ", courses)





	



	 


	


  







	
















}