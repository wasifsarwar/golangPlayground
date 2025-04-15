package main

import "fmt"

/**
 Value types
 - int
 - float
 - string
 - bool
 - structs
WE NEED POINTERS TO CHANGE THE VALUE OF THESE THINGS IN A FUNCTION
*/

/**
Reference types
- slices
- maps
- channels
- pointers
- functions

We don't need pointers to reference a value, as these types underneath points to the underlying value even when passed as a copy value in a function
*/

type person struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func (p person) details() {
	fmt.Printf("FirstName : %v\n", p.firstName)
	fmt.Printf("LastName : %v\n", p.lastName)
	fmt.Printf("email : %v\n", p.contactInfo.email)
	fmt.Printf("zip : %v\n", p.contactInfo.zip)
}

/**
 - pointers in go
 - go is a pass by value language
 - In Go, "pass by value" means that when you pass a variable to a function,
   a copy of that variable is made. This means that any modifications made to the parameter inside the function do not affect the original variable outside the function.

   Key Points:

   Value Copying: When you pass a variable (like a struct, slice, or primitive type) to a function, Go creates a copy of that variable. Changes to the parameter inside the function do not reflect on the original variable.
   Pointer Receivers: To modify the original variable, you can pass a pointer to it. This way, the function can access and modify the original data. For example, in your updateName method, using a pointer receiver (*person) allows the method to change the firstName of the original person instance.
   Performance Considerations: Passing large structs by value can be inefficient because it involves copying the entire struct. In such cases, using pointers can improve performance.

*/

func modifyValue(x int) {
	x = 10 // This change won't affect the original variable
}

func modifyPointer(x *int) {
	*x = 10 // This change will affect the original variable
}

/**
  - *person -> this is a type description - it means we are working with a pointer to a person
  - *pointerToPerson -> this is an operator, meaning we want to manipulate the value the pointer is referencing to

*/

func (pointerToPerson *person) updateName(s string) {
	(*pointerToPerson).firstName = s
}

func (p person) print() {
	// +%v gives us all the fields
	fmt.Printf("%+v\n", p)
}

func main() {
	wasif := person{firstName: "Wasif", lastName: "Siddique"}

	wasif.print()

	sornali := person{
		firstName: "Sornali",
		lastName:  "Rahman",
		contactInfo: contactInfo{
			email: "sornalirahman9@gmail.com",
			zip:   83505,
		},
	}

	sornali.details()

	/**
	* &variable = gives me the memory address of the value this variable is pointing at
	* Turn VALUE into ADDRESS with $variable
	*
	*
	* *pointer = gives me the value this memory address is pointing at
	* Turn ADDRESS into VALUE with *pointer(address)
	*/

	sornaliP := &sornali

	sornaliP.updateName("Shorna")
	sornaliP.details()

	nusaiba := person{
		firstName: "Nusaiba",
		lastName:  "Siddiqua",
		contactInfo: contactInfo{
			email: "siddiquanusaiba31@gmail.com",
			zip:   98301,
		},
	}

	/**
	This works too
	*/
	nusaiba.updateName("Nusu")
	nusaiba.details()

	/**
	- golang gotchas
	- slices can be directly modified without using pointers to address that variable's memory address value
	- Reference Type: Slices in Go are built on top of arrays and include a pointer to the underlying array, along with the length and capacity. When you modify a slice, you are modifying the data in the underlying array.
	- Direct Modification: You can directly modify the elements of a slice without needing to pass a pointer to the slice itself.
	- No Need for Pointers: Since slices are already references to the underlying data, you don't need to use pointers to modify their contents.
	*/
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
