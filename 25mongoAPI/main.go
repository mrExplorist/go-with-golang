package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/mrExplorist/mongoAPI/routes"
)

func main(){
	fmt.Println("MongoDB API...") 

	r := router.Router()
	fmt.Println("Server is getting started...") 
	log.Fatal(http.ListenAndServe(":3001", r ))
	fmt.Println("Server is listening on port 3001..")
}
