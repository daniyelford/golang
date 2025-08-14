package main

import (
	"fmt"
	"math/rand"
)

const (
	Rock     = "rock"
	Paper    = "paper"
	Scissors = "scissors"
)

//get players choices function

func getPlayerChoice() string {
	var choice string
	fmt.Println("please  select an option : rock, paper, scissors")
	fmt.Scanln(&choice)
	return choice
}

//get computers choice function

func getComputerChoice() string {

	choices := []string{Rock, Paper, Scissors}
	return choices[rand.Intn(len(choices))]

	/*  if you want as a simple function do this

	func getComputerChoice() string {
	return rock  or return paper or return scissors
	}
	*/
}

func winner(playerChoice, computerChoice string) string {
	if playerChoice == computerChoice {
		return "draw"
	}

	switch playerChoice {

	case Rock:
		if computerChoice == Scissors {
			return "you win   rock>scissors  "
		}
		return "you lose  paper>rock "

	case Paper:
		if computerChoice == Rock {
			return "you win  paper>rock "
		}
		return "you lose  paper<scissors "

	case Scissors:
		if computerChoice == Paper {
			return "you win  scissors>paper "
		}
		return "you lose  rock>scissors "
	default:
		return ("invalid option/choice")
	}
}
func main() {
	playerChoice := getPlayerChoice()
	computerChoice := getComputerChoice()

	fmt.Printf(" your choice : %s\n", playerChoice)
	fmt.Printf("computer choice :%s\n", computerChoice)

	result := winner(playerChoice, computerChoice)
	fmt.Println(result)
}
