package main

import (
	"os"
	"files"
)

func main() {
    if len(os.Args) < 2 {
        files.Help()
        return;
    }
    siteDir := os.Args[1]
	files.MakeSiteDir(siteDir)
    htmlFiles := files.GetHtmlFilesInDir(siteDir)
    for _, html := range htmlFiles {
		files.ProcessFile(siteDir, html)
    }
}


