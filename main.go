package main

import (
	"flag"
	"github.com/cledupe/sitemap-generator/adapter/input"
	"github.com/cledupe/sitemap-generator/adapter/output"
	"github.com/cledupe/sitemap-generator/service"
	"log"
	"os"
	"runtime"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "", "path's file result")

	var parallel int
	flag.IntVar(&parallel, "parallel", runtime.NumCPU(), "Number of parallel process")

	var outputFile string
	flag.StringVar(&outputFile, "output-file", "", "path's file result")

	var maxDepth int
	flag.IntVar(&maxDepth, "max-depth", 1, "Maximum level deep in url")

	flag.Parse()
	if url == "" {
		log.Println("No url found")
		os.Exit(1)
	}

	if parallel < 1 {
		parallel = runtime.NumCPU()
	}

	var db output.MappingDB
	var pageLink output.PageLink
	var sitemapResult output.XMLFile
	httpClient := input.NewHttpClient()

	siteMapGenerator := service.SiteMapGeneratorInit(outputFile, parallel, maxDepth, &db, pageLink,
		httpClient, &sitemapResult)

	siteMapGenerator.Generate(url)

}
