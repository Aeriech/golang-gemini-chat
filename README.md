## **Getting Started**

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