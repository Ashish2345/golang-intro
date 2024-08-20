package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	var want string = "Hello Test"
	var got string = Say("Hello Test")

	if want != got {
		t.Errorf("Wanted %s, got %s", want, got)
	}

}
