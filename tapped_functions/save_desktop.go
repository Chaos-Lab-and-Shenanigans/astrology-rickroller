package tappedfunctions

import (
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2/widget"
	"github.com/otiai10/copy"
)

func TapSaveDesktop(w *widget.Label, src string) func() {
	return func() {
		var text string
		userHomeDir, err := os.UserHomeDir()
		dest := filepath.Join(userHomeDir, ".backupOB")
		if err != nil {
			text = fmt.Sprintf("Could not detect OS's home directory: %v", err)
			w.SetText(text)
		}

		_, err = os.ReadDir(dest)
		if err != nil {
			if os.IsNotExist(err) {

				w.SetText("Backup directory doesn't exist, creating one...")
				err := os.MkdirAll(dest, 0o755)
				if err != nil {
					text = fmt.Sprintf("Error occured: %v", err)
					w.SetText(text)
					return
				}

				w.SetText("Backup directory created successfully!")
			} else {

				text = fmt.Sprintf("Unknown error: %v", err)
				w.SetText(text)
				return
			}
		}

		err = copy.Copy(src, dest)
		if err != nil {
			text = fmt.Sprintf("Error copying data: %v", err)
			w.SetText(text)
			return
		} else {

			text = fmt.Sprintf("Successfully created backup in %v", dest)
			w.SetText(text)
		}
	}
}
