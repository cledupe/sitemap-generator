package output

import (
	"net/url"
	"regexp"
	"strings"
)

type PageLink struct {
}

func (page PageLink) GetLinks(uri *url.URL, html string) ([]string, error) {
	var links []string
	var finalUrl string
	regexLinkTag := regexp.MustCompile("<base|<a.*?href=\"(.*?)\"")
	linksHtml := regexLinkTag.FindAllStringSubmatch(html, -1)

	for _, link := range linksHtml {
		linkUrlValid, err := url.Parse(strings.TrimSpace(link[1]))
		if err != nil {
			return []string{}, err
		}

		if linkUrlValid.IsAbs() {
			finalUrl = strings.TrimSpace(linkUrlValid.String())
		} else {
			finalUrl = uri.Scheme + "://" + uri.Host + linkUrlValid.String()
		}
		regexFinalUrl := regexp.MustCompile("^https?://")
		regexResult := regexFinalUrl.FindAllStringSubmatch(finalUrl, -1)

		if len(regexResult) > 0 {
			links = append(links, finalUrl)
		}

	}

	return links, nil
}
