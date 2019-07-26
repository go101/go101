package main

import (
	"bytes"
	"context"
	"errors"
	"go/build"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Go101 struct {
	staticHandler     http.Handler
	articleResHandler http.Handler
	isLocalServer     bool
	articlePages      map[string][]byte
	serverMutex       sync.Mutex
}

var (
	rootPath = findGo101ProjectRoot()
	go101    = &Go101{
		staticHandler:     http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(rootPath, "static")))),
		articleResHandler: http.StripPrefix("/article/res/", http.FileServer(http.Dir(filepath.Join(rootPath, "articles", "res")))),
		isLocalServer:     false, // may be modified later
		articlePages:      map[string][]byte{},
	}
)

func (go101 *Go101) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	group, item := "", ""
	tokens := strings.SplitN(r.URL.Path, "/", 3)
	if len(tokens) > 1 {
		group = tokens[1]
		if len(tokens) > 2 {
			item = tokens[2]
		}
	}

	switch go101.ConfirmLocalServer(isLocalRequest(r)); strings.ToLower(group) {
	case "static":
		w.Header().Set("Cache-Control", "max-age=360000") // 10 hours
		go101.staticHandler.ServeHTTP(w, r)
	case "article":
		item = strings.ToLower(item)
		if strings.HasPrefix(item, "res/") {
			w.Header().Set("Cache-Control", "max-age=360000") // 10 hours
			go101.articleResHandler.ServeHTTP(w, r)
			return
		} else if go101.IsLocalServer() && (strings.HasPrefix(item, "print-") || strings.HasPrefix(item, "pdf-")) {
			w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
			idx := strings.IndexByte(item, '-')
			go101.RenderPrintPage(w, r, item[:idx], item[idx+1:])
			return
		}
		go101.RenderArticlePage(w, r, item)
	case "":
		http.Redirect(w, r, "/article/101.html", http.StatusTemporaryRedirect)
	default:
		http.Redirect(w, r, "/article/101.html", http.StatusNotFound)
	}

}

func (go101 *Go101) ConfirmLocalServer(isLocal bool) {
	go101.serverMutex.Lock()
	if go101.isLocalServer != isLocal {
		go101.isLocalServer = isLocal
		if go101.isLocalServer {
			unloadPageTemplates()                    // loaded in one init function
			go101.articlePages = map[string][]byte{} // invalidate article caches
		}
	}
	go101.serverMutex.Unlock()
}

func (go101 *Go101) IsLocalServer() (isLocal bool) {
	go101.serverMutex.Lock()
	isLocal = go101.isLocalServer
	go101.serverMutex.Unlock()
	return
}

func (go101 *Go101) ArticlePage(file string) ([]byte, bool) {
	go101.serverMutex.Lock()
	page := go101.articlePages[file]
	isLocal := go101.isLocalServer
	go101.serverMutex.Unlock()
	return page, isLocal
}

func (go101 *Go101) CacheArticlePage(file string, page []byte) {
	go101.serverMutex.Lock()
	go101.articlePages[file] = page
	go101.serverMutex.Unlock()
}

//===================================================
// pages
//==================================================

type Article struct {
	Content, Title     template.HTML
	TitleWithoutTags   string
	FilenameWithoutExt string
}

var articlePagesMutex sync.Mutex
var articlePages = map[string][]byte{}
var schemes = map[bool]string{false: "http://", true: "https://"}

func (go101 *Go101) RenderArticlePage(w http.ResponseWriter, r *http.Request, file string) {
	page, isLocal := go101.ArticlePage(file)
	if page == nil {
		article, err := retrieveArticleContent(file)
		if err == nil {
			//var pageURL string
			//if !isLocal {
			//	//pageURL = r.URL.String() // looks only working for GAE
			//	pageURL = schemes[r.TLS != nil] + r.Host + r.RequestURI
			//}
			pageParams := map[string]interface{}{
				"Article":       article,
				"Title":         article.TitleWithoutTags,
				"IsLocalServer": isLocal,
				//"SocialLinkURL": pageURL, // non-blank to show social buttons
			}
			t := retrievePageTemplate(Template_Article, !isLocal)
			var buf bytes.Buffer
			if err = t.Execute(&buf, pageParams); err == nil {
				page = buf.Bytes()
			} else {
				page = []byte(err.Error())
			}
		} else if os.IsNotExist(err) {
			page = []byte{} // blank page means page not found.
		}

		if !isLocal {
			go101.CacheArticlePage(file, page)
		}
	}

	// ...
	if len(page) == 0 { // blank page means page not found.
		log.Printf("article page %s is not found", file)
		w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
		http.Redirect(w, r, "/article/101.html", http.StatusNotFound)
		return
	}

	if isLocal {
		w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	} else {
		w.Header().Set("Cache-Control", "max-age=5000") // about 1.5 hours
	}
	w.Write(page)
}

