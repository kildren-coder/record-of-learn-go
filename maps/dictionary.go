package maps

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition,nil
}

func (d Dictionary) Add(word, str string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = str
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

/*
func Search(dictionary map[string]string, word string) string{
	return dictionary[word]
}
*/

//https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/maps#shou-xian-bian-xie-ce-shi-4
//https://github.com/studygolang/learn-go-with-tests/blob/master/maps.md#write-the-test-first-4