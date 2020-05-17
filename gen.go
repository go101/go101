package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	//"errors"
	"net/http"
)

const GeneratedFolderName = "generated"

func genStaticFiles(rootPath string) {
	log.SetFlags(log.Lshortfile)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Can't get current path:", err)
	}

	fullPath := func(relPath ...string) string {
		return filepath.Join(append([]string{wd}, relPath...)...)
	}

	_, err = os.Stat(fullPath("go101.go"))
	if err != nil && os.IsNotExist(err) { //errors.Is(err, os.ErrNotExist) {
		log.Fatal("File go101.org not found. Not run in go101 folder?")
	}

	loadFile := func(uri string) []byte {
		fullPath := rootPath + uri

		res, err := http.Get(fullPath)
		if err != nil {
			log.Fatalf("Load file %s error: %s", uri, err)
		}

		content, err := ioutil.ReadAll(res.Body)

		log.Println(len(content), fullPath)

		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		return content
	}

	readFile := func(path string) []byte {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("Read file %s error: %s", path, err)
		}
		return data
	}

	readFolder := func(path string) (filenames, subfolders []string) {
		fds, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatalf("Read folder %s error: %s", path, err)
		}
		subfolders, filenames = make([]string, 0, len(fds)), make([]string, 0, len(fds))
		for _, fd := range fds {
			// ignore links, ...
			if fd.IsDir() {
				subfolders = append(subfolders, filepath.Join(path, fd.Name()))
				continue
			}
			filenames = append(filenames, filepath.Join(path, fd.Name()))
		}
		return
	}

	var readFolderRecursively func(path string) (filenames []string)
	readFolderRecursively = func(path string) (filenames []string) {
		type Folder struct {
			Path string
			Next *Folder
		}

		var head, tail *Folder
		regSubfolders := func(folders []string) {
			if len(folders) == 0 {
				return
			}
			start := 0
			if head == nil {
				head = &Folder{
					Path: folders[0],
				}
				tail = head
				start = 1
			}
			for i := start; i < len(folders); i++ {
				fldr := &Folder{
					Path: folders[i],
				}
				tail.Next = fldr
				tail = fldr
			}
		}

		var subfolders []string
		filenames, subfolders = readFolder(path)
		regSubfolders(subfolders)

		for head != nil {
			files, subfolders := readFolder(head.Path)
			head = head.Next
			regSubfolders(subfolders)
			filenames = append(filenames, files...)
		}
		return
	}

	// collect ...

	files := make(map[string][]byte, 128)

	files["index.html"] = readFile(fullPath("web", "index.html"))

	{
		dir := fullPath("articles", "res")
		filenames, _ := readFolder(dir)
		for _, f := range filenames {
			if strings.HasSuffix(f, ".png") || strings.HasSuffix(f, ".jpg") || strings.HasSuffix(f, ".jpeg") {
				name, err := filepath.Rel(dir, f)
				if err != nil {
					log.Fatalf("filepath.Rel(%s, %s) error: %s", dir, f, err)
				}
				files["article/res/"+name] = readFile(f)
			}
		}
	}

	{
		dir := fullPath("articles")
		filenames, _ := readFolder(dir)
		for _, f := range filenames {
			if strings.HasSuffix(f, ".html") {
				name, err := filepath.Rel(dir, f)
				if err != nil {
					log.Fatalf("filepath.Rel(%s, %s) error: %s", dir, f, err)
				}
				files["article/"+name] = loadFile("article/" + name)
			}
		}
	}

	{
		dir := fullPath("web", "static")
		filenames := readFolderRecursively(dir)
		for _, f := range filenames {
			name, err := filepath.Rel(dir, f)
			if err != nil {
				log.Fatalf("filepath.Rel(%s, %s) error: %s", dir, f, err)
			}
			files["static/"+name] = readFile(f)
		}
	}

	// write ...

	err = os.RemoveAll(fullPath(GeneratedFolderName))
	if err != nil {
		log.Fatalf("Remove folder %s error: %s", GeneratedFolderName, err)
	}

	for name, data := range files {
		fullFilename := fullPath(GeneratedFolderName, name)

		fullFilename = strings.Replace(fullFilename, "/", string(filepath.Separator), -1)
		fullFilename = strings.Replace(fullFilename, "\\", string(filepath.Separator), -1)

		if err := os.MkdirAll(filepath.Dir(fullFilename), 0700); err != nil {
			log.Fatalln("Mkdir error:", err)
		}

		if err := ioutil.WriteFile(fullFilename, data, 0644); err != nil {
			log.Fatalln("Write file error:", err)
		}

		log.Printf("Generated %s (size: %d).", name, len(data))
	}
}
