package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Sean")
	want := "Hello, Sean"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
