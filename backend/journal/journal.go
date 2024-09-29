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
		"entry": "1",
		"date": time.Now().Format("2006-01-02"),
		"data": data,
	}

	if check {
		entry = database.JournalEntry{
			"username": c.Locals("user").(string),
			"entry": "2",
			"date": time.Now().Format("2006-01-02"),
			"data": data,
		}
	}

	entryNum := entry["entry"].(string)


	err = CreateSidebarEntry(data, c.Locals("user").(string), entryNum)
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


func CreateSidebarEntry(data string, username string, entryNum string) error {

	data, err := ai.SummerizeJournalEntry(data, username, entryNum)

	if err != nil {
		return err
	}

	db := database.GetSummerizedJournalDB()
	fmt.Println(data)

	cleanedJSON, err := ai.CleanAndFormatJSON(data)
	if err != nil {
		return err
	}	

	entry := database.SumJournalEntry{
		"username": username,
		"entry": entryNum,
		"date": time.Now().Format("2006-01-02"),
		"data": cleanedJSON,
	}


	err = database.InsertSummerizedJournalEntry(db, entry)
	if err != nil {
		return err
	}	
	
	return nil

}


func GetTodaySummerizedJournal(c *fiber.Ctx) error {

	db := database.GetSummerizedJournalDB()

	result, err := database.GetTodaySummerizedJournal(db, c.Locals("user").(string))
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get today's journal",
		})
	}

	fmt.Println(result["data"].( string))

	dat, err := ai.CleanAndFormatJSON(result["data"].( string))


	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not convert to valid JSON",
		})
	}

	data := map[string]interface{}{
		"data": dat,
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func GetTodaySummerizedReflection(c *fiber.Ctx) error {

	db := database.GetSummerizedJournalDB()

	result, err := database.GetTodaySummerizedReflection(db, c.Locals("user").(string))
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get today's reflection",
		})
	}

	fmt.Println(result["data"].( string))

	dat, err := ai.CleanAndFormatJSON(result["data"].( string))


	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not convert to valid JSON",
		})
	}

	data := map[string]interface{}{
		"data": dat,
	}

	return c.Status(fiber.StatusOK).JSON(data)
}



func GetSummerizedJournalByDate(c *fiber.Ctx) error {
	date := c.Query("date")
	db := database.GetSummerizedJournalDB()

	result, err := database.GetSummerizedJournalByDate(db, c.Locals("user").(string), date)
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get journal by date",
		})
	}

	fmt.Println(result["data"].( string))

	dat, err := ai.CleanAndFormatJSON(result["data"].( string))


	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not convert to valid JSON",
		})
	}

	data := map[string]interface{}{	
		"data": dat,
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func GetSummerizedReflectionDate(c *fiber.Ctx) error {
	date := c.Query("date")
	convertDate, err := time.Parse("2006-01-02", date)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get reflection by date",
		})
	}	
	fmt.Println(convertDate.Format("2006-01-02"))
	db := database.GetSummerizedJournalDB()

	result, err := database.GetSummerizedReflectionByDate(db, c.Locals("user").(string), convertDate.Format("2006-01-02"))
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get reflection by date",
		})
	}

	fmt.Println(result["data"].( string))

	dat, err := ai.CleanAndFormatJSON(result["data"].( string))


	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not convert to valid JSON",
		})
	}

	data := map[string]interface{}{
		"data": dat,
	}

	return c.Status(fiber.StatusOK).JSON(data)
}