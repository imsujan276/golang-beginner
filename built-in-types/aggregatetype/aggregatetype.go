package aggregatetype

import "fmt"

// array of length 3 of type string
var myStrings [3]string

// struct of type Car
type Car struct {
	NumberOfTyres int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

// aggregate types (array, struct). array rarely used
func AggregareTypes() {
	myStrings[0] = "a"
	myStrings[1] = "b"
	myStrings[2] = "c"
	fmt.Println("First element in array is", myStrings[0])

	myCar := Car{
		NumberOfTyres: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Volvo",
		Model:         "XC",
		Year:          2021,
	}
	fmt.Printf("My car is a %d %s %s\n", myCar.Year, myCar.Make, myCar.Model)
}
