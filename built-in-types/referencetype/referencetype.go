package referencetype

import (
	"fmt"
	"sort"

	"github.com/eiannone/keyboard"
)

var myInt int

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

// receiver for type Animal
func (a *Animal) Says() {
	fmt.Printf("A %s says %s\n", a.Name, a.Sound)
}
func (a *Animal) HowManylegs() {
	fmt.Printf("A %s has %d legs\n", a.Name, a.NumberOfLegs)
}

/*
 pass only one type to channel. rune- single character used to build a string
 chan are mainly used in go routines
*/
var keyPressChan chan rune

// reference types (pointers, slices, maps, functions, channels)
func ReferenceTypes() {
	// pointerType()
	// slicesType()
	// mapType()
	// functionType()
	channelType()

}

//-------- pointer types --------
func pointerType() {
	myInt = 10
	// pointer - points to location in memory that stores myInt
	myFirstPointer := &myInt
	fmt.Println("myInt is ", myInt)
	fmt.Println("myFirstPointer is", myFirstPointer, "with value of", *myFirstPointer)
	// go to the location and update the content in that memory location
	*myFirstPointer = 15
	fmt.Println("myInt is now", myInt)
	updateValueOfPointer(myFirstPointer)
	fmt.Println("after updateValueOfPointer(), myInt is now", myInt)
}

func updateValueOfPointer(num *int) {
	*num = 25
}

//-------- slice types --------
func slicesType() {
	var animals []string
	animals = append(animals, "dog", "cat", "mouse")
	fmt.Println(animals)

	for index, animal := range animals {
		fmt.Println(index, animal)
	}

	fmt.Println("Element 0 is", animals[0])
	fmt.Println("First two elements are", animals[0:2])
	fmt.Println("The slice is", len(animals), "elements long")
	fmt.Println("Is Sorted animals", sort.StringsAreSorted(animals))

	sort.Strings(animals)
	fmt.Println("Sorted animals", animals)

	animals = deleteFromSlice(animals, 1)
	fmt.Println("Deleted animal in index 1", animals)
}

func deleteFromSlice(a []string, i int) []string {
	//copy last element to index i
	a[i] = a[len(a)-1]
	// erase last ement by giving empty string
	a[len(a)-1] = ""
	// remove last ement from a
	a = a[:len(a)-1]
	return a
}

//-------- map types --------
func mapType() {
	intMap := make(map[string]int)
	intMap["a"] = 1
	intMap["b"] = 2
	intMap["c"] = 3

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	fmt.Println("Deleting 'a':0")
	// deleting element from map
	delete(intMap, "a")
	for key, value := range intMap {
		fmt.Println(key, value)
	}

	// check if the key is in map
	el, ok := intMap["b"]
	if ok {
		fmt.Println(el, "is in intMap")
	} else {
		fmt.Println(el, "is not in intMap")
	}

	fmt.Println("Update value of 'b' to 4")
	// change value of element of the map
	intMap["b"] = 4
	for key, value := range intMap {
		fmt.Println(key, value)
	}
}

//-------- function types --------
func functionType() {
	// z := addTwoNumbers(1, 2)
	// fmt.Println(sumMany(1, 2, 3, 4, 5))
	var dog = Animal{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}
	dog.Says()
	dog.HowManylegs()
	var cat = Animal{
		Name:         "cat",
		Sound:        "meaw",
		NumberOfLegs: 4,
	}
	cat.Says()
	cat.HowManylegs()
}

func addTwoNumbers(a, b int) (sum int) {
	sum = a + b
	return
	// naked return - only used for very short function
}

// variatic function - takes any number of function
// can have only one at the end of the parameter
func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total = total + x
	}
	return total
}

//-------- channel types --------
func channelType() {
	keyPressChan = make(chan rune)

	// go routine - runs the function concurrently (runs in background)
	go listenForKeyPress()

	fmt.Println("Press any key. or q to quit")

	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}
		// send info (char) to channel (keyPressChan)
		keyPressChan <- char
	}
}

func listenForKeyPress() {
	for {
		// store value from channel (keyPressChan) to variable (key)
		key := <-keyPressChan
		fmt.Println("You Pressed", string(key))
	}
}
