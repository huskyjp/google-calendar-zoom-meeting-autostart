package google

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type GoogleCalendarMeeting struct {
	Title      string
	StartTime  time.Time
	MeetingURL string
}

const (
	calendarToken = "/google-calendar-token"
	credentials   = "credentials.json"
)

func GetUpcomingMeeting() (GoogleCalendarMeeting, error) {
	google := GoogleCalendarMeeting{}
	client, err := newGoogleCalendarClient()
	if err != nil {
		return google, err
	}
	eventService := calendar.NewEventsService(client)
	eventListner := eventService.List("primary")
	//eventListner.MaxResults(5)
	//eventListner.SingleEvents(true)
	// eventListner.OrderBy("starttime")

	eventListner.TimeMin(time.Now().Format(time.RFC3339))
	eventListner.TimeMax(time.Now().Add(time.Minute * 30).Format(time.RFC3339))

	events, err := eventListner.Do()
	if err != nil {
		fmt.Println("error retrieving events: ", err)
		return google, err
	}
	fmt.Println("current events inside: ", events)
	if len(events.Items) == 0 {
		fmt.Println("there is no meetings is scheduled at this moment")
		return google, fmt.Errorf("the event length was 0")
	}

	fmt.Println("current event lists: ", events)
	eventOne := events.Items[0]
	//fmt.Printf("%#v\n", eventOne)
	//eventTwo := events.Items[1]
	//fmt.Println(eventOne.Summary)
	//fmt.Println(eventTwo.Summary)

	// get meeting title
	google.Title = eventOne.Summary

	meetingTime, err := time.Parse(time.RFC3339, eventOne.Start.DateTime)
	if err != nil {
		return google, nil
	}
	// get meeting start time
	google.StartTime = meetingTime

	if eventOne.ConferenceData == nil {
		return google, nil
	}

	for _, entry := range eventOne.ConferenceData.EntryPoints {
		if entry.EntryPointType == "video" {
			fmt.Println("the video url is: ", entry.Uri)
			google.MeetingURL = entry.Uri
			break
		}

	}

	return google, err

}

// client for google calender to fetch its data
func newGoogleCalendarClient() (*calendar.Service, error) {

	newCredential, err := ioutil.ReadFile(credentials)
	if err != nil {
		return nil, err
	}

	// parse json to *0auth2.Config type
	newConfig, err := google.ConfigFromJSON(newCredential, calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, err
	}

	token, err := getGoogleCalendarToken(newConfig)
	if err != nil {
		return nil, err
	}

	// get access to the Calendar API
	ctx := context.Background()
	return calendar.NewService(ctx, option.WithTokenSource(newConfig.TokenSource(ctx, token)))

}

// cli handler
func fetchNewToken(config *oauth2.Config) (*oauth2.Token, error) {
	// get authURL: provide authentication page URL
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the authentication code: %s\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return &oauth2.Token{}, fmt.Errorf("Unable to read authorization ode: &v\n", err)
	}

	token, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return &oauth2.Token{}, fmt.Errorf("unable to exchange authCode for token: %s", err)
	}
	// token is type of *oauth2.Token
	return token, err
}

func getGoogleCalendarToken(config *oauth2.Config) (*oauth2.Token, error) {
	token, err := GetCredidentialGoogleCalendarTokenFromFile()
	if err != nil {
		fmt.Printf("Error getting stashed token: %s \n", err)

		token, err := fetchNewToken(config) // parse config and get new token
		fmt.Println("current   ", token)
		if err != nil {
			return token, err
		}
		WriteGoogleCalendarToken(token) // only returns error but creates new file
		return token, nil
	}
	return token, nil
}
