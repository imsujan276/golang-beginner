package main

import (
	"errors"
	"fmt"
	"math"
)

/*
	Should use Same type of Float number for mathematical operations
	rather than int or mixing them
*/

func main() {
	// basicOperators()
	shortCircuitEvaluation()
}

func shortCircuitEvaluation() {
	a := 12
	b := 0
	// if b != 0 && divideTwoNumbers(a, b) == 2 {
	// 	fmt.Println("found 2")
	// }

	c, err := divideTwoNumbers(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}

}

func divideTwoNumbers(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("can not divide by zero")
	}
	return x / y, nil
}

func basicOperators() {
	// multiplication
	radius := 12.0
	area := math.Pi * radius * radius
	fmt.Println("Area:", area)
	// o/p => 452.3893421169302

	// integer division
	half := 1 / 2
	fmt.Println("Half with Integer Division:", half)
	// o/p => 0

	// float division
	halfFloat := 1.0 / 2.0
	fmt.Println("HalfFloat with Float Division:", halfFloat)
	// o/p => 0.5

	// power
	squareOf5 := math.Pow(5.0, 2.0)
	fmt.Println("squareOf5:", squareOf5)
	// o/p => 25

	// modulus
	remainder := 5 % 3
	fmt.Println("remainder:", remainder)
	// o/p => 2

	// unary operator
	x := 3
	x++
	fmt.Println("x++ :", x)
	// o/p => 4
	x--
	x--
	fmt.Println("x is now :", x)
	// o/p => 2
}
