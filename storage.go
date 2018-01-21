package httpwiki

type Storage interface {
	Save(name string, content []byte) error
	Load(name string) ([]byte, error)
}
