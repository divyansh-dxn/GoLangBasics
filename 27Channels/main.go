package main

import "fmt"


func main()  {
	fmt.Println("Welcome to channels")

	myCh := make(chan int)
	myCh <- 5
	fmt.Println(<-myCh)
}