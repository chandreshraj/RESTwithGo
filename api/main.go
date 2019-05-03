package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":8081"
	router := NewRouter()
	fmt.Println("listing to port ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
