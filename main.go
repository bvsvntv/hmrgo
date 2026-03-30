package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	visitHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hmrgo: understanding hot module replacement")
	}

	http.HandleFunc("/", visitHandler)

	log.Println("Serving on http://localhost:2002")
	if err := http.ListenAndServe(":2002", nil); err != nil {
		log.Fatal(err)
	}
}
