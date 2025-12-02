package astrology

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func secondPage() {
	radioL := widget.NewLabel("Choose your status")
	radioW := widget.NewRadioGroup(
		[]string{"Single", "Relationship", "Married"},
		radioFunc,
	)

	radioContainer := container.New(layout.NewCenterLayout(), radioW)

	nextB := widget.NewButton(
		"Next",
		func() {
			if check() {
				thirdPage()
			}
		},
	)

	backB := widget.NewButton(
		"Back",
		func() {
			firstPage()
		},
	)

	navigation := container.New(layout.NewGridLayout(2), backB, nextB)

	cfg.Window.SetContent(container.NewVBox(
		radioL,
		radioContainer,
		layout.NewSpacer(),
		navigation,
	),
	)
}

func check() bool {
	if player.zodiacSign == nil {
		return false
	}

	switch player.status {
	case isSingle:
		return true
	case inRelationship:
		return true
	case isMarried:
		return true
	}

	return false
}

func radioFunc(s string) {
	if s == "Single" {
		player.status = isSingle
		return
	}

	if s == "Relationship" {
		player.status = isMarried
		return
	}

	if s == "Married" {
		player.status = isMarried
	}

}
