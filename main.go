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
var genFlag = flag.Bool("gen", false, "HTML generation mode?")
var themeFlag = flag.String("theme", "dark", "theme (dark | light), dark defaultly")

func main() {
	log.SetFlags(0)
	flag.Parse()

	port, isAppEngine := *portFlag, false
	if prt := os.Getenv("PORT"); prt != "" { // appengine std
		port = prt
		isAppEngine = true
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}

	go101.theme = *themeFlag

	genMode, rootURL := *genFlag, fmt.Sprintf("http://localhost:%v/", port)
	if !genMode && !isAppEngine {
		err = openBrowser(fmt.Sprintf("http://localhost:%v", port))
		if err != nil {
			log.Println(err)
		}

		go updateGo101()
	}

	runServer := func() {
		log.Println("Server started:")
		log.Printf("   http://localhost:%v (non-cached version)\n", port)
		log.Printf("   http://127.0.0.1:%v (cached version)\n", port)
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
