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

func GetTodayDataEntry1(db *AlgoeDB.Database, username string) (JournalEntry, error) {

	query := JournalEntry{"date": time.Now().Format("2006-01-02"), "username": username, "entry": "1"}

	result := db.FindOne(query)
	
	if result == nil {
		return JournalEntry{}, fmt.Errorf("journal not found")
	}
	
	fmt.Println(result)
	return result, nil
}

func GetTodayDataEntry2(db *AlgoeDB.Database, username string) (JournalEntry, error) {

	query := JournalEntry{"date": time.Now().Format("2006-01-02"), "username": username, "entry": "2"}

	result := db.FindOne(query)
	
	if result == nil {
		return JournalEntry{}, fmt.Errorf("journal not found")
	}
	
	fmt.Println(result)
	return result, nil
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

func GetSummerizedJournalDB() *AlgoeDB.Database {

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

func GetTodaySummerizedJournal(db *AlgoeDB.Database, user string) (SumJournalEntry, error) {
	
	query := SumJournalEntry{"date": time.Now().Format("2006-01-02"), "entry": "1", "username": user}

	result := db.FindOne(query)

	if result == nil {
		fmt.Println(result)
		return SumJournalEntry{}, fmt.Errorf("summerized journal not found")
	}

	return result, nil
	
}

func GetTodaySummerizedReflection(db *AlgoeDB.Database, user string) (SumJournalEntry, error) {

	query := SumJournalEntry{"date": time.Now().Format("2006-01-02"), "username": user, "entry": "2"}

	result := db.FindOne(query)

	if result == nil {
		return SumJournalEntry{}, fmt.Errorf("summerized journal not found")
	}

	return result, nil

}

func GetSummerizedReflectionByDate(db *AlgoeDB.Database, user string, date string) (SumJournalEntry, error) {

	query := SumJournalEntry{"date": date, "username": user, "entry": "2"}

	result := db.FindOne(query)
	
	if result == nil {
		return SumJournalEntry{}, fmt.Errorf("summerized journal not found")
	}
	
	return result, nil
}

func GetSummerizedJournalByDate(db *AlgoeDB.Database, user string, date string) (SumJournalEntry, error) {


	query := SumJournalEntry{"date": date, "username": user, "entry": "1"}


	result := db.FindOne(query)
	
	if result == nil {
		return SumJournalEntry{}, fmt.Errorf("summerized journal not found")
	}
	
	return result, nil
}	
