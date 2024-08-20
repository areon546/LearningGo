package main

import (
	"testing"
)

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		dictionaryInput := "test"
		got, _ := dictionary.Search(dictionaryInput)
		want := "this is just a test"

		assertStrings(t, got, want, dictionaryInput)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrWordNotFound

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	dic := Dictionary{}
	searchKey := "a"
	value := "b"
	_ = dic.Add(searchKey, value)

	t.Run("adding a new value", func(t *testing.T) {
		assertDefinition(t, dic, searchKey, value)
	})

	t.Run("overwriting a value", func(t *testing.T) {
		newValue := "new"
		err := dic.Add(searchKey, newValue)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dic, searchKey, value)
	})
}

func TestUpdate(t *testing.T) {
	key := "a"
	value := "b"
	dic := Dictionary{key: value}
	newValue := "new"

	t.Run("updating a present value", func(t *testing.T) {
		dic.Update(key, newValue)

		assertDefinition(t, dic, key, newValue)
	})

	t.Run("updating a missing value", func(t *testing.T) {
		err := dic.Update("key", newValue)

		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	key := "a"
	value := "b"
	dic := Dictionary{key: value}

	t.Run("delete a key that already exists", func(t *testing.T) {
		dic.Delete(key)
		_, err := dic.Search(key)

		if err != ErrWordNotFound {
			t.Errorf("Expected %q to be deleted", key)
		}
	})
}

func assertDefinition(t testing.TB, d Dictionary, key, value string) {
	t.Helper()
	got, err := d.Search(key)

	if err != nil {
		t.Fatalf("should find added word: %v", err)
	}

	assertStrings(t, got, value, key)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want, given string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, given)
	}
}
