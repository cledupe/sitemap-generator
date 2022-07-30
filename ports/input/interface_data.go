package input

//Interface to get data
type InterfaceData interface {
	GetData(link string) (string, error)
}
