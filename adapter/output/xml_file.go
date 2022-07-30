package output

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var FILE_NAME = "sitemap.xml"
var DEFAULT_FILE_PATH = ".\\sitemap\\"

type XMLFile struct {
	filePath string
}

var SITEMAP_HEADER = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
var SITEMAP_TAG_START = "\n<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">\n"
var SITEMAP_TAG_END = "</urlset>"

func (xMLFile *XMLFile) setFilePath(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Println(errors.New("Path not exist setting default path"))
		xMLFile.filePath = DEFAULT_FILE_PATH
	} else {
		xMLFile.filePath = filePath
	}

}

func (xMLFile *XMLFile) SaveSiteMap(filePath string, links []string) {
	xMLFile.setFilePath(filePath)
	content := createContent(links)
	if content != "" {
		xMLFile.save([]byte(content))
	} else {
		log.Println(fmt.Errorf("No file created"))
	}

}

func (xMLFile *XMLFile) save(file []byte) {
	_ = ioutil.WriteFile(xMLFile.filePath+FILE_NAME, file, 0644)
}

func createContent(links []string) string {
	var content string
	if len(links) > 0 {
		for _, link := range links {
			content = content + fmt.Sprintf("<url><loc>%s</loc></url>\n", link)
		}
	} else {
		log.Println(fmt.Errorf("No link found"))
	}
	content = SITEMAP_HEADER + SITEMAP_TAG_START + content + SITEMAP_TAG_END
	return content
}
