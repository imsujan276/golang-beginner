package main

import (
	"fmt"
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "D.", Salary: 50000, FullTime: true},
	{FirstName: "Sally", LastName: "J.", Salary: 60000, FullTime: true},
	{FirstName: "Mark", LastName: "S.", Salary: 30000, FullTime: false},
	{FirstName: "Allan", LastName: "A.", Salary: 15000, FullTime: false},
	{FirstName: "Margaret", LastName: "C.", Salary: 100000, FullTime: true},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	// log.Println(myStaff.All())
	// staff.OverpaidLimit = 50000
	log.Println("Overpaid staff:", myStaff.Overpaid())
	fmt.Println()
	log.Println("Underpaid staff:", myStaff.Underpaid())
	fmt.Println()
	// myStaff.notVisible() // unexported methods/variables can not be used in other files

}
