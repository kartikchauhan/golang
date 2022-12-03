package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const port = ":8080"

func printMethod(s string) {
	log.Printf("Method: %s", s)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	printMethod("pingHandler")

	fmt.Fprintf(w, "pong!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	printMethod("formHandler")

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "Parsing error %v", err)
		return
	}

	fmt.Fprintf(w, "Name: %s\n", r.FormValue("name"))
	fmt.Fprintf(w, "Country: %s\n", r.FormValue("country"))
}

func main() {
	printMethod("main")

	fileServer := http.FileServer(http.Dir("./static/"))

	http.Handle("/", fileServer)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/form", formHandler)

	log.Printf("Listening on port %s", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
