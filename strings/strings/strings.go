package strings

import (
	"fmt"
	"strings"
)

func Init() {
	fmt.Println()
	name := "Hello World"
	fmt.Println("String:", name)
	fmt.Println()

	fmt.Println("Bytes:")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x\n", name[i])
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("Index\tRune\tString")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("Three ways to concatenate strings")
	fmt.Println()

	h := "Hello, "
	w := "world."

	// join by + - not very efficient
	myString := h + w
	fmt.Println(myString)
	fmt.Println()

	// using fmt - more efficient
	myString = fmt.Sprintf("%s%s", h, w)
	fmt.Println(myString)
	fmt.Println()

	// using string builder - very efficient
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println(sb.String())

	fmt.Println()

	name = "abcdefghijklmnopqrstuvwxyz"
	fmt.Println("Getting a substring")
	fmt.Println(name[0:5]) // staring position, number of characters
	// o/p => abcde
}
