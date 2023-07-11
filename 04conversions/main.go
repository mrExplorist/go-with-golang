package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)




func main(){




	fmt.Println("Welcome to our pizza app!")
	fmt.Println("Please rate our pizza between 1 and 5: ")



   // Create a new reader, assuming that the default Stdin is a terminal
	reader := bufio.NewReader(os.Stdin)


	input , _:= reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)

	fmt.Println("Type of this rating is %T", input)


numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)



if err != nil{
	fmt.Println(err)
}else{
	fmt.Println("Added 1 to the rating" , numRating+1)
}




}