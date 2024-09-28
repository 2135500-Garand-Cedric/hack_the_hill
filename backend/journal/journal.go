package journal

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	"hackthehill/backend/ai"
	"hackthehill/backend/database"
)


func CheckEntry(username string) (bool, error) {

	db := database.GetJournalDB()

	_, err := database.GetTodayJournal(db, username);

	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	
	return true, nil
	
}

func CreateJournalEntry(c *fiber.Ctx) error {
	
	data := c.FormValue("data")

	db := database.GetJournalDB()

	check, err := CheckEntry(c.Locals("user").(string))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not check entry",
		})
	}

	entry := database.JournalEntry{
		"username": c.Locals("user").(string),
		"entry": 1,
		"date": time.Now().Format("2006-01-02"),
		"data": data,
	}

	if check {
		entry = database.JournalEntry{
			"username": c.Locals("user").(string),
			"entry": 2,
			"date": time.Now().Format("2006-01-02"),
			"data": data,
		}
	}

	entryNum := entry["entry"].(int)

	err = CreateSidebarEntry(data, c.Locals("user").(string), entryNum )
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create sidebar entry",
		})
	}


	err = database.InsertJournalEntry(db, entry)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not insert entry",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Entry created",
	})

}


func CreateSidebarEntry(data string, username string, entryNum int) error {

	data, err := ai.SummerizeJournalEntry(data, username, entryNum)

	if err != nil {
		return err
	}

	db := database.GetSummerizedJournal()

	entry := database.SumJournalEntry{
		"username": username,
		"entry": entryNum,
		"date": time.Now().Format("2006-01-02"),
		"data": data,
	}


	err = database.InsertSummerizedJournalEntry(db, entry)
	if err != nil {
		return err
	}	
	
	return nil

}

