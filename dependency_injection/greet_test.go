package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	fmt.Println(buffer.String())
	expect := "Hello, Chris"

	if got != expect {
		t.Errorf("expected '%s' , but got '%s'", expect, got)
	}
}
