package lib

import (
	"io"
	"net/http"

	"github.com/slack-go/slack"
)

// GetUserAvatar returns a Reader representing the user's avatar.
func GetUserAvatar(token, user string) (io.Reader, error) {
	client := slack.New(token)

	profile, err := client.GetUserProfile(user, false)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(profile.Image192)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
