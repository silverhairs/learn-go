package maps

import "errors"

var (
	ErrNotFound      = errors.New("could not find the word you are looking for")
	ErrAlreadyExists = errors.New("cannot add word because it already exists")
)

type Dictionary map[string]string

func (dictionary Dictionary) Search(word string) (string, error) {
	result, ok := dictionary[word]
	if !ok {
		return "", ErrNotFound
	}
	return result, nil
}

func (d Dictionary) Add(word, meaning string) error {
	_, err := d.Search(word)
	if err == nil {
		return ErrAlreadyExists
	}
	d[word] = meaning
	return nil
}

func (d Dictionary) Update(word, meaning string) error {
	_, r := d.Search(word)
	if r == nil {
		d[word] = meaning
		return nil
	}
	return ErrNotFound
}
