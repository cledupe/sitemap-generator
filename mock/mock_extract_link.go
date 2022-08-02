package mock

import (
	"net/url"
)

type MockExtractLink struct {
}

func (mel MockExtractLink) GetLinks(uri *url.URL, html string) ([]string, error) {

	switch uri.String() {
	case "https://example1.com.br":
		return []string{"https://example2.com.br"}, nil
	case "https://example2.com.br":
		return []string{"https://example3.com.br"}, nil
	case "https://example3.com.br":
		return []string{"https://example4.com.br"}, nil
	}
	return []string{""}, nil
}
