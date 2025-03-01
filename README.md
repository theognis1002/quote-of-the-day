# Quote of the Day API

A simple REST API service that provides AI-generated inspirational quotes. Built with Go, Docker, and modern DevOps practices.

## Technologies Used

- Go (Golang)
- Docker & Docker Compose
- OpenAI GPT API
- Make (for simplified commands)

## Requirements

- Docker (latest version)
- Docker Compose v2+
- Make (optional, but recommended)
- OpenAI API key

## Getting Started

### Setup & Installation

1. Clone this repository
2. Create a `.env` file in the root directory with your OpenAI API key:
   ```
   OPENAI_API_KEY=your-api-key-here
   ```
3. Run `make build` to build and start the containers
   - Alternatively: `docker compose up --build`

### Available Make Commands

- `make build` - Build and start containers
- `make run` - Start existing containers
- `make stop` - Stop running containers
- `make destroy` - Stop containers and remove volumes
- `make rebuild` - Rebuild all containers
- `make nuke` - Complete cleanup of all Docker resources
- `make clear-cache` - Clear the quote cache to force a new quote generation

## API Endpoints

### Root Endpoint

- **URL**: `/`
- **Method**: `GET`
- **Description**: Welcome message
- **Response Format**: JSON

```json
{
  "message": "Welcome! Please hit the `/quote-of-the-day` API to get the quote of the day."
}
```

### Quote of the Day

- **URL**: `/quote-of-the-day`
- **Method**: `GET`
- **Description**: Returns an AI-generated inspirational quote (cached daily)
- **Response Format**: JSON

```json
{
  "message": "Your quote here - Author"
}
```

### Clear Cache

- **URL**: `/clear-cache`
- **Method**: `POST`
- **Description**: Clears the current quote cache, forcing a new quote generation on next request
- **Response Format**: JSON

```json
{
  "message": "Cache cleared successfully"
}
```

## Error Handling

All endpoints return JSON responses. In case of errors, the response will include an error message:

```json
{
  "error": "Error message here"
}
```
