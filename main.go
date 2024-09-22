package main

import (
	"github.com/ssakyp/monster-slayer-game/actions"
	"github.com/ssakyp/monster-slayer-game/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundData{}

func main() {
	startGame()
	// start game
	winner := "" //contains "Player" || "Monster" || ""

	// execute until winner == "Player || "Monster"
	for winner == "" {
		// new round and return wether there is a a winner("Player", "Monster") or not ("")
		winner = executeRound()
	}

	// end game
	endGame(winner)
}

// initila greeting message
func startGame() {
	// call interaction package and print greeting message
	interaction.PrintGreeting()

}

// this will hold all the logic for user letting to choose "actions"
func executeRound() string {

	// variables declaration for round statistics
	var playerAttackDmg int
	var playerHealValue int
	var monsterAttackDmg int

	// in every three round, currentRound is considered as a special round
	currentRound++ // increment the currentRound by 1

	// check if the current round is a special one
	isSpecialRound := currentRound%3 == 0 // 1 / 3 => 1; 2 / 3 => 2; 3 / 3 => 0; 4 / 3 => 1

	// prompt appropriate actions according to round-count
	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	if userChoice == "ATTACK" {
		playerAttackDmg = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		playerAttackDmg = actions.AttackMonster(true)
	}

	// for every round after player attack to monster or heal, monster shoudl attack back
	monsterAttackDmg = actions.AttackPlayer()

	// get player-monster current health to judge winner
	playerHealth, monsterHealth := actions.GetHealthAmounts()

	// pass the values to RoundData func in output
	roundData := interaction.RoundData{
		Action:           userChoice,
		PlayerHealth:     playerHealth,
		MonsterHealth:    monsterHealth,
		PlayerAttackDmg:  playerAttackDmg,
		PlayerHealValue:  playerHealValue,
		MonsterAttackDmg: monsterAttackDmg,
	}

	// this will print the current round stats to console
	interaction.PrintRoundStatistics(&roundData)

	// append new roundData for gameRound slice
	gameRounds = append(gameRounds, roundData)

	// logic to decide winner
	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

// function will declare the winner and write the log file
func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
