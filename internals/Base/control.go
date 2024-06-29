package Base

import (
	"log"
	"sunsend/internals/DB"
	"time"
)

const (
	DB_DURATION      time.Duration = time.Minute * 1
	ARCHIVE_DURATION time.Duration = time.Minute * 1
)

var (
	DB_NOW      time.Time = time.Now().Add(DB_DURATION)
	ARCHIVE_NOW time.Time = time.Now().Add(ARCHIVE_DURATION)
)

/*
   ControlUnit Information Text
   We use this unit to control and monitor the actual server each 3 seconds
   It includes the monitoring the database and its data, Control the archive
   System and compress files, control the configuration system that include
   (Reload, Restart, Default, etc.) and more.
   This unit runs with the different thread, that means it can works aside
   the actual server.
*/

func ControlDB() {
	res := DB.QueryRow("SELECT COUNT(*) FROM Messages")
	var Count int
	res.Scan(&Count)
	log.Printf("There is %d Amount of message we have in the DB", Count)
}

func ControlUnit(timer time.Time) {
	//log.Println("I'm running this 3 seconds thread")
	ctimer := time.Now()

	switch res := ctimer; {
	case res.Compare(DB_NOW) >= 0:
		ControlDB()
		DB_NOW = time.Now().Add(DB_DURATION)
	case res.Compare(ARCHIVE_NOW) >= 0:
		ArchivCheckSystem()
		ARCHIVE_NOW = time.Now().Add(ARCHIVE_DURATION)
	}
	//ControlDB()
}
