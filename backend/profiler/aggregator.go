package profiler

import (
	"errors"
	"fmt"
	"hackthehill/backend/ai"
	"hackthehill/backend/database"


	"github.com/gofiber/fiber/v2"

	"time"
)



func GetAdvice(c *fiber.Ctx) error {

	jsonData, err := GenerateAdvice(c.Locals("user").(string))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not generate advice",
		})
	}

	return c.Status(fiber.StatusOK).JSON(jsonData)
}


func GenerateAdvice(username string) (map[string]interface{}, error) {

	db := database.GetJournalDB()
	
	data1, err := database.GetTodayDataEntry1(db, username)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("could not get data")
	}
	data2, err := database.GetTodayDataEntry2(db, username)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("could not get data")
	}
	
	response, err := ai.AggregateAdvices(data1["data"].(string), data2["data"].(string), username)
	

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"data": response}, nil

}

func AppendAdvice(username string) error {

	jsonData, err := GenerateAdvice(username)

	if err != nil {
		return fmt.Errorf("could not generate advice: %w", err)
	}

	advice := database.AdviceEntry{
		"username": username,
		"date":     time.Now().Format("2006-01-02"),
		"advice":   jsonData["data"].(string),
	}

	db := database.GetAdviceDB()

	err = database.InsertAdviceEntry(db, advice)
	if err != nil {
		return fmt.Errorf("could not append advice: %w", err)
	}

	return nil
}