const H1, _H1, MaxLen = "<h1>", "</h1>", 128

var TagSigns = [2]rune{'<', '>'}

func retrieveArticleContent(file string) (Article, error) {
	article := Article{}
	content, err := ioutil.ReadFile(filepath.Join(rootPath, "articles", file))
	if err != nil {
		return article, err
	}

	article.Content = template.HTML(content)
	article.FilenameWithoutExt = strings.TrimSuffix(file, ".html")

	// retrieve titles
	j, i := -1, strings.Index(string(article.Content), H1)
	if i >= 0 {
		i += len(H1)
		j = strings.Index(string(article.Content[i:i+MaxLen]), _H1)
		if j >= 0 {
			article.Title = article.Content[i-len(H1) : i+j+len(_H1)]
			article.Content = article.Content[i+j+len(_H1):]
			k, s := 0, make([]rune, 0, MaxLen)
			for _, r := range article.Title {
				if r == TagSigns[k] {
					k = (k + 1) & 1
				} else if k == 0 {
					s = append(s, r)
				}
			}
			article.TitleWithoutTags = string(s)
		}
	}
	if j < 0 {
		log.Println("retrieveTitlesForArticle", article.FilenameWithoutExt, "failed")
	}

	return article, nil
}

const Anchor, _Anchor, LineToRemoveTag, endl = `<li><a class="index" href="`, `">`, `(to remove)`, "\n"
const IndexContentStart, IndexContentEnd = `<!-- index starts (don't remove) -->`, `<!-- index ends (don't remove) -->`

func (go101 *Go101) RenderPrintPage(w http.ResponseWriter, r *http.Request, printTarget, item string) {
	page, isLocal := go101.ArticlePage(item)
	if page == nil {
		var err error
		var pageParams map[string]interface{}
		switch item {
		case "book101":
			pageParams, err = buildBook101PrintParams()
		}

		if err == nil {
			if pageParams == nil {
				pageParams = map[string]interface{}{}
			}
			//q := r.URL.Query().Get("showcovers")
			//pageParams["ShowCovers"] = q == "1" || q == "true"
			//pageParams["IndexTitle"] = r.URL.Query().Get("indextitle")
			pageParams["PrintTarget"] = printTarget
			pageParams["IsLocalServer"] = isLocal

			t := retrievePageTemplate(Template_PrintBook, !isLocal)
			var buf bytes.Buffer
			if err = t.Execute(&buf, pageParams); err == nil {
				page = buf.Bytes()
			}
		}

		if err != nil {
			page = []byte(err.Error())
		}

		if !isLocal {
			go101.CacheArticlePage(item, page)
		}
	}

	// ...
	if len(page) == 0 { // blank page means page not found.
		log.Printf("print page %s is not found", item)
		w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
		http.Redirect(w, r, "/article/101.html", http.StatusNotFound)
		return
	}

	if isLocal {
		w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	} else {
		w.Header().Set("Cache-Control", "max-age=5000") // about 1.5 hours
	}
	w.Write(page)
}

func buildBook101PrintParams() (map[string]interface{}, error) {
	article, err := retrieveArticleContent("101.html")
	if err != nil {
		return nil, err
	}

	// get all index article content by removing some lines
	var builder strings.Builder

	content := string(article.Content)
	i := strings.Index(content, IndexContentStart)
	if i < 0 {
		err = errors.New(IndexContentStart + " not found")
		return nil, err
	}
	i += len(IndexContentStart)
	content = content[i:]
	i = strings.Index(content, IndexContentEnd)
	if i >= 0 {
		content = content[:i]
	}

	for range [1000]struct{}{} {
		i = strings.Index(content, LineToRemoveTag)
		if i < 0 {
			break
		}
		start := strings.LastIndex(content[:i], endl)
		if start >= 0 {
			builder.WriteString(content[:start])
		}
		end := strings.Index(content[i:], endl)
		content = content[i:]
		if end < 0 {
			end = len(content)
		}
		content = content[end:]
	}
	builder.WriteString(content)

	// the index article
	articles := make([]Article, 0, 100)
	article.FilenameWithoutExt = "101"
	article.Content = template.HTML(builder.String())
	articles = append(articles, article)

	// find all articles from links
	content = string(article.Content)
	for range [1000]struct{}{} {
		i = strings.Index(content, Anchor)
		if i < 0 {
			break
		}
		content = content[i+len(Anchor):]
		i = strings.Index(content, _Anchor)
		if i < 0 {
			break
		}

		article, err := retrieveArticleContent(content[:i])
		if err != nil {
			log.Printf("retrieve article %s error: %s", content[:i], err)
		} else {
			articles = append(articles, article)
		}

		content = content[i+len(_Anchor):]
	}

	return map[string]interface{}{"Articles": articles}, nil
}

