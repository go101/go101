package main

import (
	"bytes"
	"context"
	//"errors"
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
	staticHandler http.Handler
	isLocalServer bool
	pageGroups    map[string]*PageGroup
	articlePages  Cache
	gogetPages    Cache
	serverMutex   sync.Mutex
	theme         string // default is "dark"
}

type PageGroup struct {
	resHandler   http.Handler
	indexContent template.HTML
}

var go101 = &Go101{
	staticHandler: http.StripPrefix("/static/", staticFilesHandler),
	isLocalServer: false, // may be modified later
	pageGroups:    collectPageGroups(),
}

func init() {
	for group, pg := range go101.pageGroups {
		pg.indexContent = retrieveIndexContent(group)
	}
}

func (go101 *Go101) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var group, item string
	if tokens := strings.SplitN(r.URL.Path[1:], "/", 2); len(tokens) == 2 {
		group, item = tokens[0], tokens[1]
	} else { // len(tokens) == 1
		item = tokens[0]
	}

	switch go101.ConfirmLocalServer(isLocalRequest(r)); group {
	default:
		go101.ServeGoGetPages(w, r, group, item)
	case "":
		go101.ServeGoGetPages(w, r, item, "")
	case "res":
		go101.serveGroupItem(w, r, "website", r.URL.Path[1:])
	case "static":
		w.Header().Set("Cache-Control", "max-age=31536000") // one year
		go101.staticHandler.ServeHTTP(w, r)
	case "article":
		// for history reason, fundamentals pages use "article/xxx" URLs
		go101.serveGroupItem(w, r, "fundamentals", item)
	case "optimizations", "details-and-tips", "quizzes", "generics",
		"apps-and-libs", "blog":
		go101.serveGroupItem(w, r, group, item)
	}
}

func (go101 *Go101) serveGroupItem(w http.ResponseWriter, r *http.Request, group, item string) {
	item = strings.ToLower(item)
	if strings.HasPrefix(item, "res/") {
		w.Header().Set("Cache-Control", "max-age=31536000") // one year
		go101.pageGroups[group].resHandler.ServeHTTP(w, r)
	} else if !go101.RedirectArticlePage(w, r, group, item) {
		go101.RenderArticlePage(w, r, group, item)
	}
}

func (go101 *Go101) ConfirmLocalServer(isLocal bool) {
	go101.serverMutex.Lock()
	defer go101.serverMutex.Unlock()
	if go101.isLocalServer != isLocal {
		go101.isLocalServer = isLocal
		if go101.isLocalServer {
			unloadPageTemplates()      // loaded in one init function
			go101.articlePages.Clear() // invalidate article caches
			go101.gogetPages.Clear()   // invalidate go-gets caches
		}
	}
}

func (go101 *Go101) IsLocalServer() (isLocal bool) {
	go101.serverMutex.Lock()
	defer go101.serverMutex.Unlock()
	isLocal = go101.isLocalServer
	return
}

func pullGo101Project(wd string) {
	<-time.After(time.Minute / 2)
	gitPull(wd)
	for {
		<-time.After(time.Hour * 24)
		gitPull(wd)
	}
}

//===================================================
// pages
//==================================================

type Article struct {
	Content, Title, Index template.HTML
	TitleWithoutTags      string
	Group, Filename       string
	FilenameWithoutExt    string
}

var schemes = map[bool]string{false: "http://", true: "https://"}

