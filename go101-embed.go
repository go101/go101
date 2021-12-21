//go:build go1.16
// +build go1.16

package main

import (
	"embed"
	//"errors"
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

//go:embed web
//go:embed pages
var allFiles embed.FS

var staticFilesHandler = func() http.Handler {
	if wdIsGo101ProjectRoot {
		return staticFilesHandler_NonEmbedding
	}

	staticFiles, err := fs.Sub(allFiles, path.Join("web", "static"))
	if err != nil {
		panic(fmt.Sprintf("construct static file system error: %s", err))
	}

	return http.FileServer(http.FS(staticFiles))
}()

func collectPageGroups() map[string]*PageGroup {
	if wdIsGo101ProjectRoot {
		return collectPageGroups_NonEmbedding()
	}

	entries, err := fs.ReadDir(allFiles, "pages")
	if err != nil {
		panic("collect page groups (embedding) error: " + err.Error())
	}

	pageGroups := make(map[string]*PageGroup, len(entries))

	for _, e := range entries {
		if e.IsDir() {
			group, handler := e.Name(), dummyHandler
			resFiles, err := fs.Sub(allFiles, path.Join("pages", e.Name(), "res"))
			if err == nil {
				var urlGroup string
				// For history reason, fundamentals pages uses "/article/xxx" URLs.
				if group == "fundamentals" {
					urlGroup = "/article"
				} else if group != "website" {
					urlGroup = "/" + group
				}
				handler = http.StripPrefix(urlGroup+"/res/", http.FileServer(http.FS(resFiles)))
			} else if !os.IsNotExist(err) { // !errors.Is(err, os.ErrNotExist) {
				log.Println(err)
			}

			pageGroups[group] = &PageGroup{resHandler: handler}
		}
	}

	return pageGroups
}

func loadArticleFile(group, file string) ([]byte, error) {
	if wdIsGo101ProjectRoot {
		return loadArticleFile_NonEmbedding(group, file)
	}

	content, err := allFiles.ReadFile(path.Join("pages", group, file))
	if err != nil {
		return nil, err
	}
	return content, nil
}

func parseTemplate(commonPaths []string, files ...string) *template.Template {
	if wdIsGo101ProjectRoot {
		return parseTemplate_NonEmbedding(commonPaths, files...)
	}

	cp := path.Join(commonPaths...)
	ts := make([]string, len(files))
	for i, f := range files {
		ts[i] = path.Join(cp, f)
	}
	return template.Must(template.ParseFS(allFiles, ts...))
}

func updateGo101() {
	if wdIsGo101ProjectRoot {
		updateGo101_NonEmbedding()
		return
	}

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
