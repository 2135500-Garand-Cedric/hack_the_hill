package database

import (
	// "context"
	"fmt"
	"log"

	"time"

	"github.com/wkirk01/AlgoeDB"
)


type JournalDB []map[string]interface{}
type JournalEntry map[string]interface{}


func GetJournalDB() *AlgoeDB.Database {

	config := AlgoeDB.DatabaseConfig{Path: "./db/journal.json"}	
	db, err := AlgoeDB.NewDatabase(&config)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InsertJournalEntry(db *AlgoeDB.Database, entry JournalEntry) error {

	err := db.InsertOne(entry)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
	
}

func GetTodayJournal(db *AlgoeDB.Database, username string) (JournalDB, error) {
	
	query := JournalEntry{"date": time.Now().Format("2006-01-02"), "username": username}


	result := db.FindMany(query)
	
	if result == nil {
		return JournalDB{}, fmt.Errorf("journal not found")
	}
	
	fmt.Println(result)
	return result, nil	
}



type SumJournalDB []map[string]interface{}
type SumJournalEntry map[string]interface{}

func GetSummerizedJournal() *AlgoeDB.Database {

	config := AlgoeDB.DatabaseConfig{Path: "./db/summerizedjournal.json"}	
	db, err := AlgoeDB.NewDatabase(&config)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InsertSummerizedJournalEntry(db *AlgoeDB.Database, entry SumJournalEntry) error {

	err := db.InsertOne(entry)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
	
}


