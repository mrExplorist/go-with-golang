package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("Welcome to switch case in Go and we are going to make simple Dice game! ")

	rand.Seed(time.Now().UnixNano())
	// rand.Seed(time.Now().UnixNano()) is used to generate random number everytime we run the program

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spots")
		fallthrough
	case 5:
		fmt.Println("You can move 5 spots")
		
	case 6:
		fmt.Println("You can move 6 spots and roll the dice again")
	}


}