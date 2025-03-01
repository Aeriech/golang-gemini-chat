package main

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	apiKey, err := getApiKey()
	if err != nil {
		log.Fatal(err)
		return
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
		return
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text("Write a story about a magic backpack."))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)
}

func getApiKey() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return "", err
	}

	return apiKey, nil
}
