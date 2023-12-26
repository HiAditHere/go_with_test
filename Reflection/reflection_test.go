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
				age  int
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
}
