package main

import "fmt"

const englishPrefix = "Hello, "
const basicReturn = "Hello, world."

func Hello(name string) string {
	if name == "" {
		return basicReturn
	}
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("world."))
}
