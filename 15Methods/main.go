package main

import "fmt"

func main() {
	fmt.Println("Welcome structs")

	divyansh := User{"Divyansh","divyansh@go.dev,",true,20}
	fmt.Println("User is", divyansh)
	fmt.Printf("User is %+v\n", divyansh)
	fmt.Printf("User's name is %+v email is %+v\n",divyansh.Name,divyansh.Email)
	divyansh.GetStatus()
	divyansh.NewMail()
	fmt.Printf("User's name is %+v email is %+v\n",divyansh.Name,divyansh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus()  {
	fmt.Println("Is user active: ",u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.in"
	fmt.Printf("Email of %v is %v\n",u.Name,u.Email)
}