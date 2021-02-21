// +build !go1.16

package main

var staticFilesHandler = staticFilesHandler_NonEmbedding
var resFilesHandler = resFilesHandler_NonEmbedding

func loadArticleFile(file string) ([]byte, error) {
	return loadArticleFile_NonEmbedding(file)
}

func parseTemplate(commonPaths []string, files ...string) *template.Template {
	parseTemplate_NonEmbedding(commonPaths, files...)
}

func updateGo101() {
	updateGo101_NonEmbedding()
}
