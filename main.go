package main

import (
	"fmt"
	"google-calendar-zoom-autostart/google"
	"google-calendar-zoom-autostart/zoom"
)

func main() {
	upcoming, err := google.GetUpcomingMeeting()
	if err != nil {
		fmt.Println("something happend: ", err)
		return
	}

	fmt.Printf("%#v \n", upcoming)
	fmt.Println(upcoming.Title)
	fmt.Println(upcoming.MeetingURL)

	zoom.OpenZoom(upcoming)

}
