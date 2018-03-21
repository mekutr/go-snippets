package main

import "fmt"

/*
    -Go also supports the concept of methods. They are essentially functions but they
     are owned by some structure or other type
    -In a more completely object-oriented language such as Java, you'd say that each method
     is a member of a class
    -In Go, a method is a member of a type
 */

type Dog struct {
    Breed string
    Weight int
    Sound string
}

// -To add a method declare a separate function but add the type of the function's owner
//  before the function name with an identifier wrapped in parentheses
// -When the method is called, it'll receive a reference to the object that owns it
//  that'll be an instance of the Dog structure and it will have access to all of the fields
//  of that structure
// this method receives the copy of the Dog object every time it is called
// --the changes it makes does not reflect any changes on the original object
func (d Dog) Speak() {
    fmt.Println(d.Sound)
}

// this method receives the Dog object as a pointer
// --the changes it makes changes the contents of the object because the object gets passed bby its address
func (d *Dog) SpeakThreeTimes() {
    d.Sound = fmt.Sprintf("%v! %v! %v!", d.Sound, d.Sound, d.Sound)
    fmt.Println(d.Sound)
}

func main() {
    poodle := Dog{"Poodle", 37, "Woof"}
    fmt.Println(poodle)
    poodle.Speak()

    poodle.Sound = "Arf"
    poodle.Speak()

    poodle.SpeakThreeTimes()
    poodle.SpeakThreeTimes()
}
