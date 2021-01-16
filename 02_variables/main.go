package main

import "fmt"

// Variable is the name given to a memory location to store a value of a specific type.

// Declaring a variable with an initial value
// A variable can also be provided an initial value when it is declared. The following is the syntax to declare a variable with an initial value.

func main() {
	// single variable declaration with initial value
	var name string = "Kakashi Hatake"
	// Multiple variable declaration  
	var width, height int = 100, 50

// The short declaration := can only be used withing a function and not outside 
	a := 10
	b := "golang"
	c := 4.17
	


	// This will print the type and value of the variables declared above 
	fmt.Println("width is", width, "height is", height)
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", name)


	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", name)
}
