package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

// user defined TYPE User
type User struct {
	name      string
	age       int
	favNumber float64
	ownsADog  bool
}

func main() {
	reader = bufio.NewReader((os.Stdin))

	var user User

	user.name = readString("What is your name?")
	user.age = readInt("How old are you?")
	user.favNumber = readFloat("What is your favourite number?")
	user.ownsADog = readBool("Do you own a dog? (y/n)")

	// not recommended. slower, uses more memory
	// fmt.Println("Your name is "+name+". You are", age, "years old")

	// using string interpolation
	// fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old", name, age))
	// faster, uses less memory
	fmt.Printf(
		"Your name is %s. You are %d years old. Your favourite number is %.2f. Owns a dog: %t \n",
		user.name,
		user.age,
		user.favNumber, user.ownsADog)

}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}
}

func readBool(s string) bool {
	// for {
	// 	fmt.Println(s)
	// 	prompt()
	// 	userInput, _ := reader.ReadString('\n')
	// 	userInput = strings.Replace(userInput, "\n", "", -1)
	// 	bool, err := strconv.ParseBool(userInput)
	// 	if err != nil {
	// 		fmt.Println("Please enter (y/n)")
	// 	} else {
	// 		return bool
	// 	}
	// }

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)
		prompt()
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Println("Please type y or n")
		} else if char == 'n' || char == 'N' {
			return false
		} else if char == 'y' || char == 'Y' {
			return true
		}
	}
}
