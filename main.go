package main

import (
	"fmt"
	"google-calendar-zoom-autostart/google"
)

func main() {
	upcoming, err := google.GetUpcomingMeeting()
	if err != nil {
		fmt.Println("something happend: ", err)
		return
	}

	fmt.Printf("%#v \n", upcoming)
	fmt.Println(upcoming.Title)

}
