package api

import (
	"net/http"
	"os"

	"github.com/cjdenio/slack-frames/lib"
	"github.com/gorilla/sessions"
	"github.com/slack-go/slack"
)

// ExchangeCodeRoute does stuff with the OAuth code
func ExchangeCodeRoute(res http.ResponseWriter, req *http.Request) {
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	session, _ := store.Get(req, "session")

	code := req.URL.Query().Get("code")

	if code == "" {
		res.WriteHeader(400)
		res.Write([]byte("Something went wrong :("))
		return
	}

	resp, err := slack.GetOAuthV2Response(&http.Client{}, os.Getenv("SLACK_CLIENT_ID"), os.Getenv("SLACK_CLIENT_SECRET"), code, "http://localhost:3000/api/code")
	if err != nil {
		res.Header().Add("Content-Type", "text/html")
		res.Write([]byte("Hmm... that didn't work. <a href='/api/login'>Try again?</a>"))
	} else {
		session.Values["token"] = resp.AuthedUser.AccessToken
		session.Values["user"] = resp.AuthedUser.ID
		session.Save(req, res)

		lib.Redirect(res, "/")
	}
}
