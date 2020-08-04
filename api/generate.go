package api

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
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

	fmt.Println(userProfile.ImageOriginal)

	res.Header().Add("Content-type", "image/png")

	//io.Copy(res, frame.Body)

	profileImg, err := jpeg.Decode(avatar.Body)
	if err != nil {
		fmt.Println("profileerr " + err.Error())
	}
	frameImg, err := png.Decode(frame.Body)
	if err != nil {
		fmt.Println("frameerr " + err.Error())
	}
	frameImg = resize.Resize(uint(profileImg.Bounds().Dx()), uint(profileImg.Bounds().Dy()), frameImg, resize.Bilinear)

	finalImg := image.NewRGBA(profileImg.Bounds())

	draw.Draw(finalImg, profileImg.Bounds(), profileImg, image.Point{}, draw.Over)
	draw.Draw(finalImg, profileImg.Bounds(), frameImg, image.Point{}, draw.Over)

	res.Header().Add("Content-type", "image/png")

	r, w := io.Pipe()
	body := multipart.NewWriter(w)
	textW, _ := body.CreateFormField("moo")

	textW.Write([]byte("moo"))

	//imageW, _ := body.CreateFormFile("image", "image")

	//err = png.Encode(imageW, finalImg)
	err = body.Close()
	if err != nil {
		fmt.Println(err)
	}
	w.Close()

	read, _ := ioutil.ReadAll(r)
	fmt.Println(string(read))

	// resp, err := http.Post("https://slack.com/api/users.setPhoto", "multipart/form-data", r)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// stuff, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println(string(stuff))

	res.Write(nil)
}
