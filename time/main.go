package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Printf("Nigeria time:$%v\n", presentTime)

	fmt.Println(presentTime.Format("01-02-2001 15:04:05 monday:"))

	createdDate := time.Date(2020, time.August, 07, 04, 0, 0, 0, &time.Location{})
	fmt.Println(createdDate)
}go build
