// learn golang basics
package main

import "fmt"

func basics() {
	// Basic data types
	fmt.Println("===== Basics of Go =====")
	fmt.Println("\n-- Basic Data Types --")

	var a int = 42             // Integer
	var b float64 = 3.14       // Float
	var c string = "Hello Go!" // String
	var d bool = true          // Boolean

	fmt.Println("Integer:", a)
	fmt.Println("Float:", b)
	fmt.Println("String:", c)
	fmt.Println("Boolean:", d)

	// Constants
	fmt.Println("\n-- Constants --")
	const pi = 3.14159
	fmt.Println("Constant Pi:", pi)

	// Arrays and Slices
	fmt.Println("\n-- Arrays and Slices --")
	arr := [3]int{1, 2, 3}  // Fixed-size array
	slice := []int{4, 5, 6} // Dynamic-size slice
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)

	// Map (like a dictionary)
	fmt.Println("\n-- Map --")
	myMap := map[string]int{"A": 1, "B": 2}
	fmt.Println("Map:", myMap)

	// Loops
	fmt.Println("\n-- Loops --")
	for i := 0; i < 5; i++ {
		fmt.Println("Loop iteration:", i)
	}

	// If-Else
	fmt.Println("\n-- If-Else --")
	if d {
		fmt.Println("d is true!")
	} else {
		fmt.Println("d is false!")
	}

	// Functions
	fmt.Println("\n-- Functions --")
	result := add(10, 20)
	fmt.Println("10 + 20 =", result)

	// Structs
	fmt.Println("\n-- Structs --")
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Jester", Age: 25}
	fmt.Println("Struct Person:", person)

	// Pointers
	fmt.Println("\n-- Pointers --")
	x := 100
	pointer := &x
	fmt.Println("Value of x:", x)
	fmt.Println("Pointer to x:", pointer)
	fmt.Println("Value at pointer:", *pointer)
}

// A simple function to add two numbers
func add(x int, y int) int {
	return x + y
}

// Go's main function is the entry point of any Go program.
func main() {
	basics()
	println(add(5, 12))
	// called this function from hello_world.go file
	Hello()
}
