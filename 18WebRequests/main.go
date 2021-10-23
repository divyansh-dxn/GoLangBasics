package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("lco web request")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Printf("type %T\n", res)
	// fmt.Println(res.Body)
	databytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
