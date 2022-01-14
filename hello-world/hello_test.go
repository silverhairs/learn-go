package main

import "testing"

func TestHello(tester *testing.T) {
	assertMessage := func(t testing.TB, matcher, expect string) {
		t.Helper()
		if matcher != expect {
			t.Errorf("expected %q, actual %q", expect, matcher)
		}
	}

	tester.Run("greets people", func(t *testing.T) {
		const name = "John"
		want := prefix + name
		got := Hello("John")

		assertMessage(t, got, want)
	})

	tester.Run("says hello when passed an empty string", func(t *testing.T) {
		want := prefix + "World"
		got := Hello("")

		assertMessage(t, got, want)
	})

}
