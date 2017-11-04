package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	log.SetFlags(0)
	
	flag.Parse()
	port := *flag.Int("port", 55555, "server port")

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}

	err = openBrowser(fmt.Sprintf("http://localhost:%v", port))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server started.")
	(&http.Server{Handler: go101}).Serve(l)
}

var go101 http.Handler = &Go101{
	staticHandler: http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))),
}
type Go101 struct{
	staticHandler http.Handler
}
func (go101 *Go101) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	group, item := "", ""
	tokens := strings.SplitN(r.URL.Path, "/", 3)
	if len(tokens) > 1 {
		group = tokens[1]
		if len(tokens) > 2 {
			item = tokens[2]
		}
	}
	
	// log.Println("group=", group, ", item=", item)
	
	switch strings.ToLower(group) {
	default:
		http.Error(w, "", http.StatusNotFound)
		return
	case "":
	
	case "static":
		go101.staticHandler.ServeHTTP(w, r)
		return
	case "article":
		if go101.renderArticlePage(w, r, strings.ToLower(item)) {
			return
		}
		log.Println("****************88")
	}
	
	http.Redirect(w, r, "/article/101.html", http.StatusTemporaryRedirect)
}

//===================================================
// pages
//==================================================

var articleTemplate = parseTemplate("base", "article")
var articleContents = func() map[string]template.HTML {
	if files, err := filepath.Glob("articles/*.html"); err != nil {
		log.Fatal(err)
		return nil
	} else {
		contents := make(map[string]template.HTML, len(files))
		for _, f := range files {
			contents[f] = ""
		}
		return contents
	}
}()

func retrieveArticleContent(article string, cachedIt bool) (template.HTML, error) {
	article = "articles/" + article
	html, present := articleContents[article]
	if !present {
		return "", nil
	}
	if html == "" {
		content, err := ioutil.ReadFile(article)
		if err != nil {
			return "", err
		}
		html = template.HTML(content)
		if cachedIt {
			articleContents[article] = html
		}
	}
	return html, nil
}

func (*Go101) renderArticlePage(w http.ResponseWriter, r *http.Request, file string) bool {
	content, err := retrieveArticleContent(file, !isLocalRequest(r))
	if err == nil {
		article := map[string]interface{} {
			"Content": content,
		}
		page := map[string]interface{} {
			"Article": article,
		}
		if err := articleTemplate.Execute(w, page); err == nil {
			return true
		}
	}
	w.Write([]byte(err.Error()))
	return false
}

//===================================================
// utils
//===================================================

func parseTemplate(files ...string) *template.Template {
	ts := make([]string, len(files))
	for i, f := range files {
		ts[i] = "templates/" + f
	}
	return template.Must(template.ParseFiles(ts...))
}

// https://stackoverflow.com/questions/39320371/how-start-web-server-to-open-page-in-browser-in-golang
func openBrowser(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	return exec.Command(cmd, append(args, url)...).Start()
}

func isLocalRequest(r *http.Request) bool {
	end := strings.Index(r.Host, ":")
	if end < 0 {
		end = len(r.Host)
	}
	hostname := r.Host[:end]
	return hostname == "localhost" || hostname == "127.0.0.1"
}
