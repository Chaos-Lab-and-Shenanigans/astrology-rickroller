package main

import (
	_ "embed"
	"os"
	"path/filepath"
)

//go:embed assets/wall.png
var picByte []byte

//go:embed assets/audio.mp3
var rickAudio []byte

var (
	HOMEDIR, _ = os.UserHomeDir()
	DESKTOP    = filepath.Join(HOMEDIR, "Desktop")
	BACKUPOB   = filepath.Join(DESKTOP, ".backupOB")
	DATABASE   = filepath.Join(BACKUPOB, "backupob.db")
)
