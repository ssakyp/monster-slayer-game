package main

import (
	"github.com/ssakyp/monster-slayer-game/actions"
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
		actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		actions.HealPlayer()
	} else {
		actions.AttackMonster(true)
	}

	actions.AttackPlayer()
	playerHealth, monsterHealth := actions.GetHealthAmounts()
	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame() {}
