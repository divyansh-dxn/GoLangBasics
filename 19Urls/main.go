package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paumentid=vnvnfe5i6ud"

func main() {
	fmt.Println("Welcome to url")
	fmt.Println(myurl)

	// parsing the url
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Hostname())
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	queryParams := result.Query()
	fmt.Println(queryParams)

	customUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "courses",
		RawPath: "user=divyansh",
	}
	fmt.Println(customUrl.String())
}
