package helpers

import (
	"testing"
)

func TestRemoveNewlineFromStr(t *testing.T) {

	t.Run("removes single newline at end of string", func(t *testing.T) {
		initialString := "hello this is a string\n"
		got := RemoveNewlineFromStr(initialString)
		want := "hello this is a string"

		if got != want {
			t.Fatalf("formatted string still contains \n")
		}
	})
	t.Run("removes single newline in middle of string", func(t *testing.T) {
		initialString := "hello this \nis a string"
		got := RemoveNewlineFromStr(initialString)
		want := "hello this is a string"

		if got != want {
			t.Errorf("wanted %s got %s", want, got)
		}
	})
	t.Run("removes two newline from string", func(t *testing.T) {
		initialString := "hello this \nis a string\n"
		got := RemoveNewlineFromStr(initialString)
		want := "hello this is a string"

		if got != want {
			t.Errorf("wanted %s got %s", want, got)
		}
	})

	t.Run("removes two newline from string", func(t *testing.T) {
		initialString := "hello \n\nthis \nis a string\n"
		got := RemoveNewlineFromStr(initialString)
		want := "hello this is a string"

		if got != want {
			t.Errorf("wanted %s got %s", want, got)
		}
	})

	t.Run("removes newline followed by space from string", func(t *testing.T) {
		initialString := "hello \n\n this \nis a string\n"
		got := RemoveNewlineFromStr(initialString)
		want := "hello this is a string"

		if got != want {
			t.Errorf("wanted %s got %s", want, got)
		}
	})
}
