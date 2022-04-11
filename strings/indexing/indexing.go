package indexing

import (
	"fmt"
	"strings"
)

func Init() {
	// basicStringIndexing()
	// stringsPackage()
	stringManipulation()

}

func basicStringIndexing() {
	courseName := "Learn Go for begineers"

	fmt.Println(courseName[0])
	// o/p => 76
	fmt.Println(string(courseName[0]))
	// o/p => L

	for i := 0; i <= 5; i++ {
		fmt.Print(string(courseName[i]))
	}
	// o/p => Learn
	fmt.Println()

	fmt.Println("length of courseName:", len(courseName))
	// o/p => length of courseName: 22

	var mySlice []string
	mySlice = append(mySlice, "a", "b", "c")
	fmt.Println("mySlice has", len(mySlice), "elements")
	// o/p => mySlice has 3 elements
	fmt.Println("Last element of mySlice:", mySlice[len(mySlice)-1])
	// o/p => Last element of mySlice: c
}

func stringsPackage() {
	fmt.Printf(`
---------------
Strings Package	
---------------
`)
	courses := []string{
		"learn Go",
		"learn flutter",
		"learn python",
		"learn java",
	}
	for _, x := range courses {
		if strings.Contains(x, "Go") {
			fmt.Println("Go is found in:", x, "and index:", strings.Index(x, "Go"))
		}
	}
	// o/p => Go is found in: learn Go and index: 6

	newString := "Go is a programming language, Go for it!"
	fmt.Println(strings.HasPrefix(newString, "Go"))
	// o/p => true
	fmt.Println(strings.HasSuffix(newString, "!"))
	// o/p => true
	fmt.Println(strings.Count(newString, "Go"))
	// o/p => 2
	fmt.Println(strings.Count(newString, "Fish"))
	// o/p => 0
	fmt.Println(strings.Index(newString, "Go"))
	// o/p => 0
	fmt.Println(strings.LastIndex(newString, "Go"))
	// o/p => 30
	fmt.Println(strings.LastIndex(newString, "Go00"))
	// o/p => -1
}

func stringManipulation() {
	newString := "Go is a programming language, Go for it!"

	if strings.Contains(newString, "Go") {
		// newString = strings.Replace(newString, "Go", "Golang", -1)
		newString = strings.ReplaceAll(newString, "Go", "Golang")
		// o/p => Golang is a programming language, Golang for it!

		// newString = strings.Replace(newString, "Go", "Golang", 1)
		// o/p => Golang is a programming language, Go for it!
	}
	fmt.Println(newString)

	// string comparision
	a := "A"
	if a == "A" {
		fmt.Println("a is equal to A")
	}
	// o/p => a is equal to A

	// compares the runs of the string
	if "Alpha" > "Absolute" {
		fmt.Println("A greater than B")
	} else {
		fmt.Println("A is not greater than B")
	}
	// o/p => A greater than B

	badEmail := "   Me@here.com    "
	badEmail = strings.TrimSpace(badEmail)
	fmt.Println(strings.ToLower(badEmail))
	//o/p => me@here.com

	fmt.Println()

	str := "alpha alpha alpha alpha alpha alpha"
	fmt.Println(replaceNth(str, "alpha", "beta", 1))
	// o/p => beta alpha alpha alpha alpha alpha
}

func replaceNth(s, toBeReplaced, replaceBy string, position int) string {
	// index
	i := 0

	for j := 1; j <= position; j++ {
		x := strings.Index(s[i:], toBeReplaced)
		if x < 0 {
			//did not find it
			break
		}
		i = i + x
		if j == position {
			return s[:i] + replaceBy + s[i+len(toBeReplaced):]
		}
		i += len(toBeReplaced)
	}
	return s
}
