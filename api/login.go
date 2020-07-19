package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/cjdenio/slack-frames/lib"
)

// LoginRoute redirects the user to Slack's login page.
func LoginRoute(res http.ResponseWriter, req *http.Request) {
	var redirectURI string

	if strings.Contains(req.Host, "localhost") {
		redirectURI = "http://localhost:3000/api/code"
	} else {
		redirectURI = "https://slack-frames.vercel.app/api/code"
	}

	lib.Redirect(res, fmt.Sprintf("https://slack.com/oauth/v2/authorize?user_scope=%s&client_id=%s&redirect_uri=%s", "users.profile:write,users.profile:read", os.Getenv("SLACK_CLIENT_ID"), redirectURI))
}
