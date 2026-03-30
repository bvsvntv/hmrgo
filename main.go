package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("."))
	http.Handle("/", fileServer)

	log.Println("Serving on http://localhost:2002")
	if err := http.ListenAndServe(":2002", nil); err != nil {
		log.Fatal(err)
	}
}
