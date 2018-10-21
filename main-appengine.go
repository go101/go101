// +build appengine

// todo:
// * https://cloud.google.com/appengine/docs/standard/go111/go-differences
// * https://cloud.google.com/appengine/docs/standard/go111/specifying-dependencies

package main

import (
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", go101.ServeHTTP)
	appengine.Main()
}
