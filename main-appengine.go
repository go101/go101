// +build appengine

package main

import (
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", go101.ServeHTTP)
	appengine.Main()
}
