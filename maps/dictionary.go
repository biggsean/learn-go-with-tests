package main

// Dictionary is a map od strings
type Dictionary map[string]string

const (
	// ErrNotFound error
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists error
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	// ErrWordDoesNotExists error
	ErrWordDoesNotExists = DictionaryErr("cannot update word because it does not exist")
)

// DictionaryErr type
type DictionaryErr string

// Error for type DictionaryErr
func (e DictionaryErr) Error() string {
	return string(e)
}

//Search is a search
func (d Dictionary) Search(w string) (string, error) {
	definition, ok := d[w]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add puts a word in the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update a definition in the dictionary
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

// Delete a word from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
