package main

import "fmt"

var PublicVariable = 48 // public variables start with CAPS
const ConstantVariable = "NEW CONSTANT VALUE"

func main() {

	fmt.Println(PublicVariable)
	// ConstantVariable = 50 constant cannot be reassigned
	fmt.Println(ConstantVariable)

	age := 20 // local variables can be declared without var, with :=(walrus) operator
	fmt.Println(age)

	var name string = "icecreamsandwichh"
	fmt.Println(name)
	fmt.Printf("type of name is %T\n", name)  // printf function like C

	// name = 20  string type cannot be assigned to int

	var integer8  int8 = 25  // 8 bit integer 8,32,64 are available, plain int also can be used
	// integer8 = 25556 cannot assign a number > 8-bit
	fmt.Println(integer8)
}