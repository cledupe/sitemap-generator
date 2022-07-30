package service

import (
	"fmt"
	"github.com/cledupe/sitemap-generator/ports/input"
	"github.com/cledupe/sitemap-generator/ports/output"
	"github.com/cledupe/sitemap-generator/utils/concurrency_pattern"
	"log"
	"net/url"
)

type SitemapGenerator struct {
	filePath      string
	parallel      int
	level         int
	db            output.InterfaceLink
	extractLinks  output.InterfaceExtractLink
	urlContent    input.InterfaceData
	siteMapResult output.InterfaceSiteMapResult
}

func SiteMapGeneratorInit(filePath string, parallel int, level int, db output.InterfaceLink,
	extractLinks output.InterfaceExtractLink, urlContent input.InterfaceData, siteMapResult output.InterfaceSiteMapResult) SitemapGenerator {
	return SitemapGenerator{
		filePath:      filePath,
		parallel:      parallel,
		level:         level,
		db:            db,
		extractLinks:  extractLinks,
		urlContent:    urlContent,
		siteMapResult: siteMapResult,
	}
}

func (sitemapGenerator SitemapGenerator) Generate(initialLink string) {
	var level = 0

	var links []string
	links = append(links, initialLink)
	for level <= sitemapGenerator.level && len(links) > 0 {
		links = sitemapGenerator.generate(links, level)
		level++
	}
	sitemapGenerator.siteMapResult.SaveSiteMap(sitemapGenerator.filePath, sitemapGenerator.db.FindAll())
}

func (sitemapGenerator SitemapGenerator) generate(siteUrls []string, level int) []string {
	var resultsLink []string
	workerPool := concurrency_pattern.NewThreadPool(sitemapGenerator.parallel)
	workerPool.Run()

	tasksNumber := len(siteUrls)

	channels := make(chan []string, tasksNumber)
	for _, siteUrl := range siteUrls {
		//Remove reference of siteUrls - Race Data issue
		copySiteUrl := fmt.Sprintf("%s", siteUrl)
		workerPool.AddTask(
			func() {
				sitemapGenerator.getLinks(copySiteUrl, level, channels)
			})
	}

	for i := 0; i < tasksNumber; i++ {
		resultsLink = append(resultsLink, <-channels...)
	}

	return resultsLink
}

func (sitemapGenerator SitemapGenerator) getLinks(siteUrl string, level int, ch chan []string) {
	var arrayLink []string
	urlParsed, err := url.ParseRequestURI(siteUrl)
	if err == nil {
		sitemapGenerator.db.Save(siteUrl)
		if level < sitemapGenerator.level {
			content, err := sitemapGenerator.urlContent.GetData(siteUrl)
			if err == nil {
				arrayLink, err = sitemapGenerator.extractLinks.GetLinks(urlParsed, content)
				if err != nil {
					log.Println(err)
				}
			} else {
				log.Println(err)
			}
		}

	} else {
		log.Println(err)
	}

	ch <- arrayLink
}
