package output_test

import (
	"github.com/cledupe/sitemap-generator/adapter/output"
	"net/url"
	"testing"
)

func TestExtractLink(t *testing.T) {
	var pageLink output.PageLink
	var html = `<!DOCTYPE html>
					<html>
					<body>
					
					<h1>The a href attribute</h1>
					
					<p>An absolute URL: <a href="https://www.w3schools.com">W3Schools</a></p>
					<p>A relative URL: <a href="tag_a.asp">The a tag</a></p>
					<p>mail link: <a href="mailto:abc@example.com">The a tag</a></p>
					
					</body>
					</html>`
	link, _ := url.Parse("https://www.w3schools.com")

	links, err := pageLink.GetLinks(link, html)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	if len(links) != 2 {
		t.Errorf("Array size is %d instead %d", len(links), 2)
	}

}
