![Two ancient masters playing Go under a tree](image.png)

> "Life is like the game of Go. The board is empty at the beginning, but gradually fills with possibilities - and only the wise know when to play and when to pass."
> — Ancient Chinese Proverb

# Quote of the Day API

A simple REST API service that provides AI-generated inspirational quotes using Go and OpenAI's GPT API.

## Technologies Used

- Go (Golang)
- OpenAI GPT API
- Gin Web Framework

## Requirements

- Go 1.x
- OpenAI API key

## Setup & Installation

1. Clone this repository
2. Create a `.env` file in the root directory:
   ```
   OPENAI_API_KEY=your-api-key-here
   ```
3. Run `make build` to build the application
4. Run `make run` to start the server

## Available Commands

- `make build` - Build the application
- `make run` - Run the built binary
- `make run-local` - Run directly with Go
- `make test` - Run all tests
- `make test-coverage` - Run tests with coverage report
- `make clean` - Clean build artifacts
- `make clear-cache` - Clear the quote cache

## API Endpoints

### GET /

Welcome message and API instructions

```json
{
  "message": "Welcome! Please hit the `/quote-of-the-day` API to get the quote of the day."
}
```

### GET /quote-of-the-day

Returns an AI-generated inspirational quote (cached daily)

```json
{
  "message": "Your quote here - Author"
}
```

### POST /clear-cache

Clears the current quote cache

```json
{
  "message": "Cache cleared successfully"
}
```

## Error Response

```json
{
  "error": "Error message here"
}
```

## Development

### Testing

Run tests:

```bash
make test

# With coverage report
make test-coverage
```

Coverage report is generated as `coverage.html`.
