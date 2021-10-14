package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Welcome to time")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("15:04:05 01-02-2006 Monday"))

	createdDate := time.Date(2020,time.August,10,23,23,0,0,time.Local)
	fmt.Println(createdDate)
}