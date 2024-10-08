package interaction

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
)

// roundData struct for round statistics
type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func PrintGreeting() {
	monsterFigure := figure.NewFigure("MONSTER SLAYER", "", true)
	fmt.Println(monsterFigure)
	fmt.Println("Starting a new game...")
	fmt.Println("Good luck!")
}

func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("Please chose your action")
	fmt.Println("------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")

	// only displays in a special round
	if specialAttackIsAvailable {
		fmt.Println("(3) Special Attack")
	}
}

// function to print RoundData struct
func PrintRoundStatistics(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a strong attack against monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else {
		fmt.Printf("Player healed for %v. \n", roundData.PlayerHealValue)
	}

	fmt.Printf("Monster attacked player for %v damage.\n", roundData.MonsterAttackDmg)
	fmt.Printf("Player Health: %v.\n", roundData.PlayerHealth)
	fmt.Printf("Monster Health: %v.\n", roundData.MonsterHealth)
}

func DeclareWinner(winner string) {
	fmt.Println("-------------------------")
	newFigure := figure.NewColorFigure("GAME OVER!", "", "red", true)
	fmt.Println(newFigure)
	fmt.Println("-------------------------")
	fmt.Printf("%v won!\n", winner)
}

func WriteLogFile(rounds *[]RoundData) {
	// exePath contain the current directory that exe resides also with the name.exe part
	exPath, err := os.Executable()
	if err != nil {
		fmt.Println("Writing log file failed. Exiting!")
		return
	}

	// this will get rid of the name.exe path and take the rest of other parts (absolute location)
	exPath = filepath.Dir(exPath)

	file, err := os.Create(exPath + "/gamelog.txt")
	// file, err := os.Create("gamelog.txt") // => go run .
	if err != nil {
		fmt.Println("Saving a log file failed. Exiting.")
		return
	}

	for index, value := range *rounds {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(index + 1),
			"Action":                value.Action,
			"Player Attack Damage":  fmt.Sprint(value.PlayerAttackDmg),
			"Player Heal Value":     fmt.Sprint(value.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAttackDmg),
			"Player Health":         fmt.Sprint(value.PlayerHealth),
			"Monster Health":        fmt.Sprint(value.MonsterHealth),
		}
		logLine := fmt.Sprintln(logEntry)
		_, err = file.WriteString(logLine)
		if err != nil {
			fmt.Println("Writing into log file faield. Exiting.")
			continue
		}
	}

	file.Close()
	fmt.Println("Wrote data to log!")
}
