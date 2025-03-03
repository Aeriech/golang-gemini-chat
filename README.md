## **Getting Started**

### Project Overview - GoChatAI
This project is a CLI-based chatbot built in Golang, leveraging Google's Gemini API for natural language processing (NLP). The chatbot allows users to interact with Gemini AI in a conversational manner, processing user inputs and generating intelligent responses. The application manages chat sessions, retrieves an API key from environment variables, and continuously accepts user input until the conversation is ended.

Key Features:
   - Gemini AI Integration - Uses Gemini-1.5-Flash for quick, smart responses.
   - Command-Line Chat Interface - Enables real-time chat through the terminal.
   - Session-Based Conversations - Maintains context throughout the chat session.
   - Secure API Key Handling - Uses .env file to store and retrieve API credentials.
   - Graceful Exit Handling - Allows users to type 'exit' to end the conversation.

Tech Stack:
   - Golang (Primary Language)
   - Gemini API (AI-powered chatbot)
   - Google Generative AI SDK
   - GoDotEnv (Environment variable management)

### Prerequisites

Ensure you have the following installed:

- [Go](https://go.dev/doc/install) (version 1.20 or higher)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Aeriech/golang-gemini-chat.git
   ```

2. Install dependencies:
   ```bash
   cd golang-gemini-chat
   go mod tidy
   cp .env.example .env
   ```
   
3. Setting Gemini Api Key:
    - Get your api key here: https://aistudio.google.com/app/apikey
    - Put your api key in .env (API_KEY)

4. Run development:
   ```bash
   go run main.go
   ```