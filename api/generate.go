package api

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"net/http"

	"github.com/gorilla/sessions"
	"github.com/nfnt/resize"
	"github.com/slack-go/slack"
)

// Generate generates the magic image
func Generate(res http.ResponseWriter, req *http.Request) {
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	session, _ := store.Get(req, "session")
	slackToken := session.Values["token"]

	if slackToken == nil {
		http.Error(res, "error", 400)
		return
	}

	slackClient := slack.New(slackToken.(string))

	userProfile, err := slackClient.GetUserProfile("", false)
	if err != nil {
		log.Fatal(err)
		return
	}
	res.Write([]byte(userProfile.ImageOriginal))
	return

	profile, _ := os.Open("profile.JPG")
	frame, _ := os.Open("frame.png")

	profileImg, _ := jpeg.Decode(profile)
	frameImg, _ := png.Decode(frame)
	frameImg = resize.Resize(uint(profileImg.Bounds().Dx()), uint(profileImg.Bounds().Dy()), frameImg, resize.Bilinear)

	finalImg := image.NewRGBA(profileImg.Bounds())

	draw.Draw(finalImg, profileImg.Bounds(), profileImg, image.Point{}, draw.Over)
	draw.Draw(finalImg, profileImg.Bounds(), frameImg, image.Point{}, draw.Over)

	final, _ := os.Create("final.png")

	png.Encode(final, resize.Resize(512, 512, finalImg, resize.Bilinear))
}
