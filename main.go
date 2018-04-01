// +build !appengine

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

var port = flag.Int("port", 55555, "server port")

func main() {
	log.SetFlags(0)
	flag.Parse()

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		log.Fatal(err)
	}

	err = openBrowser(fmt.Sprintf("http://localhost:%v", *port))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server started: http://localhost:%v \n", *port)
	(&http.Server{Handler: go101}).Serve(l)
}
