package main

import (
	"bytes"
	"net/http"
	"strings"
)

type GoGetInfo struct {
	RootPackage,
	GoGetSourceRepo, // only supports github now
	GoDocWebsite string
}

// ToDo: retire the SubPackage field.
var gogetInfos = map[string]GoGetInfo{
	"tinyrouter": {
		RootPackage:     "go101.org/tinyrouter",
		GoGetSourceRepo: "go101/tinyrouter",
		GoDocWebsite:    "https://pkg.go.dev/",
	},
	"skia": {
		RootPackage:     "go101.org/skia",
		GoGetSourceRepo: "go101/go-skia",
		GoDocWebsite:    "https://pkg.go.dev/",
	},
	"go101": {
		RootPackage:     "go101.org/go101",
		GoGetSourceRepo: "go101/go101",
	},
	"golang101": {
		RootPackage:     "go101.org/golang101",
		GoGetSourceRepo: "golang101/golang101",
	},
	"gold": {
		RootPackage:     "go101.org/gold",
		GoGetSourceRepo: "go101/gold",
		GoDocWebsite:    "https://pkg.go.dev/",
	},
	"golds": {
		RootPackage:     "go101.org/golds",
		GoGetSourceRepo: "go101/golds",
		GoDocWebsite:    "https://pkg.go.dev/",
	},
	"ebooktool": {
		RootPackage:     "go101.org/ebooktool",
		GoGetSourceRepo: "go101/ebooktool",
		GoDocWebsite:    "https://pkg.go.dev/",
	},
	"nstd": {
		RootPackage:     "go101.org/nstd",
		GoGetSourceRepo: "go101/nstd",
		GoDocWebsite:    "https://pkg.go.dev/",
	},
	"gotv": {
		RootPackage:     "go101.org/gotv",
		GoGetSourceRepo: "go101/gotv",
		GoDocWebsite:    "https://pkg.go.dev/",
	},
	"tmd.go": {
		RootPackage:     "go101.org/tmd.go",
		GoGetSourceRepo: "go101/tmd.go",
		GoDocWebsite:    "https://pkg.go.dev/",
	},
	"tmd": {
		RootPackage:     "go101.org/tmd",
		GoGetSourceRepo: "go101/tmd",
		GoDocWebsite:    "https://pkg.go.dev/",
	},
	"godev": {
		RootPackage:     "go101.org/godev",
		GoGetSourceRepo: "go101/godev",
		GoDocWebsite:    "https://pkg.go.dev/",
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
		if subPkg == "" {
			if rootPkg == "" {
				rootPkg = "index.html"
			}
			go101.serveGroupItem(w, r, "website", rootPkg)
		} else {
			http.Redirect(w, r, "/", http.StatusNotFound)
		}
		return
	}

	item := rootPkg
	if subPkg != "" {
		item += "/" + subPkg
	}

	page, isLocal := go101.gogetPages.Get(item, version), go101.IsLocalServer()
	if page == nil {
		info.GoGetSourceRepo = "https://github.com/" + info.GoGetSourceRepo
		if info.GoDocWebsite != "" {
			info.GoDocWebsite += info.RootPackage + "/" + subPkg + version
		} else {
			info.GoDocWebsite = info.GoGetSourceRepo
			if subPkg != "" {
				info.GoDocWebsite += "/tree/master/" + subPkg
			}
		}

		var err error
		var buf bytes.Buffer
		t := retrievePageTemplate(Template_GoGet, !isLocal)
		if err = t.Execute(&buf, &info); err == nil {
			page = buf.Bytes()
		} else {
			page = []byte(err.Error())
		}

		if !isLocal {
			go101.gogetPages.Set(item, version, page)
		}
	}

	if isLocal {
		w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	} else {
		w.Header().Set("Cache-Control", "max-age=50000") // about 14 hours
	}
	w.Write(page)
}