func (go101 *Go101) RenderArticlePage(w http.ResponseWriter, r *http.Request, group, file string) {
	page, isLocal := go101.articlePages.Get(group, file), go101.IsLocalServer()
	if page == nil {
		article, err := retrieveArticleContent(group, file)
		if err == nil {
			article.Index = disableArticleLink(go101.pageGroups[group].indexContent, file)
			pageParams := map[string]interface{}{
				"Article":       article,
				"Title":         article.TitleWithoutTags,
				"Theme":         go101.theme,
				"IsLocalServer": isLocal,
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
			go101.articlePages.Set(group, file, page)
		}
	}

	if len(page) == 0 { // blank page means page not found.
		log.Printf("article page %s/%s is not found", group, file)
		//w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
		http.Redirect(w, r, "/", http.StatusNotFound)
		return
	}

	if isLocal {
		w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	} else {
		w.Header().Set("Cache-Control", "max-age=50000") // about 14 hours
	}
	w.Write(page)
}

var H1, _H1 = []byte("<h1>"), []byte("</h1>")
var H2, _H2 = []byte("<h2>"), []byte("</h2>")

const MaxTitleLen = 256

var TagSigns = [2]rune{'<', '>'}

func retrieveArticleContent(group, file string) (Article, error) {
	article := Article{}
	content, err := loadArticleFile(group, file)
	if err != nil {
		return article, err
	}

	article.Content = template.HTML(content)
	article.Group = group
	article.Filename = file
	article.FilenameWithoutExt = strings.TrimSuffix(file, ".html")

	// retrieve titles
	splitTitleContent := func(startTag, endTag []byte) (int, int) {
		j, i := -1, bytes.Index(content, startTag)
		if i >= 0 {
			i += len(startTag)
			j = bytes.Index(bytesWithLength(content[i:], MaxTitleLen), endTag)
		}
		if j < 0 {
			return -1, 0
		}
		return i - len(startTag), i + j + len(endTag)
	}

	titleStart, contentStart := splitTitleContent(H1, _H1)
	if titleStart < 0 {
		titleStart, contentStart = splitTitleContent(H2, _H2)
	}
	if titleStart < 0 {
		//log.Println("retrieveTitlesForArticle failed:", group, file)
	} else {
		article.Title = article.Content[titleStart:contentStart]
		article.Content = article.Content[contentStart:]
		k, s := 0, make([]rune, 0, MaxTitleLen)
		for _, r := range article.Title {
			if r == TagSigns[k] {
				k = (k + 1) & 1
			} else if k == 0 {
				s = append(s, r)
			}
		}
		article.TitleWithoutTags = string(s)
	}

	return article, nil
}

func retrieveIndexContent(group string) template.HTML {
	page101, err := retrieveArticleContent(group, "101.html")
	if err != nil {
		if os.IsNotExist(err) { // errors.Is(err, os.ErrNotExist) {
			return ""
		}
		panic(err)
	}
	content := []byte(page101.Content)
	start := []byte("<!-- index starts (don't remove) -->")
	i := bytes.Index(content, start)
	if i < 0 {
		//panic("index not found")
		//log.Printf("index not found in %s/101/html", group)
		return ""
	}
	content = content[i+len(start):]
	end := []byte("<!-- index ends (don't remove) -->")
	i = bytes.Index(content, end)
	if i < 0 {
		//panic("index not found")
		//log.Printf("index not found in %s/101/html", group)
		return ""
	}
	content = content[:i]
	//comments := [][]byte{
	//	[]byte("<!-- (to remove) for printing"),
	//	[]byte("(to remove) -->"),
	//}
	//for _, cmt := range comments {
	//	i = bytes.Index(content, cmt)
	//	if i >= 0 {
	//		filleBytes(content[i:i+len(cmt)], ' ')
	//	}
	//}
	return template.HTML(content)
}

var (
	aEnd  = []byte(`</a>`)
	aHref = []byte(`href="`)
	aID   = []byte(`id="i-`)
)

func disableArticleLink(htmlContent template.HTML, page string) (r template.HTML) {
	content := []byte(htmlContent)
	aStart := []byte(`<a class="index" href="` + page)
	i := bytes.Index(content, aStart)
	if i >= 0 {
		content := content[i:]
		i = bytes.Index(content[len(aStart):], aEnd)
		if i >= 0 {
			i += len(aStart)
			//filleBytes(content[:len(start)], 0)
			//filleBytes(content[i:i+len(end)], 0)
			k := bytes.Index(content, aHref)
			if i >= 0 {
				content[1] = 'b'
				content[i+2] = 'b'
				copy(content[k:], aID)
			}
		}
	}
	return template.HTML(content)
}

//===================================================
// templates
//==================================================

type PageTemplate uint

const (
	Template_Article PageTemplate = iota
	Template_GoGet
	Template_Redirect
	NumPageTemplates
)

var pageTemplates [NumPageTemplates + 1]*template.Template
var pageTemplatesMutex sync.Mutex //
var pageTemplatesCommonPaths = []string{"web", "templates"}

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
			t = parseTemplate(pageTemplatesCommonPaths, "article")
		case Template_GoGet:
			t = parseTemplate(pageTemplatesCommonPaths, "go-get")
		case Template_Redirect:
			t = parseTemplate(pageTemplatesCommonPaths, "redirect")
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
	defer pageTemplatesMutex.Unlock()
	for i := range pageTemplates {
		pageTemplates[i] = nil
	}
}

