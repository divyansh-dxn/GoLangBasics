package main

import "fmt"

func sayHello() {
	fmt.Println("Hello programmer")
}

func add(a int, b int) int {
	return a+b
}

func proadder(values ...int) (int,string) {
	sum :=0
	for _,val := range values {
		sum+=val
	}
	return sum,"hi from pro-adder"
}

func main()  {
	fmt.Println("Welcome to functions")
	sayHello()
	fmt.Println("Addition 3 and 5 is", add(3,5))
	sum,_ := proadder(3,5,8,7,2)
	fmt.Println("Addition 3,5,8,7,and 2 is", sum)
}


// lambdas and un-named functions also exist