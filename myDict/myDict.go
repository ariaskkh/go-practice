package myDict

import "errors"

// aloas
type Dictionary map[string] string 

var (
	errNotFound = errors.New("Not Found")
	errCantUpdate = errors.New("Can't update non-existing word")
	errCantDelete = errors.New("Can't delete non-existing word")
	errWordExists = errors.New("That word already exists")
)


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
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}