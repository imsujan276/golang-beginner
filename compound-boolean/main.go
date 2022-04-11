package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

func main() {
	jack := Employee{
		Name:     "jack S.",
		Age:      27,
		Salary:   30000,
		FullTime: false,
	}

	jill := Employee{
		Name:     "Jill J.",
		Age:      33,
		Salary:   60000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, jack, jill)

	for _, e := range employees {
		if e.Age > 30 {
			fmt.Println(e.Name, "is 30 or older")
		} else {
			fmt.Println(e.Name, "is under 30")
		}

		if e.Age > 30 && e.Salary > 50000 {
			fmt.Println(e.Name, "makes more than 50000 AND is 30 or older")
		} else {
			fmt.Println("Either", e.Name, "makes more than 50000 or is under 30 ")
		}

		if e.Age > 30 || e.Salary > 50000 {
			fmt.Println(e.Name, "makes more than 50000 OR is 30 or older")
		} else {
			fmt.Println(e.Name, "makes less than 50000 and is under 30 ")
		}

		if (e.Age > 30 || e.Salary < 50000) && e.FullTime {
			fmt.Println(e.Name, "matches our unclear criteria")
		}

		println("--------------------------")
	}

}
