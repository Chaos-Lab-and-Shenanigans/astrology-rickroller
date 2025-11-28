package tappedfunctions

import (
	"database/sql"
	"fmt"
)

func TapRestoreDesktop(db *sql.DB, x chan string, path string, pathDB string) func() {
	return func() {
		errCh1 := make(chan error)
		//		errCh2 := make(chan error)

		//		go restoreWallpaper(errCh2)
		go restoreNames(db, path, errCh1)
		go stopAudio()

		err := <-errCh1
		if err != nil {
			x <- fmt.Sprintf("%v", err)
			return
		}
		/*
			err = <-errCh2
			if err != nil {
				x <- fmt.Sprintf("Successfully recovered the madness\n %v", err)
				return
			}
		*/
		x <- "Successfully recovered the madness"
	}
}
