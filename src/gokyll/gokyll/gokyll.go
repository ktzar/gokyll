package main

import "os"
import "gokyll/files"
import "strings"
import "fmt"

func help() {
	fmt.Println(`usage: gokyll directory`)
}

func copyStatic(siteDir string) {
	dirs := files.GetSiteDirs(siteDir)
	for _, dir := range dirs {
		files.CopyDirectoryToSite(dir, siteDir)
	}
	fmt.Println("Copying static directories: ", strings.Join(dirs, ", "))
}

func main() {
    if len(os.Args) < 2 {
        help()
        return;
    }
    siteDir := os.Args[1]
	files.MakeSiteDir(siteDir)
	copyStatic(siteDir)
    htmlFiles := files.GetHtmlFilesInDir(siteDir)
    for _, html := range htmlFiles {
		files.ProcessFile(siteDir, html)
    }
}


