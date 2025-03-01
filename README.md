# Quote of the Day API

A simple REST API service that provides inspirational quotes. Built with Go, Docker, and modern DevOps practices.

## Technologies Used

- Go (Golang)
- Docker & Docker Compose
- Make (for simplified commands)

## Requirements

- Docker (latest version)
- Docker Compose v2+
- Make (optional, but recommended)

## Getting Started

### Setup & Installation

1. Clone this repository
2. Run `make build` to build and start the containers
   - Alternatively: `docker compose up --build`

### Available Make Commands

- `make build` - Build and start containers
- `make run` - Start existing containers
- `make stop` - Stop running containers
- `make destroy` - Stop containers and remove volumes
- `make rebuild` - Rebuild all containers
- `make nuke` - Complete cleanup of all Docker resources

## API Endpoints

### Root Endpoint

- **URL**: `/`
- **Method**: `GET`
- **Description**: Health check endpoint

### Quote of the Day

- **URL**: `/quote-of-the-day`
- **Method**: `GET`
- **Description**: Returns a random inspirational quote
