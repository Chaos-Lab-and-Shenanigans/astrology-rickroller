package main

import (
	"os"
	"path/filepath"

	"fyne.io/fyne/v2/widget"
)

var (
	HOMEDIR, _ = os.UserHomeDir()
	DESKTOP    = filepath.Join(HOMEDIR, "Desktop")
	check      = widget.Check{}
)
