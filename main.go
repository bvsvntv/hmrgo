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
		fmt.Printf("] Network: http://%s:%s\n", getOutboundIP(), *port)
	}

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

// Get preferred outbound ip of this machine
// reference: https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
