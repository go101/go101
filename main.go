package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var portFlag = flag.String("port", "55555", "server port")

func main() {
	log.SetFlags(0)
	flag.Parse()

	port := *portFlag
	if prt := os.Getenv("PORT"); prt != "" { // appengine std
		port = prt
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}

	err = openBrowser(fmt.Sprintf("http://localhost:%v", port))
	if err != nil {
		log.Fatal(err)
	}

	go go101.Update()

	log.Println("Server started:")
	log.Printf("   http://localhost:%v (non-cached version)\n", port)
	log.Printf("   http://127.0.0.1:%v (cached version)\n", port)
	(&http.Server{
		Handler:      go101,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
	}).Serve(l)
}
