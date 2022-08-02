package mock

type MockGetData struct {
}

func (mgd MockGetData) GetData(url string) (string, error) {
	url = ""
	return url, nil
}
