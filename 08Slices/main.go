package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices!")

	var vegitables = []string{"potato", "tomato", "carrot"}
	fmt.Printf("Type of vegitables is %T\n", vegitables)
	vegitables = append(vegitables, "beans")
	fmt.Println("Fruits are", vegitables)

	// slicing slices
	fmt.Println("1,2", vegitables[1:3])  // included => 1,2
	fmt.Println("0,1,2", vegitables[:3]) // included => 1,2

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 458
	highScores[3] = 856
	// highScores[4] = 233 index out of range, but
	highScores = append(highScores, 255,778,633) // is valid
	fmt.Println("Appended scores",highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("Sorted scores",highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	var courses = []string{"reactjs","javascript","swift","python","django","java","springboot"}
	fmt.Println("courses are",courses)
	var index int = 2
	courses = append(courses[:index],courses[index+1:]... )
	fmt.Println("new courses",courses)
}
