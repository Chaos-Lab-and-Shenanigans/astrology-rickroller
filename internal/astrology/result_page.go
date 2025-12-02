package astrology

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func resultPage() func() {
	return func() {
		descL := widget.NewLabel("Your personality:")
		result := getResult()

		cfg.Window.SetContent(container.NewVBox(
			descL,
			result,
			layout.NewSpacer(),
			widget.NewSeparator(),
			widget.NewButton("See interesting information", func() { rickrollOrRestore() }),
			navigationButtons,
		))
	}
}

func getResult() *widget.Label {
	zodiac := strings.ToLower(player.zodiacSign.Text)
	result := Roasts[zodiac][player.status]

	resultL := widget.NewLabel(result)
	resultL.Alignment = fyne.TextAlignCenter
	return resultL
}
