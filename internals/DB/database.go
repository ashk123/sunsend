package DB

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func createDataBase() {
	log.Printf("Database is not exists, making one ...")
	file, err := os.Create("Storage/SunSend.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	db, err := sql.Open("sqlite3", "Storage/SunSend.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
}

func PrepairDBSystem() {
	if _, err := os.Stat("Storage/SunSend.db"); os.IsNotExist(err) {
		createDataBase()
	} else {
		log.Printf("There is a database.")
	}
}
