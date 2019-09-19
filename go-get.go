package main

import (
	"bytes"
	"net/http"
)

type GoGetInfo struct {
	SubPackage, // assume most one-depth sub-packages
	RootPackage,
	GoGetSourceRepo,
	GoDocSourceRepo string
}

var gogetInfos = map[string]GoGetInfo{
	"tinyrouter": {
		SubPackage:      "",
		RootPackage:     "go101.org/tinyrouter",
		GoGetSourceRepo: "https://github.com/go101/tinyrouter",
		GoDocSourceRepo: "https://github.com/go101/tinyrouter",
	},
	"skia": {
		SubPackage:      "",
		RootPackage:     "go101.org/skia",
		GoGetSourceRepo: "https://github.com/go101/go-skia",
		GoDocSourceRepo: "https://github.com/go101/go-skia",
	},
}

func (go101 *Go101) ServeGoGetPages(w http.ResponseWriter, r *http.Request, rootPkg, subPkg string) {
	info, exists := gogetInfos[rootPkg]
	if !exists {
		http.Redirect(w, r, "/article/101.html", http.StatusNotFound)
		return
	}

	item := rootPkg
	if subPkg != "" {
		item += "/" + subPkg
	}

	page, isLocal := go101.gogetPages.Get(item), go101.IsLocalServer()
	if page == nil {
		t := retrievePageTemplate(Template_GoGet, !isLocal)

		var err error
		var buf bytes.Buffer
		if err = t.Execute(&buf, &info); err == nil {
			page = buf.Bytes()
		} else {
			page = []byte(err.Error())
		}

		if !isLocal {
			go101.articlePages.Set(item, page)
		}
	}

	if isLocal {
		w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	} else {
		w.Header().Set("Cache-Control", "max-age=5000") // about 1.5 hours
	}
	w.Write(page)
}
