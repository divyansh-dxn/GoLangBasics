package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"plaform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")
	// lcoCourses := []course{
	// 	{"ReactJS", 299, "LCO", "abc123", []string{"web-dev", "frontend", "js"}},
	// 	{"Mern bootcamp", 199, "LCO", "bcd123", []string{"full-stack", "js"}},
	// 	{"Angular", 299, "LCO", "qwr123", []string{"google", "web-dev js"}},
	// 	{"Angular", 299, "LCO", "qwr123", []string{}},
	// }
	// EncodeJson(lcoCourses)
	jsonData := []byte(`	
	{
		"coursename": "ReactJS",
		"price": 299,
		"plaform": "LCO",
		"tags": [
				"web-dev",
				"frontend",
				"js"
		]
	}`)
	DecodeJson(jsonData)
}

func EncodeJson(courses []course) {
	finalJson, error := json.MarshalIndent(courses, "", "\t")
	if error != nil {
		panic(error)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson(data []byte) {
	isDataValid := json.Valid([]byte(data))
	var lcoCourse course
	if isDataValid {
		fmt.Println("Data is valid")
		json.Unmarshal(data,&lcoCourse)
		fmt.Printf("%#v\n",lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// convert data to map
	var dataMap map[string]interface{}
	json.Unmarshal(data,&dataMap)
	fmt.Printf("%#v\n",dataMap)

}
