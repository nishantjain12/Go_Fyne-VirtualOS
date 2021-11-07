package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func UnmarshalLanguage(data []byte) (Language, error) {
	var r Language
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Language) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Language struct {
	Data Data `json:"data"`
}

type Data struct {
	Detections [][]Detection `json:"detections"`
}

type Detection struct {
	Language   string `json:"language"`
	IsReliable bool   `json:"isReliable"`
	Confidence int64  `json:"confidence"`
}

type language struct {
	name string
	code string
}

func showLanguageDetector(w fyne.Window) {
	languages := []language{
		{name: "Afrikaans", code: "af"},
		{name: "Albanian ", code: "sq"},
		{name: "Amharic ", code: "am"},
		{name: "Arabic ", code: "ar"},
		{name: "Armenian ", code: "hy"},
		{name: "Azerbaijani ", code: "az"},
		{name: "Basque ", code: "eu"},
		{name: "Belarusian ", code: "be"},
		{name: "Bengali", code: "bn"},
		{name: "Bosnian", code: "bs"},
		{name: "Bulgarian", code: "bg"},
		{name: "Catalan", code: "ca"},
		{name: "Croatian", code: "hr"},
		{name: "Corsican", code: "co"},
		{name: "Czech", code: "cs"},
		{name: "Danish", code: "da"},
		{name: "Dutch", code: "nl"},
		{name: "English", code: "en"},
		{name: "Esperanto", code: "eo"},
		{name: "Estonian", code: "et"},
		{name: "Finnish", code: "fi"},
		{name: "French", code: "fr"},
		{name: "Hindi", code: "hi"},
		{name: "Italian", code: "it"},
		{name: "Korean", code: "ko"},
		{name: "Latin", code: "la"},
		{name: "Marathi", code: "mr"},
		{name: "Nepali", code: "ne"},
		{name: "Persian", code: "fa"},
		{name: "Punjabi", code: "pa"},
		{name: "Russian", code: "ru"},
		{name: "Serbian", code: "sr"},
		{name: "Tamil", code: "ta"},
		{name: "Telugu", code: "te"},
		{name: "Thai", code: "th"},
	}

	// myApp := app.New()
	// myWindow := myApp.NewWindow("Language Detector")

	label1 := canvas.NewText("Language Detector", color.White)
	label1.Alignment = fyne.TextAlignCenter
	label1.TextSize = 30

	entry := widget.NewMultiLineEntry()
	entry.SetPlaceHolder("Enter Your Text...")
	entry.Wrapping = fyne.TextWrapBreak

	message := widget.NewLabel("Language: ")
	message.Alignment = fyne.TextAlignCenter
	message.TextStyle = fyne.TextStyle{Monospace: true}
	message.Hide()

	detectBtn := widget.NewButtonWithIcon("Detect Language", theme.SearchIcon(), func() {
		url := "https://google-translate1.p.rapidapi.com/language/translate/v2/detect"

		payload := strings.NewReader("q=" + strings.ReplaceAll(strings.ReplaceAll(entry.Text, " ", "%20"), ",", "%2c"))

		req, _ := http.NewRequest("POST", url, payload)

		req.Header.Add("content-type", "application/x-www-form-urlencoded")
		req.Header.Add("accept-encoding", "application/JSON")
		req.Header.Add("x-rapidapi-host", "google-translate1.p.rapidapi.com")
		req.Header.Add("x-rapidapi-key", "--------------api-Key-------------------")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Print(err)
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		data, err := UnmarshalLanguage(body)
		if err != nil {
			fmt.Print(err)
		} else {
			var flag bool = false
			for i := 0; i < len(languages); i++ {
				if languages[i].code == data.Data.Detections[0][0].Language {
					message.Text = "Language: " + languages[i].name
					message.Refresh()
					message.Show()
					flag = true
					break
				}
			}
			if !flag {
				message.Text = "Language: Not Found"
				message.Refresh()
			}
		}
	})

	myWindow.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, container.NewVScroll(
			container.NewVBox(
				container.New(
					layout.NewCenterLayout(),
					label1,
				),
				container.New(
					layout.NewGridLayout(1),
					entry,
				),
				container.New(
					layout.NewGridLayout(1),
					detectBtn,
				),
				container.New(
					layout.NewCenterLayout(),
					message,
				),
			),
		)),
	)
	w.CenterOnScreen()
	// myWindow.Resize(fyne.NewSize(400, 400))
	w.Show()
}
