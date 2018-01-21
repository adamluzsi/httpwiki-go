package httpwiki_test

import (
	"errors"
)

type Storage struct {
	memory map[string][]byte
	errors []error
}

func (s *Storage) expectedError() error {
	if s.errors == nil || len(s.errors) == 0 {
		return nil
	}

	err := s.errors[len(s.errors)-1]
	s.errors = s.errors[:len(s.errors)-1]

	return err
}

func (s *Storage) Save(name string, content []byte) error {

	if expectedErr := s.expectedError(); expectedErr != nil {
		return expectedErr
	}

	s.memory[name] = content

	return nil

}

func (s *Storage) Load(name string) ([]byte, error) {

	if expectedErr := s.expectedError(); expectedErr != nil {
		return nil, expectedErr
	}

	data, found := s.memory[name]

	var err error

	if !found {
		err = errors.New("not found!")
	}

	return data, err
}

func NewStorage(es []error) *Storage {
	return &Storage{memory: map[string][]byte{}, errors: es}
}
