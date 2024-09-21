package main

import (
	"github.com/ssakyp/monster-slayer-game/interaction"
)

var currentRound = 0

func main() {
	startGame()

	winner := "" // "Player" || "Monster" || ""

	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()

}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0 // 1 / 3 => 1; 2 / 3 => 2; 3 / 3 => 0; 4 / 3 => 1

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	if userChoice == "ATTACK" {

	} else if userChoice == "HEAL" {

	} else {

	}

	return ""
}

func endGame() {}
