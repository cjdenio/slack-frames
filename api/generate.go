package api

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/nfnt/resize"

	"net/http"

	"github.com/gorilla/sessions"
	"github.com/slack-go/slack"
)

// Generate generates the magic image
func Generate(res http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("<cries in internal server error>"))
		}
	}()

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

	avatar, err := http.Get(userProfile.ImageOriginal)
	if err != nil {
		log.Fatal(err)
		return
	}

	frame, err := http.Get(req.URL.Query().Get("frame"))
	if err != nil {
		log.Fatal(err)
		return
	}

	profileImg, _, err := image.Decode(avatar.Body)
	if err != nil {
		fmt.Println("profileerr " + err.Error())
	}
	frameImg, _, err := image.Decode(frame.Body)
	if err != nil {
		fmt.Println("frameerr " + err.Error())
	}
	frameImg = resize.Resize(uint(profileImg.Bounds().Dx()), uint(profileImg.Bounds().Dy()), frameImg, resize.Bilinear)

	finalImg := image.NewRGBA(profileImg.Bounds())

	draw.Draw(finalImg, profileImg.Bounds(), profileImg, image.Point{}, draw.Over)
	draw.Draw(finalImg, profileImg.Bounds(), frameImg, image.Point{}, draw.Over)

	r, w := io.Pipe()
	body := multipart.NewWriter(w)

	go func() {
		textW, _ := body.CreateFormField("moo")

		textW.Write([]byte("moo"))

		imageW, _ := body.CreateFormFile("image", "image")

		err = png.Encode(imageW, finalImg)
		if err != nil {
			fmt.Println(err)
		}
		body.WriteField("token", slackToken.(string))

		err = body.Close()
		if err != nil {
			fmt.Println(err)
		}
		w.Close()
	}()

	request, _ := http.NewRequest(http.MethodPost, "https://slack.com/api/users.setPhoto", r)
	request.Header.Add("Content-Type", body.FormDataContentType())
	client := &http.Client{}
	client.Do(request)

	res.Write(nil)
}
