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

var winPhrases = []string{
	"You lucky!",
	"Well done",
	"Great!",
}

var drawPhrases = []string{
	"its a draw!",
	"no one wins",
	"mm both chose the same!",
}

var losePhrases = []string{
	"try again!",
	"computer wins",
	"well, you are not good at this game",
}

type Round struct {
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
		break
	case PAPER:
		computerChoice = "Computer chose PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
		break
	default:
	}

	rand := rand.Intn(3)
	message := ""

	if playerValue == computerValue {
		roundResult = "Its a draw"
		message = drawPhrases[rand]

	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins"
		message = winPhrases[rand]
	} else {
		roundResult = "Computer wins"
		message = losePhrases[rand]
	}

	var result Round
	result.Message = message
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	return result
}
