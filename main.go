package main

import (
	"fmt"
	"google-calendar-zoom-autostart/google"
	"log"
)

func main() {
	upcoming, err := google.GetUpcomingMeeting()
	if err != nil {
		log.Fatal("something happend: ", err)
		return
	}

	fmt.Printf("%#v \n", upcoming)
	fmt.Println(upcoming.Title)

}
