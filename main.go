package main

import (
	"fmt"
	"log"
	"time"
)

func init() {
	flags()
	if initDB {
		fmt.Println("WARNING! You have launched with the `init` option!\nYOUR DB " + FirSyncDBPath + " IS GOING TO BE DELETED, THIS CANNOT BE UNDONE.")
		for i := 0; i < 61; i++ {
			secondsLeft := 60 - i
			if secondsLeft <= 0 {
				break
			} else {
				fmt.Printf("You have %d seconds to press CTRL-C to abort before this happens.\r", secondsLeft)
			}
			time.Sleep(1 * time.Second)
		}
		if fileExists(FirSyncDBPath) {
			log.Println("Looking for DB: ")
			destroyDB()
		}
		setupDB()
	}
}
func main() {

	log.Println("fir server up")
	go api()
	select {}
}
