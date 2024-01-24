# Web Hotlist Middleware

This project is a middleware written in Go for collecting and analyzing web request data. It provides an API for retrieving and submitting statistics based on different types (POST, COMMENT, VOTE, etc.) and time ranges (daily, weekly, monthly, and total). All statistics data are stored and managed using Redis.

## Features

- API endpoints for retrieving and submitting statistics data.
- Flexible data structure design to accommodate different types of statistics and time ranges.
- High performance and scalability to handle a large number of concurrent requests.

## Getting Started

### Prerequisites

- Go 1.16 or later
- Redis 6.0 or later

### Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/web-hotlist-middleware.git
```

2. Navigate to the project directory:

```bash
cd web-hotlist-middleware
```

3. Install the dependencies:

```bash
go mod download
```

4. Build the project:

```bash
go build -o web-hotlist-middleware ./cmd/main.go
```

### Usage

1. Start the middleware:

```bash
./web-hotlist-middleware
```

2. Use the following API endpoints:

- GET /api/statistics: Retrieve statistics data.
- POST /api/submit-statistics: Submit new statistics data.

## Testing

To run the unit tests:

```bash
go test ./...
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.