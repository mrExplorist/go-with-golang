package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)




func main(){


	fmt.Println("Hello World! Learning Mod in Go!")

	greeter()


	// mux.NewRouter returns a router that routes requests to the specified endpoint 

	r := mux.NewRouter()  
	// r.HandleFunc("/", HomeHandler) // this is the handler function for the route "/"
    r.HandleFunc("/", HomeHandler).Methods("GET")

	// run this router on the server
	log.Fatal(http.ListenAndServe(":8080", r))

}



func greeter()  {
	fmt.Println("Hello World! Greetings from Go!")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World from HomeHandler</h1>"))
}

