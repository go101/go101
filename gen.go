package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const GeneratedFolderName = "generated"

func genStaticFiles(rootURL string) {
	log.SetFlags(log.Lshortfile)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Can't get current path:", err)
	}

	fullPath := func(relPath ...string) string {
		return filepath.Join(append([]string{wd}, relPath...)...)
	}

	_, err = os.Stat(fullPath("web", "static", "go101"))
	if err != nil {
		if os.IsNotExist(err) { //errors.Is(err, os.ErrNotExist) {
			log.Fatal("File web/static/go101 not found. Not run in go101 folder?")
		}
		log.Fatal(err)
	}

	// load from http server
	loadFile := func(uri string) []byte {
		fullURL := rootURL + uri

		res, err := http.Get(fullURL)
		if err != nil {
			log.Fatalf("Load file %s error: %s", uri, err)
		}

		content, err := ioutil.ReadAll(res.Body)

		log.Println(len(content), fullURL)

		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		return content
	}

	// read from OS file system
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

	// md -> html
	md2htmls := func(group string) {
		dir := fullPath("pages", group)
		outputs, err := runShellCommand(time.Minute/2, dir, "ebooktool", "-md2htmls")
		if err != nil {
			log.Fatalf("ebooktool failed to execute in directory: %s.\n%s", dir, outputs)
		}
	}

	// tmd -> html
	tmd2htmls := func(group string) {
		dir := fullPath("pages", group)
		filenames, _ := readFolder(dir)
		for _, filename := range filenames {
			if strings.HasSuffix(filename, ".tmd") {
				outputs, err := runShellCommand(time.Minute/2, dir, "tmd", "gen", "--enabled-custom-apps=html", filename)
				if err != nil {
					log.Fatalf("tmd failed to execute in directory: %s.\n%s", dir, outputs)
				}
			}
		}
	}

	// collect ...

	files := make(map[string][]byte, 128)

	files["index.html"] = loadFile("")

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

	collectPageGroupFiles := func(group, urlPrefix string, collectRes bool) {
		if collectRes {
			dir := fullPath("pages", group, "res")
			filenames, _ := readFolder(dir)
			for _, f := range filenames {
				if strings.HasSuffix(f, ".png") || strings.HasSuffix(f, ".jpg") {
					name, err := filepath.Rel(dir, f)
					if err != nil {
						log.Fatalf("filepath.Rel(%s, %s) error: %s", dir, f, err)
					}
					files[urlPrefix+"res/"+name] = readFile(f)
				}
			}
		}

		md2htmls(group)
		tmd2htmls(group)

		{
			dir := fullPath("pages", group)
			filenames, _ := readFolder(dir)
			for _, f := range filenames {
				if strings.HasSuffix(f, ".html") {
					name, err := filepath.Rel(dir, f)
					if err != nil {
						log.Fatalf("filepath.Rel(%s, %s) error: %s", dir, f, err)
					}
					files[urlPrefix+name] = loadFile(urlPrefix + name)
				}
			}
		}
	}

	{
		infos, err := ioutil.ReadDir(fullPath("pages"))
		if err != nil {
			panic("collect page groups error: " + err.Error())
		}

		for _, e := range infos {
			if e.IsDir() {
				group := e.Name()

				var urlPrefix string
				if group == "fundamentals" {
					// For history reason, fundamentals pages use "/article/xxx" URLs.
					urlPrefix = "article/"
				} else if group != "website" {
					urlPrefix = group + "/"
				}

				var collectRes bool
				if _, err := os.Stat(fullPath("pages", group, "res")); err == nil {
					collectRes = true
				}

				collectPageGroupFiles(group, urlPrefix, collectRes)
			}
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
