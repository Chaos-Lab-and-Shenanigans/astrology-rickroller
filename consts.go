package main

import (
	_ "embed"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

//go:embed assets/wall.png
var rickWall []byte

//go:embed assets/audio.mp3
var rickAudio []byte

// Main window's
var (
	myApp     = app.New()
	windowAst = myApp.NewWindow("Astrology")
	logsL     = widget.NewLabel("")
	logsCh    = setUpdaterChannel(logsL)
	controlCh = getControlCh()
)
