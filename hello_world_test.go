package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Lyndon")
		want := "Hello, Lyndon"

		if got != want {
			t.Errorf("Got '%s' want '%s'", got, want)
		}
	})

	t.Run("Say 'Hello, world.' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world."

		if got != want {
			t.Errorf("Got '%s' want '%s'", got, want)
		}
	})

}
