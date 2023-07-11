package main

import (
	"fmt"
	"sync"
)


func main(){
	fmt.Println("Race Condition in - GoLang") 

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{} // mutex is used to lock the variable so that only one go routine can access it at a time

	// why locking is required? 
	// because if two go routines try to access the same variable at the same time then the value of the variable will be unpredictable 



	var score = []int{0} 


	wg.Add(3)



	// wait group is used to wait for all go routines to finish before exiting the main function 
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("Go Routine 1")
		mut.Lock() // lock the variable so that only one go routine can access it at a time 
		score = append(score, 1)
		mut.Unlock() // unlock the variable so that other go routines can access it
		wg.Done()
		

	}(wg, mut) // go routine
	go func(wg *sync.WaitGroup , m *sync.RWMutex){
		fmt.Println("Go Routine 2")
		mut.Lock() 
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup , m *sync.RWMutex){
		fmt.Println("Go Routine 3")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)



	 wg.Wait() // wait for all go routines to finish
	mut.RLock()
	fmt.Println(score)
	mut.RUnlock()




}

