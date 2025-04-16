package main

import (
	"fmt"
	"io"
	"os"
)

/**
- Create a program that reads the contents of a text file then prints its contents out to the terminal
- The file to open should be provided as an argument to the program when it is executed at the terminal. For example, go run main.go mytextfile.txt
- Use os.Args to read the arguments provided to program
-
*/
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a valid file")
	}

	file, err := os.Open(os.Args[1])
	if (err != nil) {
		fmt.Printf("found err %v when trying to open file", err)
	}
	io.Copy(os.Stdout, file)
}