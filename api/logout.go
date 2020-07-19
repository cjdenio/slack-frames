package api

import (
	"net/http"
	"os"

	"github.com/cjdenio/slack-frames/lib"
	"github.com/gorilla/sessions"
)

func LogoutRoute(res http.ResponseWriter, req *http.Request) {
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	session, _ := store.Get(req, "session")

	delete(session.Values, "token")
	delete(session.Values, "user")

	session.Save(req, res)

	lib.Redirect(res, "/")
}
