package profiler

import (
	"errors"
	"hackthehill/backend/database"
)





func GenerateAdvice(username string) error{

	db := database.GetJournalDB()
	
	data1, err := database.GetTodayDataEntry1(db, username)
	if err != nil {
		return errors.New("could not get data entry 1")
	}
	data2, err := database.GetTodayDataEntry2(db, username)
	if err != nil {
		return errors.New("could not get data entry 2")
	}
	
	data := data1["data"].(string) + data2["data"].(string)
	
	summary, err := ai.SummarizeJournalEntry(data, username, "1")



	if err != nil {
		return err
	}	

	return nil

}