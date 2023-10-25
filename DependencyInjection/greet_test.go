package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {

	buffer := bytes.Buffer{}
	Greet(&buffer, "Adit")

	got := buffer.String()
	want := "Hello, Adit"

	if got != want {
		t.Errorf("Wanted %s, got %s", want, got)
	}
}
