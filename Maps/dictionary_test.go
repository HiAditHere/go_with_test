package maps

import "testing"

func TestDictionary(t *testing.T) {

	t.Run("Known Word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is a test"}
		//got, _ := Search(dictionary, "test")
		got, _ := dictionary.Search("test")
		want := "This is a test"
		assertStrings(t, got, want)
	})

	t.Run("Now known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is a test"}
		//_, err := Search(dictionary, "not test")
		_, err := dictionary.Search("not test")
		assertError(t, err, ErrNotFound)
	})

	t.Run("Add a not existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		_ = dictionary.Add("test", "This is a test")
		assertDefinition(t, dictionary, "test", "This is a test")
	})

	t.Run("Add an existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is a test"}
		err := dictionary.Add("test", "new_test")

		assertError(t, err, ErrAlreadyExists)
	})

	t.Run("Update an existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is a test"}
		new_definition := "This is a new test"

		err := dictionary.Update("test", new_definition)

		if err != nil {
			t.Error("Did not expect an error but got one")
		}
		assertDefinition(t, dictionary, "test", new_definition)
	})

	t.Run("Update an unexisting word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is a test"}
		new_word := "new_test"
		definition := "Some new definition"

		err := dictionary.Update(new_word, definition)

		assertError(t, err, ErrNotFound)
	})

	t.Run("Delete an existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is a test"}
		word := "test"

		_, err := dictionary.Search(word)

		assertNoError(t, err)

		dictionary.Delete(word)
		assertNotFindWord(t, dictionary, word)

	})

	t.Run("Delete an unexisting word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"

		_, err := dictionary.Search(word)

		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Wanted %q, got %q", want, got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("Wanted %q error, got %q error", want, got)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Error("Wanted No error but got one")
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	assertNoError(t, err)
	assertStrings(t, got, definition)
}

func assertNotFindWord(t testing.TB, dictionary Dictionary, word string) {
	t.Helper()

	_, err := dictionary.Search(word)

	assertError(t, err, ErrNotFound)
}
