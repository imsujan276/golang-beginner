package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Game struct {
	DisplayChannel chan string
	RoundChannel   chan int
	Round          Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

// Rounds receiver for type Game
func (g *Game) Rounds() {
	// use select to process input in channels
	for {
		select {
		case round := <-g.RoundChannel:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChannel <- 1
		case msg := <-g.DisplayChannel:
			fmt.Println(msg)
			// send info back when done
			g.DisplayChannel <- ""
		}
	}
}

// clearScreen clears the screen
func (g *Game) ClearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (g *Game) Intro() {
	g.DisplayChannel <- fmt.Sprintf(`
Rock, Paper, & Scissors
-----------------------
Game is played for three rounds, and best of three wins the game. Good luck!
`)
	<-g.DisplayChannel
}

// PlayRound returns true if round is complete else return false
func (g *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := 1

	g.DisplayChannel <- fmt.Sprintf(`

Round %d
-------------
`, g.Round.RoundNumber)
	<-g.DisplayChannel

	fmt.Print("Please enter rock, paper or scissors\n->")

	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(3)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	g.DisplayChannel <- fmt.Sprintf("\nPlayer choose: %s", strings.ToUpper(playerChoice))
	<-g.DisplayChannel

	switch computerValue {
	case ROCK:
		g.DisplayChannel <- "Computer choose ROCK"
		<-g.DisplayChannel
		break
	case PAPER:
		g.DisplayChannel <- "Computer choose PAPER"
		<-g.DisplayChannel
		break
	case SCISSORS:
		g.DisplayChannel <- "Computer choose SCISSORS"
		<-g.DisplayChannel
		break
	default:
		break
	}

	fmt.Println()
	if playerValue == computerValue {
		g.DisplayChannel <- "Its a DRAW"
		<-g.DisplayChannel
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		default:
			g.DisplayChannel <- "Invalid Choice!"
			<-g.DisplayChannel
			return false
		}
	}

	return true
}

func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChannel <- "Computer wins"
	<-g.DisplayChannel

}

func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChannel <- "Player wins"
	<-g.DisplayChannel
}

func (g *Game) Summary() {
	g.DisplayChannel <- fmt.Sprintf(`

Final Score
-----------
Player: %d/3, Computer: %d/3
`, g.Round.PlayerScore, g.Round.ComputerScore)
	<-g.DisplayChannel

	if g.Round.PlayerScore > g.Round.ComputerScore {
		g.DisplayChannel <- "Player wins the game!"
		<-g.DisplayChannel
	} else {
		g.DisplayChannel <- "Computer wins the game!"
		<-g.DisplayChannel
	}
}
