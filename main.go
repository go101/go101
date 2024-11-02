package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var portFlag = flag.String("port", "55555", "server port")
var genFlag = flag.Bool("gen", false, "HTML generation mode?")
var themeFlag = flag.String("theme", "", "theme (dark | light)")
var nobFlag = flag.Bool("nob", false, "not open browswer?")

var listenConfig net.ListenConfig

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
	//l, err := net.ListenTCP("tcp", addr)
	l, err := listenConfig.Listen(context.Background(), "tcp", addr.String())
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

	httpServer := &http.Server{
		Handler:      go101,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	runServer := func() {
		log.Println("Server started:")
		log.Printf("   http://localhost:%v (non-cached version)\n", addr.Port)
		log.Printf("   http://127.0.0.1:%v (cached version)\n", addr.Port)
		httpServer.Serve(l)
	}

	shutdownServer := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatalf("Server shutdown error: %s", err)
		}
		log.Println("Server shutdown.")
	}

	if genMode {
		go runServer()
		genStaticFiles(rootURL)
		shutdownServer()
		return
	}

	go runServer()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	shutdownServer()
}
