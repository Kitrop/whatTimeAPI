# WhatTimeAPI

A lightweight API service written in Go that returns the day of the week for a given date. Uses Redis for caching and Chi for routing.

## Features

- Get weekday name by date in multiple formats
- Redis caching for previously calculated dates
- Custom error handling for invalid dates and routes
- Environment variables configuration
- Chi router with middleware support

## Installation

### Prerequisites
- Go 1.23+
- Redis server

1. Clone the repository:
```bash
git clone https://github.com/Kitrop/whatTimeAPI.git
cd whatTimeAPI
```

2. Install dependencies:
```bash
go mod download
```

3. Create environment file:
```bash
touch .env
```

4. Configure .env file:
```bash
PORT=:3000
REDIS_HOST=localhost:6379
REDIS_PASSWORD=yourpassword
```

## Prerequisites
1. Start Redis server

2.
```bash
go run main.go
```

### Usage

Make GET requests to /weekdate/{date} endpoint with date in one of these formats:

- DD.MM.YYYY
- DD-MM-YYYY
- DD:MM:YYYY

Example Requests

### Example Requests

*Valid date:*
```bash
curl http://localhost:3000/weekdate/31.12.2024
```

*Response:*
```bash
Tuesday
```


*Invalid date:*
```bash
invalid date format
```

*Response:*
```bash
invalid date format
```