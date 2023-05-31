package main

import (
	"log"
	"time"
)

func main() {
	var date = "20230531"
	getDate, err := GetDate(date)
	if err != nil {
		log.Println("error cause:")
		return
	}
	log.Println(getDate)
}

func GetDate(date string) (string, error) {
	strToTime, err := time.Parse("20060102", date)
	if err != nil {
		log.Println("error cause:", err)
		return "", err
	}
	nextDate := strToTime.AddDate(0, 1, 0)
	var newDate = nextDate.Format("20060102")
	execDay := strToTime.Day()
	if execDay != nextDate.Day() {
		strToTime = strToTime.AddDate(0, 2, 0)
		newDate = strToTime.Format("20060102")
		if execDay == strToTime.Day() {
			return newDate, nil
		}

	}
	return newDate, nil
}
