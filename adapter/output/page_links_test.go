package output

import (
	"net/url"
	"testing"
)

func extractLink(t *testing.T) {
	var pageLink PageLink
	var html = `<!DOCTYPE html>
					<html>
					<body>
					
					<h1>The a href attribute</h1>
					
					<p>An absolute URL: <a href="https://www.w3schools.com">W3Schools</a></p>
					<p>A relative URL: <a href="tag_a.asp">The a tag</a></p>
					
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
