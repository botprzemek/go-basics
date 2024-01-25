package main

import (
	"fmt"
)

const sixAsInteger int8 = 6
const six = "6"

var integer int8
var hello string
var isPies bool
var pi float32

var letters = [...]string{"a", "b", "c", "d"}
var numbers = [3]int8{1, 2, 3}
var zeros = make([]int8, 3, 5)

var index int8

func main() {
	integer = 21
	hello = "World"
	isPies = false
	pi = 3.14

	//sixAsInteger = 21 // Can't assign new value
	fmt.Println(sixAsInteger)
	fmt.Println(six)

	fmt.Printf("%v Letters: %v\n", len(letters), letters)

	index = 2
	fmt.Printf("Letter at %v: %v\n", index, letters[index])

	slice := letters[1:3]
	fmt.Printf("Sliced letters: %v\n", slice)

	for index := 0; index < len(numbers); index++ {
		fmt.Printf("Number at %v: %v\n", index, numbers[index])
	}

	for index, value := range numbers {
		fmt.Printf("Number at %v: %v\n", index, value)
	}

	fmt.Printf("Zeros: %v\n", zeros)
	zeros = append(zeros, 0)
	fmt.Printf("Zeros: %v\n", zeros)

	people := make(map[string]string)
	people["Lebron"] = "James"
	people["Kevin"] = "Durant"
	people["Stephen"] = "Curry"

	for index, value := range people {
		fmt.Printf("Number at %v: %v\n", index, value)
	}
}
