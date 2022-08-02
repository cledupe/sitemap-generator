package service

import (
	"github.com/cledupe/sitemap-generator/adapter/output"
	"github.com/cledupe/sitemap-generator/mock"
	"log"
	"os"
	"strings"
	"testing"
)

const InitialLink = "https://example1.com.br"

const FileName = "sitemap.xml"

var contentFake = []string{"<loc>https://example1.com.br</loc>",
	"<loc>https://example2.com.br</loc>",
	"<loc>https://example3.com.br</loc>",
	"<loc>https://example4.com.br</loc>"}

func TestGenerateSiteMap(t *testing.T) {

	db := output.MappingDB{}
	el := mock.MockExtractLink{}
	uc := mock.MockGetData{}
	sr := output.XMLFile{}
	sm := SiteMapGeneratorInit("", 8, 3, &db, el, uc, &sr)

	sm.Generate(InitialLink)
	data, err := os.ReadFile(FileName)
	if err != nil {
		t.Error(err.Error())
	}
	content := string(data)
	for _, link := range contentFake {
		if !strings.Contains(content, link) {
			t.Errorf("Link %s not contains in %s the content \n", link, content)
		}
	}
	e := os.Remove(FileName)
	if e != nil {
		log.Fatal(e)
	}
}