//===================================================
// templates
//==================================================

type PageTemplate uint

const (
	Template_Article PageTemplate = iota
	Template_PrintBook
	NumPageTemplates
)

var pageTemplates [NumPageTemplates + 1]*template.Template
var pageTemplatesMutex sync.Mutex //

func init() {
	for i := range pageTemplates {
		retrievePageTemplate(PageTemplate(i), true)
	}
}

func retrievePageTemplate(which PageTemplate, cacheIt bool) *template.Template {
	if which > NumPageTemplates {
		which = NumPageTemplates
	}

	pageTemplatesMutex.Lock()
	t := pageTemplates[which]
	pageTemplatesMutex.Unlock()

	if t == nil {
		switch which {
		case Template_Article:
			t = parseTemplate(filepath.Join(rootPath, "templates"), "base", "article")
		case Template_PrintBook:
			t = parseTemplate(filepath.Join(rootPath, "templates"), "pdf")
		default:
			t = template.New("blank")
		}

		if cacheIt {
			pageTemplatesMutex.Lock()
			pageTemplates[which] = t
			pageTemplatesMutex.Unlock()
		}
	}
	return t
}

func unloadPageTemplates() {
	pageTemplatesMutex.Lock()
	for i := range pageTemplates {
		pageTemplates[i] = nil
	}
	pageTemplatesMutex.Unlock()
}

//===================================================
// git
//===================================================

func gitPull() {
	output, err := runShellCommand(time.Minute/2, "git", "pull")
	if err != nil {
		log.Println("git pull:", err)
	} else {
		log.Printf("git pull: %s", output)
	}
}

func goGet(pkgPath string) {
	_, err := runShellCommand(time.Minute/2, "go", "get", "-u", pkgPath)
	if err != nil {
		log.Println("go get -u "+pkgPath+":", err)
	} else {
		log.Println("go get -u " + pkgPath + " succeeded.")
	}
}

func (go101 *Go101) Update() {
	<-time.After(time.Minute / 2)

	//output, err := runShellCommand(time.Minute/2, "git", "remote")
	//if err != nil {
	//	log.Println("list git remotes error:", err)
	//	return
	//}
	//k := bytes.IndexRune(output, '\n')
	//if k < 0 {
	//	log.Println("find git remote failed:", output)
	//	return
	//}

	//configItem := "remote." + string(bytes.TrimSpace(output[:k])) + ".url"
	//output, err = runShellCommand(time.Minute/2, "git", "config", "--get", configItem)
	//if err != nil {
	//	log.Println("get "+configItem+" error:", err)
	//	return
	//}
	//a, b := bytes.Index(output, []byte("://")), bytes.Index(output, []byte(".git"))
	//if a += 3; a < 3 {
	//	a = 0
	//}
	//if b < 0 {
	//	b = len(output)
	//}

	//pkgPath := string(bytes.TrimSpace(output[a:b]))
	gitPull() // goGet(pkgPath)
	for {
		<-time.After(time.Hour * 24)
		gitPull() // goGet(pkgPath)
	}
}

//===================================================
// utils
//===================================================

func parseTemplate(path string, files ...string) *template.Template {
	ts := make([]string, len(files))
	for i, f := range files {
		ts[i] = filepath.Join(path, f)
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
	if _, err := os.Stat(filepath.Join(".", "go101.go")); err == nil {
		return "."
	}

	for _, name := range []string{
		"gitlab.com/go101/go101", "gitlab.com/Go101/go101",
		"github.com/go101/go101", "github.com/Go101/go101",
	} {
		pkg, err := build.Import(name, "", build.FindOnly)
		if err == nil {
			return pkg.Dir
		}
	}

	return "."
}

func isLocalRequest(r *http.Request) bool {
	end := strings.Index(r.Host, ":")
	if end < 0 {
		end = len(r.Host)
	}
	hostname := r.Host[:end]
	return hostname == "localhost" // || hostname == "127.0.0.1" // 127.* for local cached version now
}

func runShellCommand(timeout time.Duration, cmd string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	command := exec.CommandContext(ctx, cmd, args...)
	command.Dir = rootPath
	return command.Output()
}
