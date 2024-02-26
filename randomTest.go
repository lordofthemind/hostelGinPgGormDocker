package main

import "fmt"

// Animal represents a basic animal.
type Animal struct {
	Name string
}

// Speak is a method of the Animal type.
func (a *Animal) Speak() {
	fmt.Println("Generic animal sound")
}

// Dog represents a dog, which embeds the Animal type.
type Dog struct {
	Animal
	Breed string
}

// Speak is overridden for the Dog type.
func (d *Dog) Speak() {
	fmt.Println("Woof!")
}

func main() {
	// Create an instance of the Dog type.
	myDog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Labrador",
	}

	// Call the Speak method for the Dog.
	myDog.Speak()

	// Call the Speak method for the embedded Animal.
	myDog.Animal.Speak()
}
