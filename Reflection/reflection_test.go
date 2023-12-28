package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Location string
	Age      int
}

func TestWalk(t *testing.T) {

	want := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Input is as Expected Calls",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},

		{
			"More than 1 input",
			struct {
				Name    string
				Surname string
			}{"Chris", "Brown"},
			[]string{"Chris", "Brown"},
		},

		{
			"If all input is not string",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},

		{
			"If some input is struct",
			&Person{
				"Chris",
				Profile{"London", 33},
			},
			[]string{"Chris", "London"},
		},

		{
			"Slices",
			[]Profile{
				{"London", 33},
				{"Reykjavik", 34},
			},
			[]string{"London", "Reykjavik"},
		},
	}

	for _, test := range want {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("wanted %v got %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("Testing for Maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string

		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("Testing for channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{"London", 33}
			aChannel <- Profile{"Katowice", 34}
			close(aChannel)
		}()

		var got []string

		want := []string{"London", "Katowice"}

		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %q, got %q", want, got)
		}
	})

	t.Run("Testing for functions", func(t *testing.T) {

		aFunc := func() (Profile, Profile) {
			return Profile{"London", 33}, Profile{"Katowice", 34}
		}

		var got []string

		want := []string{"London", "Katowice"}

		Walk(aFunc, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if contains == false {
		t.Errorf("expecred %v to contain %q but did not", haystack, needle)
	}
}
