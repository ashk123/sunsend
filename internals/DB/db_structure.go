package DB

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

// type Channel struct {
// 	ID          int
// 	Name        string
// 	Description string
// 	Owner       string
// }

func createBaseTable() int {
	db_work := getbasedb()

	base_channels_table := `CREATE TABLE Channels (
		"ID" integer NOT NULL PRIMARY KEY,		
		"Name" TEXT,
		"Description" TEXT,
		"Owner" TEXT
	  );` // SQL Statement for Create Table
	defer db_work.Close()
	log.Println("Create Base Channels table...")
	statement_channel, err := db_work.Prepare(base_channels_table) // Prepare SQL Statement
	if err != nil {
		fmt.Println("this is where I stoped the program hahahahaha")
		log.Fatalln(err.Error())
	}
	statement_channel.Exec() // Execute SQL Statements
	log.Println("Base Channels table created")
	// CreateMessageTable()

	base_messages_table := `CREATE TABLE Messages (
		"CID" integer NOT NULL,	
		"MID" integer NOT NULL PRIMARY KEY,	
		"Author" TEXT,
		"Content" TEXT,
		"Date" TEXT,
		"ReplyID" Integer
	  );` // SQL Statement for Create Table
	defer db_work.Close()
	log.Println("Create Base Messages table...")
	statement_msg, err := db_work.Prepare(base_messages_table) // Prepare SQL Statement
	if err != nil {
		// log.Fatalln(err.Error())
		return 16
	}
	statement_msg.Exec() // Execute SQL Statements
	log.Println("Base Messages table created")
	return 0
}
func QueryRows(query string) (*sql.Rows, int) {
	db := getbasedb()
	row, err := db.Query(query)
	if err != nil {
		// log.Fatal(err)
		return nil, 16
	}
	defer db.Close()
	// defer row.Close()
	return row, 0
}
func QueryRow(query string) (*sql.Row, int) {
	db := getbasedb()
	row := db.QueryRow(query)
	if row == nil {
		// log.Fatal(err)
		return nil, 16
	}
	defer db.Close()
	// defer row.Close()
	return row, 0
}

// func QueryChannel() (*sql.Rows, int) {
// 	db := getbasedb()
// 	row, err := db.Query("SELECT * FROM Channels")
// 	if err != nil {
// 		// log.Fatal(err)
// 		return nil, 16
// 	}
// 	defer db.Close()
// 	// defer row.Close()
// 	return row, 0
// }

// func QueryMessages(query string) (*sql.Rows, int) {
// 	db := getbasedb()
// 	row, err := db.Query("SELECT * FROM Messages")
// 	if err != nil {
// 		// log.Fatal(err)
// 		return nil, 16
// 	}
// 	defer db.Close()
// 	// defer row.Close()
// 	return row, 0
// }

func InsertChannel(user_ID int, user_Name string, user_Descriptin string, user_Owner string) int {
	db := getbasedb()

	insertChannelBase := `INSERT INTO Channels(ID, Name, Description, Owner) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertChannelBase) // Prepare statement.
	defer db.Close()
	// This is good to avoid SQL injections
	if err != nil {
		// log.Fatalln(err.Error())
		log.Println(err.Error())
		return 16
	}
	_, err = statement.Exec(user_ID, user_Name, user_Descriptin, user_Owner)
	if err != nil {
		// log.Fatalln(err.Error())
		log.Println(err.Error())
		return 16
	}
	log.Println("Channel", user_Name, "Created SuccsessFuly.")
	return 0
}

func InsertMsg(CID int, MID int, Author string, Content string, Date string, ReplyID int) int {
	db := getbasedb()
	inserMessageBase := `INSERT INTO Messages(CID, MID, Author, Content, Date, ReplyID) VALUES (?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(inserMessageBase) // Prepare statement.
	defer db.Close()
	// This is good to avoid SQL injections
	if err != nil {
		// log.Fatalln(err.Error())
		return 16
	}

	_, err = statement.Exec(CID, MID, Author, Content, Date, ReplyID)
	if err != nil {
		// log.Fatalln(err.Error())
		return 16
	}
	log.Println("User", Author, "Sent Message to channel", CID)
	return 0
}
