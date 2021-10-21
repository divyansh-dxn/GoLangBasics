package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	// for loops
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	fmt.Println("")
	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println("")
	for index, day := range days {
		fmt.Printf("index %v day %v\n", index, day)
	}

	// equivalent to while
	i := 0
	for i < 10 {
		if i == 6 {
			goto jumpHere
		}
		if i == 2 {
			i++
			continue
		}
		fmt.Println(i)
		if i == 8 {
			break
		}
		i++
	}

jumpHere:
	fmt.Println("Jumped here")

}
