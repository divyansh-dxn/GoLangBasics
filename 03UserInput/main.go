package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome := "Hello welcome new user"
	fmt.Println(welcome)
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza")

	// comma ok syntax
	input, err := reader.ReadString('\n')
	fmt.Print("Given input:",input)
	fmt.Println(err)

}