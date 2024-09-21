package interactions

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

var reader bufio.NewReader(os.Stdin)

GetUserInput() string {
	fmt.Println("Please choose your action: ")
	fmt.Println("---------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")

	userChoice, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("INVALID INPUT")
		return
	}

	userChoice = strings.Replace(userChoice, "\n", "", -1)
	fmt.Printf("Your choice: %v\n", userChoice)
	return userChoice
}

func GameStartUp() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting a new game ...")
	fmt.Println("Good luck!")
}