package main

import (
	"image/color"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type image struct {
	name      string
	path      string
	size      int64
	extension string
}

func showGalleryApp(w fyne.Window) {
	// a := app.New()
	// myWindow := a.NewWindow("Gallery App")

	// temp := container.NewCenter(widget.NewLabel("Hello Fyne!"))

	root_src := "C:\\Users\\nisha\\OneDrive\\Pictures\\Camera Roll"

	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Print(files[0])

	images := []image{}
	for _, file := range files {
		if !file.IsDir() {
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "png" || extension == "jpg" || extension == "jpeg" {
				img := image{name: strings.Split(file.Name(), ".")[0], path: root_src + "\\" + file.Name(), size: file.Size(), extension: extension}
				images = append(images, img)
			}
		}
	}

	tabs := container.NewAppTabs()

	for i := 0; i < len(images); i++ {
		name := widget.NewEntry()
		name.TextStyle = fyne.TextStyle{Bold: true}
		name.SetPlaceHolder("null")
		name.Disable()
		name.SetText(images[i].name)

		extension := widget.NewEntry()
		extension.TextStyle = fyne.TextStyle{Bold: true}
		extension.SetPlaceHolder("null")
		extension.Disable()
		extension.SetText(images[i].extension)

		Size := widget.NewEntry()
		Size.TextStyle = fyne.TextStyle{Bold: true}
		Size.SetPlaceHolder("null")
		Size.Disable()
		Size.SetText(strconv.FormatInt(images[i].size, 10))

		Path := widget.NewEntry()
		Path.TextStyle = fyne.TextStyle{Bold: true}
		Path.SetPlaceHolder("null")
		Path.Disable()
		Path.SetText(images[i].path)

		HsplitContainer := container.NewHSplit(
			container.NewAdaptiveGrid(1, canvas.NewImageFromFile(images[i].path)),
			container.NewVBox(
				container.New(
					layout.NewGridLayout(2),
					canvas.NewText("Name: ", color.White),
					name,
				),
				container.New(
					layout.NewGridLayout(2),
					canvas.NewText("Extension: ", color.White),
					extension,
				),
				container.New(
					layout.NewGridLayout(2),
					canvas.NewText("Size: ", color.White),
					Size,
				),
				container.New(
					layout.NewGridLayout(2),
					canvas.NewText("Path: ", color.White),
					Path,
				),
			),
		)
		HsplitContainer.SetOffset(0.7)
		tabs.Append(container.NewTabItem(images[i].name, HsplitContainer))
	}

	tabs.SetTabLocation(container.TabLocationBottom)
	myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, tabs))
	w.CenterOnScreen()
	w.Show()
}
