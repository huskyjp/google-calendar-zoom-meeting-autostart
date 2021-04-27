package google

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"golang.org/x/oauth2"
)

func GetCredidentialGoogleCalendarTokenFromFile() (*oauth2.Token, error) {
	token, err := ioutil.ReadFile(calendarToken) // how to read file from the file

	if err != nil || len(token) == 0 {
		return &oauth2.Token{}, fmt.Errorf("token file empty")
	}

	var authToken oauth2.Token
	err = json.Unmarshal(token, &authToken)
	if err != nil {
		return nil, err
	}

	return &authToken, nil
}

func WriteGoogleCalendarToken(token *oauth2.Token) error {
	tokenMarshaled, err := json.Marshal(token)
	fmt.Println("MARSHAR: ", tokenMarshaled)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(calendarToken, tokenMarshaled, 0644)
}
