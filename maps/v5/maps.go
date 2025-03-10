package maps

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (dictionary Dictionary) Search(word string) (string, error) {
	definition, ok := dictionary[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (dictionary Dictionary) Add(word, definition string) error {
	_, err := dictionary.Search(word)
	switch err {
	case ErrNotFound:
		dictionary[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err

	}

	return nil
}

func (dictionary Dictionary) Update(word, definition string) {
	dictionary[word] = definition
}
