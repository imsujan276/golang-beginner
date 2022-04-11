package staff

import "fmt"

// exported - starts by Capital letters
var OverpaidLimit = 75000

// not exported - starts by small letters
var underpaidLimit = 20000

// exported - starts by Capital letters
type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

// receiver function for office type
func (o *Office) All() []Employee {
	return o.AllStaff
}

func (o *Office) Overpaid() []Employee {
	var overpaid []Employee

	for _, e := range o.AllStaff {
		if e.Salary > OverpaidLimit {
			overpaid = append(overpaid, e)
		}
	}
	return overpaid
}

func (o *Office) Underpaid() []Employee {
	var underpaid []Employee

	for _, e := range o.AllStaff {
		if e.Salary <= underpaidLimit {
			underpaid = append(underpaid, e)
		}
	}
	return underpaid
}

func (o *Office) notVisible() {
	fmt.Println("not exported function")
}

/*
	package level function
	can be used inside same package
	but can not be used it outside or as an receiver
*/
func myFunction() {
	fmt.Println("I am a function")
}
