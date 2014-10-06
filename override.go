package main

import (
    "fmt"
)

type Animal struct {
    name string
}

func (it *Animal) think() string {
    return "The world is good."
}

func (it *Animal) Say() {
    fmt.Printf("%s thinks, '%s'.\n", it.name, it.think())
}

type Dog struct {
    Animal
}

func NewDog(name string) *Dog {
    return &Dog{Animal: Animal{name}}
}

func (d *Dog) think() string {
    return "Bones are really good."
}

// $ go run override.go
// Snoopy thinks, 'The world is good.'.
// Bones are really good.
func main() {
    dog := NewDog("Snoopy")
    dog.Say() // think() is not virtual so dog will not say "Bones are really good."
    fmt.Println(dog.think())
}
