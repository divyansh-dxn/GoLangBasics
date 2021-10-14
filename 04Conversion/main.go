package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Welcome to our pizza app!")
	fmt.Println("Please rate out pizza on scale of 1-5")

	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input),64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Thanks for rating ",numrating+1)
	}
 }