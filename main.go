package main

import (
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	tappedfunctions "github.com/Chaos-Lab-and-Shenanigans/order-breaker/tapped_functions"
)

func main() {
	app := app.New()
	window := app.NewWindow("2nd window")

	backupL := widget.NewLabel("Logs will appear here")
	window.SetContent(container.NewVBox(
		backupL,
		widget.NewButton("Save backup", tappedfunctions.TapSaveDesktop(backupL, DESKTOP)),
		widget.NewButton("Exit", func() { os.Exit(0) }),
	))

	// window.ShowAndRun()
	window.ShowAndRun()
}
