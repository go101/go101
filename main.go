package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

var portFlag = flag.String("port", "55555", "server port")
var genFlag = flag.Bool("gen", false, "HTML generation mode?")
var themeFlag = flag.String("theme", "", "theme (dark | light)")
var nobFlag = flag.Bool("nob", false, "not open browswer?")

func main() {
	log.SetFlags(0)
	flag.Parse()

	port, isAppEngine := *portFlag, false
	if prt := os.Getenv("PORT"); prt != "" { // appengine std
		port = prt
		isAppEngine = true
	}
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}

Retry:
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		if strings.Index(err.Error(), "bind: address already in use") >= 0 {
			addr.Port++
			if addr.Port < 65535 {
				goto Retry
			}
		}
		log.Fatal(err)
	}

	go101.theme = *themeFlag

	genMode, rootURL := *genFlag, fmt.Sprintf("http://localhost:%v/", addr.Port)
	if !genMode && !isAppEngine {
		if !*nobFlag {
			err = openBrowser(rootURL)
			if err != nil {
				log.Println(err)
			}
		}

		go updateGo101()
	}

	runServer := func() {
		log.Println("Server started:")
		log.Printf("   http://localhost:%v (non-cached version)\n", addr.Port)
		log.Printf("   http://127.0.0.1:%v (cached version)\n", addr.Port)
		(&http.Server{
			Handler:      go101,
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  5 * time.Second,
		}).Serve(l)
	}

	if genMode {
		go runServer()
		genStaticFiles(rootURL)
		return
	}

	runServer()
}
