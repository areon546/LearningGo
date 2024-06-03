package main

var (
	ErrWordNotFound     = DictionaryErr("key not found in dictionary")
	ErrWordExists       = DictionaryErr("key already present in dictionary")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because word doesn't exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, found := d[key]
	var err error

	if !found {
		err = ErrWordNotFound
	}
	return value, err
}

func (d Dictionary) Add(key, value string) error {

	_, err := d.Search(key)

	switch err {
	case ErrWordNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
