package httpwiki_test

type Templates struct{}

func (t Templates) ShowPage(data []byte) []byte {
	return []byte("show: " + string(data))
}

func (t Templates) EditPage(data []byte) []byte {
	return []byte("edit: " + string(data))
}
