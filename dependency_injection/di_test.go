package main

import "bytes"
import "testing"

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Sean")

	got := buffer.String()
	want := "Hello, Sean"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
