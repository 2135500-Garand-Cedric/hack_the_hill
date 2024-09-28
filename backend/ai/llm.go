package ai

import (
	"fmt"
	"hackthehill/backend/database"

	"github.com/joho/godotenv"
	"github.com/jpoz/groq"

	"os"
)

func InitGroq() (*groq.Client, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	apiKey := os.Getenv("GROQ_API_KEY")
	client := groq.NewClient(groq.WithAPIKey(apiKey))

	return client, nil
}




func CallLLM(prompt string, userPrompt string) (string, error) {

	client, err := InitGroq()
	if err != nil {
		return "", err
	}

	response, err := client.CreateChatCompletion(groq.CompletionCreateParams{
		Model: "llama3-8b-8192",
		Messages: []groq.Message{
			{
				Role:    "system",
				Content: prompt,
			},
			{
				Role:    "user",
				Content: userPrompt,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	
	return response.Choices[0].Message.Content, nil

}


func SummerizeJournalEntry(entry string, username string, entryType int) (string, error) {

	db := database.GetDB()
	userProfile, err := database.FindProfilesByUsername(db, username)

	if err != nil {
		return "", err
	}
	
	var prompt string

	if entryType == 1 {		
		prompt = fmt.Sprintf(`You are an API who takes in a user's planned task and summarizes it into this format task:description. Feel free to expand the task and description name whatever seems right based on this person:
		%s
		Split the task into multiple cells. Should be ordered in chronological order. Only return JSON without text.`, userProfile)
	} else if entryType == 2 {

		prompt = fmt.Sprintf(`You are an API who takes in a user's reflections of the day and summarizes it into this format task:description. Feel free to expand the task and description name whatever seems right based on this person:
		%s
		Split the task into multiple cells. Should be ordered in chronological order. Only return JSON without text.`, userProfile)
	}

	data, err := CallLLM(prompt, entry)

	if err != nil {
		return "", err
	}

	return data, nil
}

func TestAI() (string, error) {

	client, err := InitGroq()
	if err != nil {
		return "", err
	}

	response, err := client.CreateChatCompletion(groq.CompletionCreateParams{
		Model: "llama3-8b-8192",
		Messages: []groq.Message{
			{
				Role:    "system",
				Content: "You are an API who takes in a users planned task and you split it up into three times sections. Morning, evening, and night. Organize into \"task\":\"description\" . Each time of day can have more than 1 task to split it up like that so on the user interface it can be displayed like todo and calendar. Only return json and no text",
			},
			{
				Role:    "user",
				Content: "i woke up at 8 today and i plan to go to school at 9 and then imma eat at 4 pm and then spend 2 hours working on math and then watch a movie for going to bed at around 11",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	
	println(response.Choices[0].Message.Content)


	return "hello", nil
}