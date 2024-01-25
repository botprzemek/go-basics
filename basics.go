package main

import (
	"fmt"
)

var godzilla string

func greet(arg string) (name string) {
	fmt.Printf("Hello, %s!\n", arg)
	return
}

func main() {
	fmt.Println("Hello, World!")

	const hello = "World"
	fmt.Printf("Hello, %s!\n", hello)

	var a, b = "Hello", 6
	fmt.Print(a, ", ", b, "!\n")
	fmt.Printf("%v is type of %T, but %v is type of %T\n", a, a, b, b)

	/* fmt.Println("This is a comment")
	fmt.Println("This is also a comment") */
	fmt.Println("This is not a comment") // But this won't print

	godzilla = "Godzilla"
	greet(godzilla)
}
