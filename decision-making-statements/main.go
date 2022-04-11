package main

import (
	"myapp/game"
)

func main() {

	displayChannel := make(chan string)
	roundChannel := make(chan int)

	game := game.Game{
		DisplayChannel: displayChannel,
		RoundChannel:   roundChannel,
		Round: game.Round{
			RoundNumber:   0,
			PlayerScore:   0,
			ComputerScore: 0,
		},
	}

	go game.Rounds()
	game.ClearScreen()
	game.Intro()

	for {
		// assigm imt value to round channel
		game.RoundChannel <- 1
		// wait for the value to be assigned
		<-game.RoundChannel

		if game.Round.RoundNumber > 3 {
			break
		}

		if !game.PlayRound() {
			game.RoundChannel <- -1
			<-game.RoundChannel
		}

	}

	game.Summary()

}
