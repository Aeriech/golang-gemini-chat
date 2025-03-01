package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

const (
	modelName = "gemini-1.5-flash"
)

type ChatSession struct {
	Chat    *genai.ChatSession
	Context context.Context
}

func main() {
	err := makeChatRequests()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func makeChatRequests() error {
	ctx := context.Background()
	apiKey, err := getApiKey()
	if err != nil {
		return err
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		return fmt.Errorf("error creating client: %w", err)
	}
	defer client.Close()

	gemini := client.GenerativeModel(modelName)
	chat := gemini.StartChat()

	chatSession := &ChatSession{
		Chat:    chat,
		Context: ctx,
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Start chatting with Gemini. Type 'exit' to end.")

	for {
		fmt.Print("You: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Ending chat with Gemini.")
			break
		}

		chatSession.sendChat(input)
	}

	return nil
}

func (chatSession *ChatSession) sendChat(message string) {
	respond, err := chatSession.Chat.SendMessage(chatSession.Context, genai.Text(message))
	if err != nil {
		log.Printf("Error sending message: %v", err)
		fmt.Println("Error communicating with Gemini. Please try again.")
		return
	}

	printRespond(respond)
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

func printRespond(respond *genai.GenerateContentResponse) {
	if len(respond.Candidates) == 0 {
		fmt.Println("Gemini: No response content found.")
		return
	}

	for _, candidate := range respond.Candidates {
		for _, part := range candidate.Content.Parts {
			text, ok := part.(genai.Text)
			if ok {
				fmt.Printf("Gemini: %s\n", string(text))
			} else {
				fmt.Println("Gemini: Response is not text.")
			}
		}
	}
}
