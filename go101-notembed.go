// +build !embed

package main

import (
	"go/build"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var staticFilesHandler = http.FileServer(http.Dir(filepath.Join(rootPath, "web", "static")))
var resFilesHandler = http.FileServer(http.Dir(filepath.Join(rootPath, "articles", "res")))

func loadArticleFile(file string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(rootPath, "articles", file))
}

func parseTemplate(commonPath string, files ...string) *template.Template {
	ts := make([]string, len(files))
	for i, f := range files {
		ts[i] = filepath.Join(rootPath, commonPath, f)
	}
	return template.Must(template.ParseFiles(ts...))
}

func updateGo101() {
	pullGo101Project(rootPath)
}

//=================================

var rootPath = findGo101ProjectRoot()

func findGo101ProjectRoot() string {
	if _, err := os.Stat(filepath.Join(".", "go101.go")); err == nil {
		return "."
	}

	for _, name := range []string{
		"gitlab.com/go101/go101",
		"gitlab.com/Go101/go101",
		"github.com/go101/go101",
		"github.com/Go101/go101",
	} {
		pkg, err := build.Import(name, "", build.FindOnly)
		if err == nil {
			return pkg.Dir
		}
	}

	return "."
}
