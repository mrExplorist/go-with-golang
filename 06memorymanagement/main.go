package main

//~------------------> memory management in Go! <------------------~//

func main(){
	
	//  ways to allocate memory in Go
	// 1. new() function
	// 2. make() function
	// 3. GarbageCollection()





	//* new(): ---------> * * * * * * * * * * * * * *
	// allocates memory and returns a pointer to it
	// zero value is returned
	// used for value types like int, float, struct

	//* make() --------------> * * * * * * * * * * *
	// creates slices, maps and channels only
	// returns an initialized (not zero) value of type T (not *T)
	// used for reference types like slices, maps and channels

	//* GarbageCollection() ---------> * * * * * * * * * *
	// Go has a garbage collector that runs in the background and cleans up memory that is no longer being used
	// Garbage collector is a part of the runtime
	// Garbage collector is a mark and sweep algorithm
	// Garbage collector is non deterministic
	// Garbage collector is not exposed to the user
	// Garbage collector is not a part of the language  specification
	// Garbage collector is not a part of the core language
	// Garbage collector is not a part of the compiler

	//& why GC is needed?
	// 1. to free up memory
	// 2. to avoid memory leaks
	// 3. to avoid fragmentation of memory
	// 4. to avoid running out of memory
	// 5. to avoid having to manually clean up memory

	
}