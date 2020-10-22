package main

import (
	"bytes"
	"log"
	"net/http"
)

var redirectPages = map[string]string{
	"go-sdk.html":    "go-toolchain.html",
	"tool-gold.html": "tool-golds.html",
}

func (go101 *Go101) RedirectArticlePage(w http.ResponseWriter, r *http.Request, file string) bool {
	redirectPage, ok := redirectPages[file]
	if ok {
		page, isLocal := go101.articlePages.Get(file), go101.IsLocalServer()
		if page == nil {
			pageParams := map[string]interface{}{
				"RedirectPage":  "/article/" + redirectPage,
				"IsLocalServer": isLocal,
				"Value": func() func(string, ...interface{}) interface{} {
					var kvs = map[string]interface{}{}
					return func(k string, v ...interface{}) interface{} {
						if len(v) == 0 {
							return kvs[k]
						}
						kvs[k] = v[0]
						return ""
					}
				}(),
			}

			t := retrievePageTemplate(Template_Redirect, !isLocal)
			var buf bytes.Buffer
			if err := t.Execute(&buf, pageParams); err == nil {
				page = buf.Bytes()
			} else {
				page = []byte(err.Error())
			}

			if !isLocal {
				go101.articlePages.Set(file, page)
			}
		}

		if len(page) == 0 { // blank page means page not found.
			log.Printf("article page %s is not found", file)
			//w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
			http.Redirect(w, r, "/article/101.html", http.StatusNotFound)
		} else if isLocal {
			w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
		} else {
			w.Header().Set("Cache-Control", "max-age=50000") // about 14 hours
		}
		w.Write(page)
	}

	return ok
}
