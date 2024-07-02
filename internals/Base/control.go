package Base

import (
	"fmt"
	"log"
	"sunsend/internals/DB"
	"time"
)

const (
	DB_DURATION      time.Duration = time.Second * 5 // 1 Minute
	ARCHIVE_DURATION time.Duration = time.Minute * 1 // ~~~
)

var (
	DB_NOW      time.Time = time.Now().Add(DB_DURATION)
	ARCHIVE_NOW time.Time = time.Now().Add(ARCHIVE_DURATION) // TODO: Use AddDate instead of Add
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
	log.Printf("[INFO] There is %d Amount of message we have in the DB", Count)
}

func ControlUnit(timer time.Time) {
	ctimer := time.Now()
	//log.Println(ctimer, ARCHIVE_DURATION)
	switch res := ctimer; {
	case res.Compare(DB_NOW) >= 0:
		ControlDB()
		DB_NOW = time.Now().Add(DB_DURATION)
	case res.Compare(ARCHIVE_NOW) >= 0:
		oldmessageorg, res := GetOldMessages()
		if res == 0 {
			for _, message := range *oldmessageorg {
				// Archive the Last 5 Messages
				AddArchive(&message)
				// Delete The Last 5 Messages
				res := DB.QueryRow(fmt.Sprintf("DELETE FROM Messages WHERE MID = %d", message.MID))
				var count int
				res.Scan(&count)
				log.Println("[INFO] Archived and Removed from the database", message.MID)
			}
		}
		//log.Println("The result value is:", res)
		//log.Println(oldmessageorg)
		ARCHIVE_NOW = time.Now().Add(ARCHIVE_DURATION)
	}
	//ControlDB()

}
