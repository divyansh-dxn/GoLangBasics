package main

import "fmt"

func main() {
	fmt.Println("Welcome if-else!")

	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount >= 10 && loginCount <= 20 {
		result = "Suspicious user"
	} else {
		result = "Block him"
	}
	fmt.Println("Analysis result: ", result)



	var num = 57
	if num%2==0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// initialization and checking at same time
	if num :=3; num<10 {
		fmt.Println("Num is less than 10")
	}

}
