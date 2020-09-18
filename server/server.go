package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", homePageHandler)

	fmt.Println("Server listening on Port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello from the other side")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
