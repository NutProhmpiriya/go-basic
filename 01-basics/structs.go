// This file demonstrates the usage of structs in Go
// Structs are user-defined types that group related data together
// They can have methods associated with them, similar to classes in other languages

package main

import "fmt"

// Person represents a person with basic information
// Struct fields are typically capitalized to make them exported (public)
type Person struct {
	Name    string   // Person's full name
	Age     int      // Person's age in years
	Address Address  // Nested struct for address information
}

// Address represents a physical address
// This demonstrates struct composition (nested structs)
type Address struct {
	Street  string  // Street name and number
	City    string  // City name
	Country string  // Country name
}

// Method for Person struct
// This is a value receiver method - it works with a copy of the Person
// Syntax: func (receiverName ReceiverType) MethodName() ReturnType
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I'm %d years old", p.Name, p.Age)
}

// Method with pointer receiver to modify the struct
// This is a pointer receiver method - it can modify the original Person
// Use pointer receivers when:
// 1. You need to modify the receiver
// 2. The struct is large and you want to avoid copying
// 3. You want consistency with other methods of the type
func (p *Person) Birthday() {
	p.Age++
}

func main() {
	// Creating a nested struct (Address)
	// Method 1: Create struct with field names
	address := Address{
		Street:  "123 Main St",
		City:    "Bangkok",
		Country: "Thailand",
	}

	// Creating the main struct (Person)
	// Method 2: Create struct with field names and nested struct
	person := Person{
		Name:    "John",
		Age:     25,
		Address: address,
	}

	// Accessing struct fields using dot notation
	// You can access nested struct fields by chaining dots
	fmt.Printf("Person: %+v\n", person)         // %+v prints field names and values
	fmt.Printf("Address: %+v\n", person.Address)

	// Using struct methods
	// Call value receiver method
	fmt.Println(person.Greet())

	// Using pointer receiver method
	// Go automatically handles the pointer conversion
	person.Birthday()
	fmt.Printf("After birthday: %d years old\n", person.Age)

	// Anonymous struct
	// Useful for one-off struct definitions
	// Common in test cases or when you need a temporary structure
	employee := struct {
		ID     int
		Role   string
		Active bool
	}{
		ID:     1,
		Role:   "Developer",
		Active: true,
	}

	fmt.Printf("Employee: %+v\n", employee)
}
