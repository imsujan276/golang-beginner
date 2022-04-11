package interfacetype

import (
	"fmt"
)

/*
 interface of type Animal
 list all of the function that the insterface must have in order to be that interface
 for the struct to satisfy an interface, if must receive the the functions that the interface has
*/
type Animal interface {
	Says() string
	HowManylegs() int
}

// struct of type Dog
type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManylegs() int {
	return d.NumberOfLegs
}

// struct of type Cat
type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func (c *Cat) Says() string {
	return c.Sound
}

func (c *Cat) HowManylegs() int {
	return c.NumberOfLegs
}

// interface type
func InterfaceTypes() {
	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}
	riddle(&dog)

	cat := Cat{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
		HasTail:      true,
	}
	riddle(&cat)
}

func riddle(a Animal) {
	r := fmt.Sprintf("This animal says %s and has %d legs. What animal is it?", a.Says(), a.HowManylegs())
	fmt.Println(r)
}
