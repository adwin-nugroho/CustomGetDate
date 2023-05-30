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
	var count = 1
	strToTime, err := time.Parse("20060102", date)
	if err != nil {
		log.Println("error cause:", err)
		return "", err
	}
	strToTime = strToTime.AddDate(0, +count, 0)
	var newDate = strToTime.Format("20060102")
	// log.Println("new date b4 looping:", newDate)

	if date[6:8] != newDate[6:8] {
		for {
			count++
			strToTime, err = time.Parse("20060102", date)
			if err != nil {
				log.Println("error cause:", err)
				return "", err
			}
			strToTime = strToTime.AddDate(0, +count, 0)
			newDate = strToTime.Format("20060102")
			// log.Println("new date after looping:", newDate)
			if date[6:8] == newDate[6:8] {
				break
			}
		}
	}

	// log.Println("check last 2 digit:", newDate[6:8])

	return newDate, nil
}
