package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func basic() string {
	return "World!"
}

func main() {

	js := mewn.String("./frontend/dist/my-app/main.js")
	css := mewn.String("./frontend/dist/my-app/styles.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "OPCUA Certificate Generator",
		JS:     js,
		CSS:    css,
		Colour: "#FFF",
	})
	app.Bind(basic)
	app.Run()
}
