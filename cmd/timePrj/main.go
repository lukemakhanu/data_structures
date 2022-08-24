package main

import (
	"fmt"
	"time"
)

func main() {

	timeString := "2022-08-24T11:00:00"
	stringToTime, err := time.Parse("2006-01-02T15:04:05", timeString)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("stringToTime", stringToTime)
	updatedTime := stringToTime.Format("2006-01-02 15:04:05")
	fmt.Println("Updated format :", updatedTime)

}
