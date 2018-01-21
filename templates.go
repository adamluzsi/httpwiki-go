package httpwiki

type Templates interface {
	ShowPage(data []byte) []byte
	EditPage(data []byte) []byte
}
