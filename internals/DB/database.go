package DB

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite" // sqlite without cgo
)

var db *sql.DB

// Create the Base Database
func createDataBase() {
	log.Println("\nDatabase is not exists, making one ...")
	file, err := os.Create("Storage/SunSend.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	db, err = sql.Open("sqlite", "Storage/SunSend.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
}

// A function for prepare some works about database
func PrepairDBSystem() {
	if _, err := os.Stat("Storge"); os.IsNotExist(err) {
		log.Println("There is not any Storage folder, let's craete one ...")
		err := os.Mkdir("Storage", os.ModePerm)
		if err != nil {
			fmt.Println(err.Error()) // If program can't make the Storage folder
		} else {
			log.Println("Storage folder Created.")
		}
	}
	if _, err := os.Stat("Storage/SunSend.db"); os.IsNotExist(err) {
		createDataBase()
		createBaseTable()
		InsertChannel(
			123,
			"Channel_1",
			"this is the default channel",
			"Eiko",
		) // Create default chat Channel
	} else {
		// db = getbasedb()
		log.Println("\nThere is a database.")
	}
}

// it's better to get database from something like getter
func getbasedb() *sql.DB {
	db, err := sql.Open("sqlite", "Storage/SunSend.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
