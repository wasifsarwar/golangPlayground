package main

import "fmt"

func main() {
	//key string, value string

	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "someGreenHex",
	}

	//add a key
	colors["white"] = "#hexCodeForWHITE"

	printMap(colors)

	//delete a key
	delete(colors, "white")
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("the hex for %v is %v\n", color, hex)
	}
}

/**
Difference between structs and maps

Structs
- Values can be of different type
- Keys don't support indexing
- You need to know all different fields at compile time
- Used to represent a thing with a lot of different properties
- THIS IS A VALUE TYPE SO NEEDS POINTERS TO CHANGE VALUES


Maps
- All keys & values must be of the same type
- Keys are indexed, you can iterate over them
- You don't need to know all the keys at compile time
- Used to represent a collection of related properties
- THIS IS A REFERENCE TYPE, DOESN'T NEED POINTERS TO CHANGE VALUES

*/
