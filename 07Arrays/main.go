package main

import "fmt"

func main()  {
	fmt.Println("Welcome arrays!")
	var fruits [5]string

	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[3] = "peach"
	fruits[4] = "mango"

	fmt.Println("Fruits are", fruits)
	fmt.Println("Fruits are", len(fruits))


	var vegitables = [3]string{"potato","tomato","carrot"}
	fmt.Println("Vegitables are: ",vegitables)
	fmt.Println("Length is: ",len(vegitables))

}