package tappedfunctions

import (
	"database/sql"
	"fmt"
)

func TapRickRollDesktop(db *sql.DB, x chan string, path string, rickyWall []byte, rickAudioBytes []byte) func() {
	return func() {
		//		errCh1 := make(chan error)
		errCh2 := make(chan error)
		errCh3 := make(chan error)

		//		go setWallpaper(rickyWall, path, errCh1)
		go rickrollNames(db, path, errCh2)
		go playAudio(rickAudioBytes, errCh3)
		/*
			err := <-errCh1
			if err != nil {
				x <- fmt.Sprintf("Error occured while setting wallpaper: %v", err)
				return
			}
		*/
		err := <-errCh2
		if err != nil {
			x <- fmt.Sprintf("Error occured while rickrolling names: %v", err)
			return
		}

		err = <-errCh3
		if err != nil {
			x <- fmt.Sprintf("Error occured while playing audio: %v", err)
			return
		}

		x <- "Check out your desktop brother 🙂"
	}
}
