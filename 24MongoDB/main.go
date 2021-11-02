package main

import (
	"fmt"
	"gomongo/routes"
	"net/http"
)

func main() {
	fmt.Println("Welcome to mongodb")
	fmt.Println("SERVER IS STARTING...")
	r := routes.Router()
	http.ListenAndServe(":4000",r)
	fmt.Println("LISTENING AT PORT 4000...")
}
