package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	var tests = []struct {
		name     string
		expected float64
		given    Shape
	}{
		{"Rectangle", 72.0, Rectangle{12, 6}},
		{"Circle", 314.1592653589793, Circle{10}},
		{"Triangle", 36.0, Triangle{12, 6}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := (tt.given.Area())
			if got != tt.expected {
				t.Errorf("got %.2f expected %.2f", got, tt.expected)
			}
		})
	}
}
