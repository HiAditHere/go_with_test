package maps

import (
	"errors"
)

type Dictionary map[string]string

var (
	ErrNotFound      = errors.New("Could not find what you were looking for")
	ErrAlreadyExists = errors.New("The word already exists in the dictionary")
)

func (d Dictionary) Search(word string) (string, error) {
	meaning, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return meaning, nil
}

func (d Dictionary) Add(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, new_definition string) error {

	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = new_definition
	case ErrNotFound:
		return ErrNotFound
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
