package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	port := flag.String("port", "2002", "port to listen on")
	host := flag.Bool("host", false, "expose on LAN")
	flag.Parse()

	fileServer := http.FileServer(http.Dir("."))
	http.Handle("/", fileServer)

	bind := "127.0.0.1"
	if *host {
		bind = "0.0.0.0"
	}
	addr := net.JoinHostPort(bind, *port)

	fmt.Printf("] Local: http://localhost:%s\n", *port)
	if *host {
		fmt.Printf("] Network: <show_local_ip>:%s\n", *port)
	}

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
