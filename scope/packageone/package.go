package packageone

import "fmt"

/*
	if package level variable/function names started with small letters are private
	and can not be sued outside the package.
	If the package level variable/function name is started with capital letter,
	it can be reused from other packages also i.e. global
*/
var privateVar = "I am private"

func notExported() {}

var PublicVar = "I am public (or exported)"

func Exported() {}

var PackageVar = "This is the package level variable in packageone"

func PrintMe(s1, s2 string) {
	fmt.Println(s1, s2, PackageVar)
}
