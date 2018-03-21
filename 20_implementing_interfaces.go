package main

import "fmt"

/*
    -Go supports the use of interfaces to define high-level abstractions and create programming contracts
    -Unlike other object-oriented languages, you don't explicitly declare a type to be an implementation of
     an interface
    -Instead there is a very simple rule. If a type implements all of the methods defined in an interface,
     then it's an implementation of that interface. There is no keyword "implements" or similar such as in
     Java and C++
    -Interface and type relationships are implied by the presence of the methods

    Interesting fact
    -In fact, every type of Go is an implementation of some interface
    -And whether or not there's a specific interface that it implements with specific set of methods
     every types also is an implementation of a no-methods interface simply named "Interface"
*/

// this defines the animal interface
type Animal interface{
    Speak() string // this defines the Speak method with the return type string
}

// -- Dog

type Dog struct {

}

// once the method structure is exactly the same as interface, it automatically implements
// it from the interface
func (d Dog) Speak() string {
    return "Woof!"
}

// -- Cat

type Cat struct {

}

func (c Cat) Speak() string {
    return "Meow!"
}

// -- Cow

type Cow struct {

}

func (c Cow) Speak() string {
    return "Moo!"
}

// -- MAIN

func main() {
    poodle := Animal(Dog{}) // this proves that the Dog type is an implementation of the Animal interface
    fmt.Println(poodle)

    animals := []Animal{Dog{}, Cat{}, Cow{}}

    // throw away values are represented by the underscore(_) character
    // the throw away values would normally be the index
    // animal will be a reference to the current Animal object
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}