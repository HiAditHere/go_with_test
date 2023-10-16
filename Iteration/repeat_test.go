package iteration

import "testing"

func TestRepeat(t *testing.T) {
	get := Repeat("a")
	want := "aaaaa"

	if get != want {
		t.Errorf("Wanted '%q' got '%q'", want, get)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Repeat("a")
	}
}
