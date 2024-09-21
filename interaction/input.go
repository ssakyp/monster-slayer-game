package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialAttackIsAvailable bool) string {
	for {
		playerChoice, _ := getPlayerInput()

		if playerChoice == "1" {
			return "ATTACK"
		} else if playerChoice == "2" {
			return "HEAL"
		} else if playerChoice == "3" && specialAttackIsAvailable {
			return "SPECIAL_ATTACK"
		}
		fmt.Println("Fetching the use input failed. Please try again.")
	}
}

func getPlayerInput() (string, error) {
	fmt.Print("Your choice: ")

	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	userInput = strings.ReplaceAll(userInput, "\n", "")

	return userInput, nil
}
