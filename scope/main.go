package main

import (
	"myapp/packageone"
)

var myVar = "This is package level variable"

func main() {
	var blockVar = "This is block level variable"

	packageone.PrintMe(myVar, blockVar)

}
