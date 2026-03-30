package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "2002", "port to listen on")
	flag.Parse()

	fileServer := http.FileServer(http.Dir("."))
	http.Handle("/", fileServer)

	fmt.Printf("] Serving on http://localhost:%s\n", *port)
	if err := http.ListenAndServe(":2002", nil); err != nil {
		log.Fatal(err)
	}
}
