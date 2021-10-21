package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch-case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice value", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice: 1")
	case 2:
		fmt.Println("Dice: 2")
	case 3:
		fmt.Println("Dice: 3")
	case 4:
		fmt.Println("Dice: 4")
	case 5:
		fmt.Println("Dice: 5")
	case 6:
		fmt.Println("Dice: 6")
	default:
		fmt.Println("What was that!")
	}
}
