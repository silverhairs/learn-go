package maps

import (
	"testing"
)

var assertString = func(t testing.TB, expect, got string) {
	t.Helper()
	if expect != got {
		t.Errorf("expected %s, but got %s", expect, got)
	}
}

var assertError = func(t testing.TB, expect, got error) {
	t.Helper()
	if expect != got {
		t.Fatal("expected an error but got none")
	}
}

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("Search", func(t *testing.T) {
		t.Run("search a word in the dictionary", func(t *testing.T) {
			got, _ := dictionary.Search("test")
			expect := "this is just a test"
			assertString(t, expect, got)
		})

		t.Run("returns an error when the item is not found", func(t *testing.T) {
			_, err := dictionary.Search("boris9")
			expect := ErrNotFound

			if err != nil {
				assertError(t, expect, err)
			}
		})
	})

	t.Run("Add", func(t *testing.T) {
		t.Run("new word", func(t *testing.T) {
			dictionary.Add("boris", "Awesome person")
			expect := "Awesome person"
			got, _ := dictionary.Search("boris")
			assertString(t, expect, got)
		})
		t.Run("existing word", func(t *testing.T) {
			word := "test"
			def := "this is just a test"
			err := dictionary.Add(word, "another test")
			want, _ := dictionary.Search(word)
			assertError(t, ErrAlreadyExists, err)
			assertString(t, def, want)
		})
	})
	t.Run("Update", func(t *testing.T) {
		new_word := "not a test"
		t.Run("existing word", func(t *testing.T) {
			dictionary.Update("test", new_word)
			got, _ := dictionary.Search("test")
			assertString(t, new_word, got)
		})
		t.Run("word not found", func(t *testing.T) {
			expect := ErrNotFound
			dictionary.Update("testi", new_word)
			_, got := dictionary.Search("testi")
			assertError(t, expect, got)
		})
	})
}
