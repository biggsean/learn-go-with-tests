package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestWalk(t *testing.T) {
	var tests = []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Sean"},
			[]string{"Sean"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Sean", "Great Falls"},
			[]string{"Sean", "Great Falls"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Sean", 42},
			[]string{"Sean"},
		},
		{
			"Nested structs",
			StructWithNestedStruct{
				"Sean",
				Profile{42, "Great Falls"},
			},
			[]string{"Sean", "Great Falls"},
		},
		{
			"Pointers to things",
			&StructWithNestedStruct{
				"Sean",
				Profile{42, "Great Falls"},
			},
			[]string{"Sean", "Great Falls"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		{
			"Maps",
			map[string]string{
				"foo": "bar",
				"biz": "baz",
			},
			[]string{"bar", "baz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string
			walk(tt.Input, func(input string) {
				got = append(got, input)
			})

			if reflect.ValueOf(tt.Input).Kind() == reflect.Map {
				sort.Strings(got)
			}

			if !reflect.DeepEqual(got, tt.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, tt.ExpectedCalls)
			}
		})
	}
}

type StructWithNestedStruct struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
