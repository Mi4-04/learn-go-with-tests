package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct{ Name string }{"man"},
			[]string{"man"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"man", "London"},
			[]string{"man", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"man", 33},
			[]string{"man"},
		},
		{
			"nested fields",
			Person{
				"man",
				Profile{33, "London"},
			},
			[]string{"man", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
