package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()
var myWindow fyne.Window = myApp.NewWindow("Virtual Os")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget
var btn6 fyne.Widget
var DesktopBtn fyne.Widget

var img fyne.CanvasObject

var panelContent *fyne.Container

func main() {
	img = canvas.NewImageFromFile("desktop.jpg")

	btn1 = widget.NewButtonWithIcon("Gallery App", theme.MediaPhotoIcon(), func() {
		showGalleryApp(myWindow)
	})
	btn2 = widget.NewButtonWithIcon("Weather App", theme.ListIcon(), func() {
		showWheatherApp(myWindow)
	})
	btn3 = widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon(), func() {
		showCalculator()
	})
	btn4 = widget.NewButtonWithIcon("Text Editor", theme.DocumentCreateIcon(), func() {
		showTextEditor()
	})
	btn5 = widget.NewButtonWithIcon("Detector", theme.SearchIcon(), func() {
		showLanguageDetector(myWindow)
	})
	btn6 = widget.NewButtonWithIcon("News App", theme.InfoIcon(), func() {
		showNewsApp(myWindow)
	})

	DesktopBtn = widget.NewButtonWithIcon("This Pc", theme.HomeIcon(), func() {
		myWindow.SetContent(
			container.NewBorder(panelContent, nil, nil, nil, img),
		)
	})

	panelContent = container.NewVBox(
		container.NewGridWithColumns(7, DesktopBtn, btn1, btn2, btn3, btn4, btn5, btn6),
	)

	myWindow.Resize(fyne.NewSize(900, 600))
	myWindow.CenterOnScreen()
	myWindow.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, img),
	)
	myWindow.ShowAndRun()
}
