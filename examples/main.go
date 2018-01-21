package main

import (
	"errors"
	"net/http"

	httpwiki "github.com/adamluzsi/httpwiki-go"
)

func main() {
	wiki := httpwiki.New(NewStorage(), Templates{})

	http.ListenAndServe(":8080", wiki)
}

type Storage struct {
	memory map[string][]byte
}

func (s *Storage) Save(name string, content []byte) error {
	s.memory[name] = content
	return nil
}

func (s *Storage) Load(name string) ([]byte, error) {
	data, found := s.memory[name]

	var err error

	if !found {
		err = errors.New("not found!")
	}

	return data, err
}

func NewStorage() *Storage {
	return &Storage{memory: map[string][]byte{}}
}

type Templates struct{}

func (t Templates) ShowPage(data []byte) []byte {
	return data
}

func (t Templates) EditPage(data []byte) []byte {
	return data
}
