package config

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Fix the Extra label postiton
var (
	count int
	Extra = widget.NewLabel("")
)

func GetHomeExitButtons() *fyne.Container {
	Cfg.Window.Resize(AstWindowSize) //Resize to default size
	Extra.SetText("")                //Make blanck initially
	count = 0
	Extra.Alignment = fyne.TextAlignCenter

	homeExitButtons := container.New(
		layout.NewGridLayout(2),
		widget.NewButton("Exit", HandleExit),
		widget.NewButton("Home", func() { Cfg.ControlCh <- "restart" }),
	)

	if GotRickRolled {
		Extra.Show()
	} else {
		Extra.Hide()
	}

	return homeExitButtons
}

func HandleExit() {
	if GotRickRolled {
		count += 1
		if count == 10 {
			Extra.SetText(Extra.Text + "It'll be your headache if it gets outta screen\n")
		} else {
			Extra.SetText(Extra.Text + "Exit your pathetic life, mf\n")
		}
		Extra.Show()
	} else {
		Cfg.ControlCh <- "exit"
	}

}

func Check() bool {
	if User.ZodiacSign == nil {
		return false
	}

	switch User.Status {
	case IsSingle:
		return true
	case InRelationship:
		return true
	case IsMarried:
		return true
	}

	return false
}
