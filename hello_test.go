package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Fiona")
	want := "Hello, Fiona"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
