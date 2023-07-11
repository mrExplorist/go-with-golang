package main

import (
	"fmt"
	"sync"
)



func main(){
	fmt.Println("Channels in Golang")

	// Creating a channel in Golang 
	myCh := make(chan int, 2)

	wg := &sync.WaitGroup{}

	// Pushing data into the channel

	// myCh <- 42
	// // print the data from the channel

	// fmt.Println(<-myCh)  

	wg.Add(2)



	//!Receive only channel 
	go func(ch <-chan int, wg *sync.WaitGroup){ // ch <-chan int means we can only read from the channel 

		// close(myCh)
		val, isChannelOpen := <-myCh
		if isChannelOpen{
			fmt.Println(val)
		}
		// Reading data from the channel
		// fmt.Println(<-myCh)
		

		wg.Done()
	} (myCh, wg)

	//! Send only channel
	go func(ch chan<- int, wg *sync.WaitGroup){ // ch chan<- int means we can only write into the channel
		// Pushing data into the channel
		
		// myCh <- 0
		myCh <- 42
		// myCh <- 27
		// Closing the channel
		close(myCh)
		
		wg.Done()
	} (myCh, wg)
	wg.Wait()	





	

}