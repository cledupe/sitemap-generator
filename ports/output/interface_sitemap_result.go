package output

type InterfaceSiteMapResult interface {
	SaveSiteMap(filePath string, links []string) error
}
