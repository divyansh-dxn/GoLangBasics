package main

import "fmt"

func main()  {
	fmt.Println("Welcome maps!")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["KT"] = "Kotlin"

	fmt.Println("All languages",languages)
	fmt.Println("My favourite language",languages["KT"])

	delete(languages,"RB")
	fmt.Println("Nice languages",languages)

	// loops
	for key,val := range languages {
		fmt.Printf("For key %v language is %v\n",key,val)
	}

}