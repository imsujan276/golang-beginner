package main

import "fmt"

type Vehicle struct {
	NumberOfWheels    int
	NumberOfpassenger int
}

type Car struct {
	Year       int
	Make       string
	Model      string
	isElectric bool
	isHybrid   bool
	Vehicle    Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Println("Number of Passengers", v.NumberOfpassenger)
	fmt.Println("Number of Wheels", v.NumberOfWheels)
}

func (c Car) show() {
	fmt.Println("Year:", c.Year)
	fmt.Println("Make:", c.Make)
	fmt.Println("Model:", c.Model)
	fmt.Println("isHybrid:", c.isHybrid)
	fmt.Println("isElectric:", c.isElectric)
	c.Vehicle.showDetails()
}

func main() {
	suv := Vehicle{
		NumberOfWheels:    4,
		NumberOfpassenger: 6,
	}

	volvoXC90 := Car{
		Year:       2021,
		Make:       "Volvo",
		Model:      "XC90 T8",
		isElectric: false,
		isHybrid:   true,
		Vehicle:    suv,
	}
	volvoXC90.show()

	fmt.Println()

	teslaModelX := Car{
		Year:       2021,
		Make:       "Tesla",
		Model:      "Model X",
		isElectric: true,
		isHybrid:   false,
		Vehicle:    suv,
	}
	teslaModelX.Vehicle.NumberOfpassenger = 8
	teslaModelX.show()
}
