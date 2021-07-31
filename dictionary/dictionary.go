package dictionary

import "errors"

// Dictionary Type
type Dictionary map[string]string

var errNotFound = errors.New("Not found")
var errWordExist = errors.New("That word already exists")
var errWordCantUpdate = errors.New("That word cannot be updated")
var errWordCantDelete = errors.New("That word does not exist")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExist
	}
	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errWordCantUpdate
	case nil:
		d[word] = def
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errWordCantDelete
	case nil:
		delete(d, word)
	}
	return nil
}