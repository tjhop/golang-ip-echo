package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	// 3rd party packages
	"github.com/julienschmidt/httprouter"
)

// based on:
// https://blog.golang.org/context/userip/userip.go
// and
// https://stackoverflow.com/a/33301173
func getIP(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// returns in <ip, port, err> format
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		//return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)

		fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		//return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
		fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
		return
	}

	// This will only be defined when site is accessed via non-anonymous proxy
	// and takes precedence over RemoteAddr
	// Header.Get is case-insensitive
	forward := req.Header.Get("X-Forwarded-For")

	fmt.Fprintf(w, "%s\n", forward)
}

func main() {
	// make new httprouter
	r := httprouter.New()

	// add handler for default path to retrieve IP addresses
	r.GET("/", getIP)

	// get socket to listen
	sock, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.Serve(sock, r))
}
