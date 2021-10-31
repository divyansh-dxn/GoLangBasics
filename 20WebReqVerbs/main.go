package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web request verbs!")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformFormRequest()
}

func PerformGetRequest() {
	var responseString strings.Builder

	const myurl = "http://localhost:3000/posts"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code", response.StatusCode)
	fmt.Println("Content length", response.ContentLength)
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println(responseString.String())
	fmt.Println("Bytecount is", byteCount)
	// fmt.Println(string(content))
}

func PerformPostRequest() {
	const url = "http://localhost:3000/posts/"
	requestBody := strings.NewReader(`
		{
			"id": 3,
			"title": "icecream-server",
			"author": "sandwichh"
		}
	`)
	response,error := http.Post(url,"application/json",requestBody)
	if error!=nil {
		panic(error)
	}
	defer response.Body.Close()
	content,_:=ioutil.ReadAll(response.Body)
	fmt.Println("Content",string(content))
}

func PerformFormRequest() {
	const myUrl = "http://localhost:3000/posts/"
	// formdata
	data := url.Values{}
	data.Add("id","4")
	data.Add("title","icecream-parlour")
	data.Add("author","sandwichh")
	response,err := http.PostForm(myUrl,data)
	if(err!=nil) {
		panic(err)
	}
	defer response.Body.Close()
	content,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}