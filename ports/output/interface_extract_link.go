package output

import "net/url"

type InterfaceExtractLink interface {
	GetLinks(uri *url.URL, html string) ([]string, error)
}
