package database

import (
	"github.com/wkirk01/AlgoeDB"

	"time"

	"fmt"
	"log"
)

type AdviceDB []map[string]interface{}
type AdviceEntry map[string]interface{}


func GetAdviceDB() *AlgoeDB.Database {
	
	config := AlgoeDB.DatabaseConfig{Path: "./db/advice.json"}

	db, err := AlgoeDB.NewDatabase(&config)
	if err != nil {
		log.Fatal(err)
	}

	return db
	
}


func InsertAdviceEntry(db *AlgoeDB.Database, entry AdviceEntry) error {

	err := db.InsertOne(entry)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
	
}

func GetTodaysAdvice(db *AlgoeDB.Database, username string) (AdviceEntry, error) {

	query := AdviceEntry{"date": time.Now().Format("2006-01-02"), "username": username}

	result := db.FindOne(query)

	if result == nil {
		return AdviceEntry{}, fmt.Errorf("advice not found")
	}

	return result, nil
}