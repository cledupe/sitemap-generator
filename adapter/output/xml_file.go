package output

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const FileName = "sitemap.xml"
const DefaultFilePath = ".\\sitemap\\"
const SitemapHeader = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
const SitemapTagStart = "\n<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">\n"
const SitemapTagEnd = "</urlset>"

type XMLFile struct {
	filePath string
}

func (xMLFile *XMLFile) SaveSiteMap(filePath string, links []string) error {
	err := xMLFile.setFilePath(filePath)
	if err != nil {
		return err
	}
	content, err := createContent(links)
	if err == nil {
		err = xMLFile.save([]byte(content))
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println(err)
	}
	return err
}

func (xMLFile *XMLFile) setFilePath(filePath string) error {
	var err error
	filePathPrepared := prepareFilePath(filePath)

	if _, err = os.Stat(filePathPrepared); os.IsNotExist(err) {
		err = errors.New("path not exist setting default path")
		xMLFile.filePath = DefaultFilePath
	} else {
		xMLFile.filePath = filePathPrepared
	}
	return err
}

func (xMLFile *XMLFile) save(file []byte) error {
	return ioutil.WriteFile(filepath.Join(xMLFile.filePath, FileName), file, 0644)
}

func createContent(links []string) (string, error) {
	var content string
	var err error
	if len(links) > 0 {
		for _, link := range links {
			content = content + fmt.Sprintf("<url><loc>%s</loc></url>\n", link)
		}
		content = SitemapHeader + SitemapTagStart + content + SitemapTagEnd
	} else {
		err = fmt.Errorf("no link found")
	}

	return content, err
}

func prepareFilePath(filePath string) string {
	absolutePath, _ := filepath.Abs(filePath)
	return absolutePath
}
