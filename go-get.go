package main

import (
	"bytes"
	"net/http"
	"strings"
)

type GoGetInfo struct {
	SubPackage, // assume most one-depth sub-packages
	RootPackage,
	GoGetSourceRepo,
	GoDocSourceRepo,
	// starts with '@'
	Version string
}

// ToDo: retire the SubPackage field.
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
	"go101": {
		SubPackage:      "",
		RootPackage:     "go101.org/go101",
		GoGetSourceRepo: "https://github.com/go101/go101",
		GoDocSourceRepo: "https://github.com/go101/go101",
	},
	"golang101": {
		SubPackage:      "",
		RootPackage:     "go101.org/golang101",
		GoGetSourceRepo: "https://github.com/golang101/golang101",
		GoDocSourceRepo: "https://github.com/golang101/golang101",
	},
	"gold": {
		SubPackage:      "",
		RootPackage:     "go101.org/gold",
		GoGetSourceRepo: "https://github.com/go101/gold",
		GoDocSourceRepo: "https://github.com/go101/gold",
	},
	"golds": {
		SubPackage:      "",
		RootPackage:     "go101.org/golds",
		GoGetSourceRepo: "https://github.com/go101/golds",
		GoDocSourceRepo: "https://github.com/go101/golds",
	},
}

func (go101 *Go101) ServeGoGetPages(w http.ResponseWriter, r *http.Request, rootPkg, subPkg string) {
	var version string
	if subPkg != "" {
		atIndex := strings.IndexByte(subPkg, '@')
		if atIndex >= 0 {
			subPkg = subPkg[:atIndex]
			version = subPkg[atIndex:]
		}
	} else {
		atIndex := strings.IndexByte(rootPkg, '@')
		println(111, atIndex)
		if atIndex > 0 {
			version = rootPkg[atIndex:]
			rootPkg = rootPkg[:atIndex]
		}
	}

	// simple handling for pkg.go.dev
	if len(version) < 3 || version[1] != 'v' || version[2] < '0' || version[2] > '9' {
		version = ""
	}

	info, exists := gogetInfos[rootPkg]
	if !exists {
		http.Redirect(w, r, "/article/101.html", http.StatusNotFound)
		return
	}

	item := rootPkg
	if subPkg != "" {
		item += "/" + subPkg
	}
	item += version

	page, isLocal := go101.gogetPages.Get(item), go101.IsLocalServer()
	if page == nil {
		t := retrievePageTemplate(Template_GoGet, !isLocal)
		info.SubPackage = subPkg
		info.Version = version
		println("info.Version=", info.Version)

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
		w.Header().Set("Cache-Control", "max-age=50000") // about 14 hours
	}
	w.Write(page)
}
