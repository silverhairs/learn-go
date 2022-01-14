package main

import "fmt"

func main() {

	const name = "Boris"
	message := Hello(name)
	fmt.Println(message)
}

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return prefix + name
}

const prefix = "Hello, "
