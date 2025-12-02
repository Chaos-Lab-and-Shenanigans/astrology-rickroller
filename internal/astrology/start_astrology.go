package astrology

import (
	"database/sql"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func StartAstro() func() {
	return func() {
		initNav()
		firstPage()
	}
}

func initNav() {
	homeB := widget.NewButton("Home", func() { cfg.RestartCh <- "restart" })
	exitB := widget.NewButton("Exit", func() { fyne.CurrentApp().Quit() })

	navigationButtons = container.New(layout.NewGridLayout(2), exitB, homeB)
}

func InitConfig(db *sql.DB, path string, dbName string, rickAudioBytes *[]byte, rickWallBytes *[]byte, w fyne.Window, logsCh chan string, restartCh chan string) {
	cfg.DB = db
	cfg.LogsCh = logsCh
	cfg.RestartCh = restartCh
	cfg.Path = path
	cfg.PathDB = filepath.Join(path, dbName)
	cfg.RickyAudioBytes = rickAudioBytes
	cfg.RickyWall = rickWallBytes
	cfg.Window = w
}
