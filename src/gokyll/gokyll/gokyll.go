package main

import "os"
import "gokyll/files"
import "fmt"

func help() {
	fmt.Println(`usage: gokyll directory`)
}

func main() {
    if len(os.Args) < 2 {
        help()
        return;
    }
    siteDir := os.Args[1]
	files.MakeSiteDir(siteDir)
    dirs := files.GetSiteDirs(siteDir)
	for _, dir := range dirs {
		fmt.Println("Copying: ", dir)
		files.CopyDirectoryToSite(dir, siteDir)
	}
    htmlFiles := files.GetHtmlFilesInDir(siteDir)
    for _, html := range htmlFiles {
		files.ProcessFile(siteDir, html)
    }
}


