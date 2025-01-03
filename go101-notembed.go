//go:build !go1.16
// +build !go1.16

package main

import "html/template"

var staticFilesHandler = staticFilesHandler_NonEmbedding

func collectPageGroups() map[string]*PageGroup {
	return collectPageGroups_NonEmbedding()
}

func loadArticleFile(group, file string) ([]byte, error) {
	return loadArticleFile_NonEmbedding(group, file)
}

func parseTemplate(commonPaths []string, files ...string) *template.Template {
	return parseTemplate_NonEmbedding(commonPaths, files...)
}

func updateGo101() {
	updateGo101_NonEmbedding()
}
