package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
)

func init() {
	//设置中文字体:解决中文乱码问题
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		if strings.Contains(path, "msyh.ttf") || strings.Contains(path, "simhei.ttf") || strings.Contains(path, "simsun.ttc") || strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Markdown Editor")

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <markdown_file>")
	}

	mdFile := os.Args[1]
	content, err := ioutil.ReadFile(mdFile)
	if err != nil {
		log.Fatal(err)
	}

	// Create a MultiLineEntry to allow text editing
	editor := widget.NewMultiLineEntry()
	editor.SetText(string(content))

	// Add a save button to save the changes
	saveButton := widget.NewButton("Save", func() {
		err := ioutil.WriteFile(mdFile, []byte(editor.Text), 0644)
		if err != nil {
			log.Println("Failed to save file:", err)
		}
	})

	// Create a container with the editor and save button
	contentContainer := container.NewVBox(editor, saveButton)

	// Use a Scroll container to make content scrollable
	scroll := container.NewScroll(contentContainer)

	w.SetContent(scroll)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
