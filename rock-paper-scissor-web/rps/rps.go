package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER    = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS = 2 // beats paper. (paper + 1) % 3 = 2
)

type Round struct {
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

var winMessages = []string{
	"Good job!",
	"Nice work!",
	"You should buy a lottery ticket.",
}
var loseMessages = []string{
	"Too bad!",
	"Try again!",
	"This is just not your day.",
}
var drawMessages = []string{
	"Great minds think alike.",
	"Uh oh. Try again!",
	"Nobody wins, but you can try again.",
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""

	switch computerValue {
	case ROCK:
		computerChoice = "Computer choose ROCK"
		break
	case PAPER:
		computerChoice = "Computer choose PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer choose SCISSORS"
		break
	default:
		break
	}

	messageInt := rand.Intn(3)
	message := ""

	if playerValue == computerValue {
		roundResult = "Its a DRAW"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins"
		message = winMessages[messageInt]
	} else {
		roundResult = "Computer wins"
		message = loseMessages[messageInt]
	}

	result := Round{
		Message:        message,
		ComputerChoice: computerChoice,
		RoundResult:    roundResult,
	}
	return result
}
