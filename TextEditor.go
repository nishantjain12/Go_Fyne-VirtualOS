package main

import (
	"image/color"
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type tabs struct {
	fileName string
	text     *widget.Entry
}

func showTextEditor() {
	var count int = 0
	myWindow := myApp.NewWindow("Text Editor")

	tabNodes := []tabs{}

	Tabs := container.NewDocTabs()

	content := container.New(
		layout.NewVBoxLayout(),
	)

	content.Add(widget.NewButton("AddFile", func() {
		count++
		content.Add(container.New(
			layout.NewHBoxLayout(),
			widget.NewLabel("-New File"+strconv.Itoa(count)+".txt"),
		))
		textArea := widget.NewMultiLineEntry()
		textArea.SetPlaceHolder("Enter text...")
		entry := tabs{fileName: "New File" + strconv.Itoa(count) + ".txt", text: textArea}
		tabNodes = append(tabNodes, entry)
		Tabs.Append(container.NewTabItem(entry.fileName, entry.text))
	}))

	hSplitContainer := container.NewHSplit(
		container.NewVBox(
			canvas.NewText("History", color.White),
			content,
		),
		Tabs,
	)
	hSplitContainer.SetOffset(0.2)
	myWindow.SetContent(hSplitContainer)

	//=================================================MainMenu====================================================
	fileItem1 := fyne.NewMenuItem("New File", func() {
		count++
		content.Add(container.New(
			layout.NewHBoxLayout(),
			widget.NewLabel("-New File"+strconv.Itoa(count)+".txt"),
		))
		textArea := widget.NewMultiLineEntry()
		textArea.SetPlaceHolder("Enter text...")
		entry := tabs{fileName: "New File" + strconv.Itoa(count) + ".txt", text: textArea}
		tabNodes = append(tabNodes, entry)
		Tabs.Append(container.NewTabItem(entry.fileName, entry.text))
	})

	fileItem2 := fyne.NewMenuItem("open File", func() {
		openFileDialog := dialog.NewFileOpen(
			func(uc fyne.URIReadCloser, _ error) {
				ReadData, _ := ioutil.ReadAll(uc)

				output := fyne.NewStaticResource("NewFile", ReadData)
				viewText := widget.NewMultiLineEntry()

				viewText.SetText(string(output.StaticContent))

				entry := tabs{fileName: "New File" + strconv.Itoa(count) + ".txt", text: viewText}
				tabNodes = append(tabNodes, entry)
				Tabs.Append(container.NewTabItem(entry.fileName, entry.text))
				content.Add(container.New(
					layout.NewHBoxLayout(),
					widget.NewLabel("-"+entry.fileName),
				))
			}, myWindow)

		openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		openFileDialog.Show()
	})

	fileItem3 := fyne.NewMenuItem("save File", func() {
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(tabNodes[Tabs.SelectedIndex()].text.Text)
				uc.Write(textData)
			}, myWindow)

		saveFileDialog.SetFileName(Tabs.Selected().Text)
		saveFileDialog.Show()
	})

	newMenu1 := fyne.NewMenu("File", fileItem1, fileItem2, fileItem3)

	menu := fyne.NewMainMenu(newMenu1)

	myWindow.SetMainMenu(menu)
	myWindow.Resize(fyne.NewSize(1100, 700))

	//=============================================================================================================
	myWindow.Show()
}
