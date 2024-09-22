package actions

import (
	"math/rand"
	"time"
)

// this contains the source value to generate a random number
var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)

// player-monster health values
var currentMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) int {
	// min-max attack values for general attack
	minAttackValue := PLAYER_ATTACK_MIN_DMG
	maxAttackValue := PLAYER_ATTACK_MAX_DMG

	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN_DMG
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_DMG
	}

	// calculate damage value
	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentMonsterHealth -= dmgValue // deduct damage value from current monster health
	return dmgValue
}

func HealPlayer() int {
	// calculate heal value
	healValue := generateRandBetween(PLAYER_HEAL_MIN_VALUE, PLAYER_HEAL_MAX_VALUE)

	// system won't allow playerHealth to be greater than 100
	healDiff := PLAYER_HEALTH - currentPlayerHealth

	if healDiff >= healValue {
		currentPlayerHealth += healValue
		return healValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
		return healDiff
	}
}

func AttackPlayer() int {
	minAttackValue := MONSTER_ATTACK_MIN_DMG
	maxAttackValue := MONSTER_ATTACK_MAX_DMG

	// calculate attack player value
	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)

	// update current player health
	currentPlayerHealth -= dmgValue
	return dmgValue
}

func generateRandBetween(min, max int) int {
	return randGenerator.Intn(max-min) + min
}

func GetHealthAmounts() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}