//===================================================
// non-embedding functions
//===================================================

var dummyHandler http.Handler = http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})

var staticFilesHandler_NonEmbedding = http.FileServer(http.Dir(filepath.Join(rootPath, "web", "static")))

func collectPageGroups_NonEmbedding() map[string]*PageGroup {
	infos, err := ioutil.ReadDir(filepath.Join(rootPath, "pages"))
	if err != nil {
		panic("collect page groups error: " + err.Error())
	}

	pageGroups := make(map[string]*PageGroup, len(infos))

	for _, e := range infos {
		if e.IsDir() {
			group, handler := e.Name(), dummyHandler
			resPath := filepath.Join(rootPath, "pages", group, "res")
			if _, err := os.Stat(resPath); err == nil {
				var urlPrefix string
				// For history reason, fundamentals pages uses "/article/xxx" URLs.
				if group == "fundamentals" {
					urlPrefix = "/article"
				} else if group != "website" {
					urlPrefix = "/" + group
				}
				handler = http.StripPrefix(urlPrefix+"/res/", http.FileServer(http.Dir(resPath)))
			} else if !os.IsNotExist(err) { // !errors.Is(err, os.ErrNotExist) {
				log.Println(err)
			}

			pageGroups[group] = &PageGroup{resHandler: handler}
		}
	}

	return pageGroups
}

func loadArticleFile_NonEmbedding(group, file string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(rootPath, "pages", group, file))
}

func parseTemplate_NonEmbedding(commonPaths []string, files ...string) *template.Template {
	cp := filepath.Join(commonPaths...)
	ts := make([]string, len(files))
	for i, f := range files {
		ts[i] = filepath.Join(rootPath, cp, f)
	}
	return template.Must(template.ParseFiles(ts...))
}

func updateGo101_NonEmbedding() {
	pullGo101Project(rootPath)
}

var rootPath, wdIsGo101ProjectRoot = findGo101ProjectRoot()

func findGo101ProjectRoot() (string, bool) {
	if _, err := os.Stat(filepath.Join(".", "go101.go")); err == nil {
		return ".", true
	}

	for _, name := range []string{
		"gitlab.com/go101/go101",
		"gitlab.com/Go101/go101",
		"github.com/go101/go101",
		"github.com/Go101/go101",
	} {
		pkg, err := build.Import(name, "", build.FindOnly)
		if err == nil {
			return pkg.Dir, false
		}
	}

	return ".", false
}

//===================================================
// utils
//===================================================

func bytesWithLength(s []byte, n int) []byte {
	if n > len(s) {
		n = len(s)
	}
	return s[:n]
}

func filleBytes(s []byte, b byte) {
	for i := range s {
		s[i] = b
	}
}

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
	return hostname == "localhost" // || hostname == "127.0.0.1" // 127.* for local cached version now
}

func runShellCommand(timeout time.Duration, wd string, cmd string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	command := exec.CommandContext(ctx, cmd, args...)
	command.Dir = wd
	return command.Output()
}

func gitPull(wd string) {
	output, err := runShellCommand(time.Minute/2, wd, "git", "pull")
	if err != nil {
		log.Println("git pull:", err)
	} else {
		log.Printf("git pull: %s", output)
	}
}

func goGet(pkgPath, wd string) {
	_, err := runShellCommand(time.Minute/2, wd, "go", "get", "-u", pkgPath)
	if err != nil {
		log.Println("go get -u "+pkgPath+":", err)
	} else {
		log.Println("go get -u " + pkgPath + " succeeded.")
	}
}

//===================================================
// cache
//===================================================

type Cache struct {
	sync.Mutex
	pages map[[2]string][]byte
}

func (c *Cache) Get(group, name string) []byte {
	c.Lock()
	defer c.Unlock()
	page := c.pages[[2]string{group, name}]
	return page
}

func (c *Cache) Set(group, name string, page []byte) {
	c.Lock()
	defer c.Unlock()
	if c.pages == nil {
		c.pages = map[[2]string][]byte{}
	}
	c.pages[[2]string{group, name}] = page
}

func (c *Cache) Clear() {
	c.Lock()
	defer c.Unlock()
	c.pages = map[[2]string][]byte{}
}
