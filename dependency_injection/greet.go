package main

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, person string) {
	message := "Hello, " + person
	fmt.Fprint(writer, message)
}
