package tappedfunctions

import (
	"fmt"

	"github.com/reujab/wallpaper"
)

const (
	limit = 14
	sep   = "~"
)

var (
	backupWall         = ""
	currentWallPath, _ = wallpaper.Get()
	retErr             = fmt.Errorf("return")
)
