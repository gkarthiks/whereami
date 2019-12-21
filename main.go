package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main(){
	fmt.Println("Starting the server... ")
	http.HandleFunc("/", whereami)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func whereami(w http.ResponseWriter, r *http.Request) {
	name, _ := os.Hostname()
	fmt.Fprintf(w, "Serving request from %s!", name)
}