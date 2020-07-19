package lib

import (
	"fmt"
	"net/http"
)

// Redirect redirects the user to some far away land...
func Redirect(res http.ResponseWriter, url string) {
	res.Header().Add("Location", url)
	res.WriteHeader(301)
	res.Write([]byte(fmt.Sprintf("Redirecting to <a href='%[1]s'>%[1]s</a>...", url)))
}
