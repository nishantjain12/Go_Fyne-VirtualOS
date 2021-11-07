package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func UnmarshalNews(data []byte) (News, error) {
	var r News
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *News) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type News struct {
	Status       string    `json:"status"`
	TotalResults int64     `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Article struct {
	Source      Source  `json:"source"`
	Author      *string `json:"author"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	URL         string  `json:"url"`
	URLToImage  string  `json:"urlToImage"`
	PublishedAt string  `json:"publishedAt"`
	Content     string  `json:"content"`
}

type Source struct {
	ID   *string `json:"id"`
	Name string  `json:"name"`
}

var count int = 0
var news1 News

func showNewsApp(w fyne.Window) {
	// myApp := app.New()
	// myWindow := myApp.NewWindow("News App")

	label1 := canvas.NewText("News App", color.White)
	label1.TextStyle = fyne.TextStyle{Bold: true}

	searchBox := widget.NewEntry()
	searchBox.TextStyle = fyne.TextStyle{Bold: true}
	searchBox.SetPlaceHolder("Search News...")

	label2 := widget.NewLabel("Latest News...")
	label2.TextStyle = fyne.TextStyle{Bold: true}
	label2.Alignment = fyne.TextAlignCenter
	label2.Wrapping = fyne.TextWrapBreak

	label3 := widget.NewLabel(fmt.Sprintf("News: %s", strconv.Itoa(count+1)))
	label3.TextStyle = fyne.TextStyle{Bold: true}
	label3.Alignment = fyne.TextAlignCenter
	label3.Wrapping = fyne.TextWrapBreak
	label3.Hide()

	img := canvas.NewImageFromResource(nil)
	img.FillMode = canvas.ImageFillContain

	imgcontainer := container.NewGridWrap(
		fyne.NewSize(400, 200),
		img,
	)
	image := container.New(
		layout.NewCenterLayout(),
		imgcontainer,
	)

	title := widget.NewLabel("loading...")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Wrapping = fyne.TextWrapBreak

	source := widget.NewLabel("loading...")
	source.TextStyle = fyne.TextStyle{Monospace: true}
	source.Wrapping = fyne.TextWrapBreak

	Authour := widget.NewLabel("loading...")
	Authour.TextStyle = fyne.TextStyle{Italic: true}
	Authour.Wrapping = fyne.TextWrapBreak
	Authour.Alignment = fyne.TextAlignTrailing

	description := widget.NewLabel("loading...")
	description.TextStyle = fyne.TextStyle{Italic: true}
	description.Wrapping = fyne.TextWrapBreak

	ReadMore := widget.NewHyperlink("Read more", nil)
	ReadMore.SetURLFromString("www.google.com")
	ReadMore.Refresh()

	publish := widget.NewLabel("loading...")
	publish.TextStyle = fyne.TextStyle{Bold: true}
	publish.Wrapping = fyne.TextWrapBreak
	publish.Alignment = fyne.TextAlignLeading

	prevButton := widget.NewButton("Prev", func() {
		if count > 0 {
			count--
		}
		r, _ := fyne.LoadResourceFromURLString(news1.Articles[count].URLToImage)
		img.Resource = r
		img.Refresh()

		title.Text = news1.Articles[count].Title
		title.Refresh()

		source.Text = news1.Articles[count].Source.Name
		source.Refresh()

		Authour.Text = *news1.Articles[count].Author
		Authour.Refresh()

		description.Text = news1.Articles[count].Description
		description.Refresh()

		publish.Text = news1.Articles[count].PublishedAt
		publish.Refresh()

		ReadMore.SetURLFromString(news1.Articles[count].URL)
		ReadMore.Refresh()

		label3.Text = fmt.Sprintf("News: %s", strconv.Itoa(count+1))
		label3.Refresh()

	})
	nextButton := widget.NewButton("Next", func() {
		if count < len(news1.Articles)-1 {
			count++
		}
		r, _ := fyne.LoadResourceFromURLString(news1.Articles[count].URLToImage)
		img.Resource = r
		img.Refresh()

		title.Text = news1.Articles[count].Title
		title.Refresh()

		source.Text = news1.Articles[count].Source.Name
		source.Refresh()

		Authour.Text = *news1.Articles[count].Author
		Authour.Refresh()

		description.Text = news1.Articles[count].Description
		description.Refresh()

		publish.Text = news1.Articles[count].PublishedAt
		publish.Refresh()

		ReadMore.SetURLFromString(news1.Articles[count].URL)
		ReadMore.Refresh()

		label3.Text = fmt.Sprintf("News: %s", strconv.Itoa(count+1))
		label3.Refresh()
	})

	changeNews := container.New(
		layout.NewGridLayout(2),
		prevButton,
		nextButton,
	)
	changeNews.Hide()
	changeNews.Refresh()

	res1, err := http.Get("https://newsapi.org/v2/top-headlines?country=in&apiKey=---------------api-key-------------")
	if err != nil {
		fmt.Print(err)
	} else {
		body, err := ioutil.ReadAll(res1.Body)
		if err != nil {
			fmt.Print(err)
		} else {
			latestNews, err := UnmarshalNews(body)
			if err != nil {
				fmt.Print(err)
			} else {

				r, _ := fyne.LoadResourceFromURLString(latestNews.Articles[0].URLToImage)
				img.Resource = r
				img.Refresh()

				title.Text = latestNews.Articles[0].Title
				title.Alignment = fyne.TextAlignCenter
				title.Refresh()

				source.Text = latestNews.Articles[0].Source.Name
				source.Alignment = fyne.TextAlignTrailing
				source.Refresh()

				Authour.Text = *latestNews.Articles[0].Author
				Authour.Alignment = fyne.TextAlignTrailing
				source.Refresh()

				description.Text = latestNews.Articles[0].Description
				description.Alignment = fyne.TextAlignCenter
				description.Refresh()

				publish.Text = latestNews.Articles[0].PublishedAt
				publish.Alignment = fyne.TextAlignLeading
				publish.Refresh()

				ReadMore.SetURLFromString(latestNews.Articles[0].URL)
				ReadMore.Refresh()
			}
		}
	}

	searchBtn := widget.NewButton("Search", func() {
		count = 0

		changeNews.Show()
		changeNews.Refresh()

		label3.Show()
		label3.Refresh()

		res, err := http.Get("https://newsapi.org/v2/everything?qInTitle=" + strings.ReplaceAll(searchBox.Text, " ", "%20") + "&from=2021-10-31&sortBy=popularity&apiKey=---------------api-key-------------")
		if err != nil {
			fmt.Print(err)
			fmt.Print("1")
		} else {
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Print(err)
				fmt.Print("2")
			} else {
				news, err := UnmarshalNews(body)
				if err != nil {
					fmt.Print(err)
					fmt.Print("3")
				} else {
					label2.Text = (strconv.FormatInt(news.TotalResults, 10) + " results found for: " + "\"" + searchBox.Text + "\"")
					label2.Refresh()
					news1 = news

					r, _ := fyne.LoadResourceFromURLString(news1.Articles[count].URLToImage)
					img.Resource = r
					img.Refresh()

					title.Text = news1.Articles[count].Title
					title.Refresh()

					source.Text = news1.Articles[count].Source.Name
					source.Refresh()

					Authour.Text = *news1.Articles[count].Author
					Authour.Refresh()

					description.Text = news1.Articles[count].Description
					description.Refresh()

					publish.Text = news1.Articles[count].PublishedAt
					publish.Refresh()

					ReadMore.SetURLFromString(news1.Articles[count].URL)
					ReadMore.Refresh()

					label3.Text = fmt.Sprintf("News: %s", strconv.Itoa(count+1))
					label3.Refresh()
				}
			}
		}
	})

	myWindow.SetContent(
		container.NewVScroll(
			container.NewVBox(
				container.New(
					layout.NewCenterLayout(),
					label1,
				),
				container.New(
					layout.NewGridLayoutWithRows(2),
					searchBox,
					searchBtn,
				),
				container.New(
					layout.NewGridLayout(1),
					label2,
				),
				container.New(
					layout.NewGridLayout(1),
					label3,
				),
				container.New(
					layout.NewGridLayout(1),
					image,
				),
				container.New(
					layout.NewGridLayout(1),
					title,
				),
				container.New(
					layout.NewGridLayout(1),
					description,
				),
				container.New(
					layout.NewCenterLayout(),
					ReadMore,
				),
				container.New(
					layout.NewGridLayout(1),
					source,
				),
				container.New(
					layout.NewGridLayoutWithColumns(2),
					publish,
					Authour,
				),
				container.New(
					layout.NewGridLayout(1),
					changeNews,
				),
			),
		),
	)
	w.CenterOnScreen()
	// myWindow.Resize(fyne.NewSize(500, 600))
	w.Show()
}
