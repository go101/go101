package main

import (
	"flag"
	"fmt"
	"go/build"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var port = flag.Int("port", 55555, "server port")

func main() {
	log.SetFlags(0)
	flag.Parse()

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		log.Fatal(err)
	}

	err = openBrowser(fmt.Sprintf("http://localhost:%v", *port))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server started: http://localhost:%v \n", *port)
	(&http.Server{Handler: go101}).Serve(l)
}

var (
	rootPath              = findGo101ProjectRoot()
	go101    http.Handler = &Go101{
		staticHandler:     http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.ToSlash(rootPath + "static")))),
		articleResHandler: http.StripPrefix("/article/res/", http.FileServer(http.Dir(filepath.ToSlash(rootPath + "articles/res")))),
	}
)

type Go101 struct {
	staticHandler     http.Handler
	articleResHandler http.Handler
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
		item = strings.ToLower(item)
		if strings.HasPrefix(item, "res/") {
			go101.articleResHandler.ServeHTTP(w, r)
			return
		}
		
		if go101.renderArticlePage(w, r, item) {
			return
		}
	}

	http.Redirect(w, r, "/article/101.html", http.StatusTemporaryRedirect)
}

//===================================================
// pages
//==================================================

type Article struct {
	Content, Title   template.HTML
	TitleWithoutTags string
	FileWithoutExt   string
}

var articleTemplate = parseTemplate("base", "article")

var articleContents = func() map[string]Article {
	path := filepath.ToSlash(rootPath + "articles/")
	if files, err := filepath.Glob(path + "*.html"); err != nil {
		log.Fatal(err)
		return nil
	} else {
		contents := make(map[string]Article, len(files))
		for _, f := range files {
			contents[strings.TrimPrefix(f, path)] = Article{}
		}
		return contents
	}
}()

func retrieveArticleContent(file string, cachedIt bool) (Article, error) {
	article, present := articleContents[file]
	if !present {
		return Article{}, nil
	}
	if article.Content == "" {
		content, err := ioutil.ReadFile(filepath.ToSlash(rootPath + "articles/" + file))
		if err != nil {
			return Article{}, err
		}
		article.Content = template.HTML(content)
		article.FileWithoutExt = strings.TrimSuffix(file, ".html")
		retrieveTitlesForArticle(&article)
		if cachedIt {
			articleContents[file] = article
		}
	}
	return article, nil
}

func retrieveTitlesForArticle(article *Article) {
	const H1, _H1, MaxLen = "<h1>", "</h1>", 128
	i := strings.Index(string(article.Content), H1)
	if i >= 0 {
		i += len(H1)
		j := strings.Index(string(article.Content[i:i+MaxLen]), _H1)
		if j >= 0 {
			article.Title = article.Content[i:i+j]
			k, tags, s := 0, [2]rune{'<', '>'}, make([]rune, 0, MaxLen)
			for _, r := range article.Title {
				if r == tags[k] {
					k = (k+1) & 1
				} else if k == 0 {
					s = append(s, r)
				}
			}
			article.TitleWithoutTags = string(s)
		}
	}
}

func (*Go101) renderArticlePage(w http.ResponseWriter, r *http.Request, file string) bool {
	article, err := retrieveArticleContent(file, !isLocalRequest(r))
	if err == nil {
		page := map[string]interface{}{
			"Article": article,
			"Title":   article.TitleWithoutTags,
		}
		if err = articleTemplate.Execute(w, page); err == nil {
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
		ts[i] = filepath.ToSlash(rootPath + "templates/" + f)
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

func findGo101ProjectRoot() string {
	if _, err := os.Stat(filepath.ToSlash("./go101.go")); err == nil {
		return ""
	}

	pkg, err := build.Import("github.com/go101/go101", "", build.FindOnly)
	if err != nil {
		log.Fatal("Can't find pacakge: github.com/go101/go101")
		return ""
	}
	return filepath.ToSlash(pkg.Dir + "/")
}

func isLocalRequest(r *http.Request) bool {
	end := strings.Index(r.Host, ":")
	if end < 0 {
		end = len(r.Host)
	}
	hostname := r.Host[:end]
	return hostname == "localhost" || hostname == "127.0.0.1"
}
