package output

//Interface save result
type InterfaceLink interface {
	Save(url string)
	FindAll() []string
}
