package main

import "os"
import "files"
import "fmt"

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
    dirs := files.GetSiteDirs(siteDir)
	for _, dir := range dirs {
		fmt.Println("Copying: ", dir)
		files.CopyDirectoryToSite(dir, siteDir)
	}
}


