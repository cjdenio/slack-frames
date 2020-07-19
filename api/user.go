package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/slack-go/slack"

	"github.com/gorilla/sessions"
)

// AuthedRoute tells the frontend whether or not the user is authed.
func AuthedRoute(res http.ResponseWriter, req *http.Request) {
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	session, _ := store.Get(req, "session")

	token := session.Values["token"]

	if token == nil {
		res.WriteHeader(400)
		res.Write([]byte("yer not authed, harry"))
		return
	}

	client := slack.New(token.(string))
	profile, err := client.GetUserProfile(session.Values["user"].(string), false)

	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte("uh oh, there be a 500 error"))
		return
	}

	marshalledJson, _ := json.Marshal(profile)

	res.Header().Add("Content-Type", "application/json")
	res.Write(marshalledJson)
}
