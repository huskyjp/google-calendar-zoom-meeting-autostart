package zoom

import (
	"fmt"
	"google-calendar-zoom-autostart/google"
	"log"
	"os/exec"
)

func OpenZoom(zoomInfo google.GoogleCalendarMeeting) {

	//&& zoomInfo.MeetingURL != "" && zoomInfo.ConferencePassword != ""
	if zoomInfo.ConferenceID != "" {
		zoom := fmt.Sprintf("zoommtg://zoom.us/join?action=join&confno=%s", zoomInfo.ConferenceID)
		if zoomInfo.ConferencePassword != "" {
			zoom = fmt.Sprintf(zoom+"&"+"pwd=%s", zoomInfo.ConferencePassword)
		}
		// zoom := strings.Replace(zoomInfo.MeetingURL, "https", "zoommtg", 1) + zoomInfo.ConferenceID + zoomInfo.ConferencePassword
		// zoom = zoom + zoomInfo.ConferenceID + zoomInfo.ConferencePassword

		// TODO: Explore how to run command in windows
		err := exec.Command("cmd", "start", zoom)
		if err != nil {
			fmt.Println(err)
			fmt.Println(zoom)

			return
		}
	} else {
		log.Fatalln("we could not open the zoom link or there is no verified meeting... see you again!")
	}
}
