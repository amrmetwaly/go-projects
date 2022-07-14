package main

import (
	"fmt"
	"log"
	"net/http"
)

// w is what the server responds with
// r is the request from the user (client)
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "mehtod is not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "hello!")

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseFrom() err: %v", err)
	}

	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

}

func main() {
	// Serve the files in the directory called static
	fileServer := http.FileServer(http.Dir("./static"))
	// Send the / route requests to the fileServer
	http.Handle("/", fileServer)
	//any incoming requests to /form should call formHandler
	http.HandleFunc("/form", formHandler)
	//any incoming requests to /hello should call helloHandler
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server on port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}