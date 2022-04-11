package basictype

import "log"

// basic types (number, strings, booleans)
var myInt int
var myUInt uint // hols +ve value and 0 (unsigned integer)
var myFloat32 float32
var myFloat64 float64 // can habdle really large numbers
var myString string
var myBool bool

func BasicTypes() {
	myInt = -10
	myUInt = 10
	myFloat32 = 10.1
	myFloat64 = 100.8237
	log.Println(myInt, myUInt, myFloat32, myFloat64)

	// strings are immutable
	myString = "hello"
	log.Println(myString)

	myBool = true
	log.Println(myBool)
}
