// +build go1.16

package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

//go:embed web articles/*.html
//go:embed articles/res/*.png
//go:embed articles/res/*.jpg
var allFiles embed.FS

var staticFilesHandler, resFilesHandler = func() (http.Handler, http.Handler) {
	staticFiles, err := fs.Sub(allFiles, path.Join("web", "static"))
	if err != nil {
		panic(fmt.Sprintf("construct static file system error: %s", err))
	}
	resFiles, err := fs.Sub(allFiles, path.Join("articles", "res"))
	if err != nil {
		panic(fmt.Sprintf("construct res file system error: %s", err))
	}

	//paths1 := printFS("static files", staticFiles)
	//paths2 := printFS("res files", resFiles)
	//for _, path := range paths1 {
	//	openFileInFS(staticFiles, path)
	//}
	//for _, path := range paths2 {
	//	openFileInFS(resFiles, path)
	//}

	return http.FileServer(http.FS(staticFiles)), http.FileServer(http.FS(resFiles))
}()

func loadArticleFile(file string) ([]byte, error) {
	content, err := allFiles.ReadFile(path.Join("articles", file))
	if err != nil {
		return nil, err
	}
	return content, nil
}

func parseTemplate(commonPath string, files ...string) *template.Template {
	ts := make([]string, len(files))
	for i, f := range files {
		ts[i] = path.Join(commonPath, f)
	}
	return template.Must(template.ParseFS(allFiles, ts...))
}

func updateGo101() {
	if _, err := os.Stat(filepath.Join(".", "go101.go")); err == nil {
		pullGo101Project("")
		return
	}
	if filepath.Base(os.Args[0]) == "go101" {
		log.Println("go", "install", "go101.org/go101@latest")
		output, err := runShellCommand(time.Minute/2, "", "go", "install", "go101.org/go101@latest")
		if err != nil {
			log.Printf("error: %s\n%s", err, output)
		} else {
			log.Printf("done.")
		}
	}

	// no ideas how to update
}

//==============================================

//func printFS(title string, s fs.FS) []string {
//	println("==================", title)
//	files := make([]string, 0, 256)
//	fs.WalkDir(s, ".", func(path string, d fs.DirEntry, err error) error {
//		println(path)
//		files = append(files, path)
//		return nil
//	})
//	return files
//}

//func openFileInFS(s fs.FS, name string) {
//	f, err := s.Open(name)
//	if err != nil {
//		println("open", name, "error:", err.Error())
//		return
//	}
//	f.Close()
//	println("open", name, "ok")
//}
