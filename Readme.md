## Project Workshop

จะกำหนดโจทย์ให้ทุกคนมาร่วมเพิ่ม feature ไปพร้อมกัน

## Setup project

### Requirements
- Go version 1.18 or higher
- Docker and Docker Compose (for PostgreSQL database)

### Installation

1. Clone the repository
```bash
git clone https://github.com/yourusername/project-workshop.git
cd project-workshop
```

2. Install dependencies
```bash
go mod download
```

### Database Setup

The project uses PostgreSQL as its database. You can start the database using Docker Compose:

```bash
docker-compose up -d
```

This will start a PostgreSQL instance with the following configuration:
- Host: localhost
- Port: 5432
- Username: postgres
- Password: postgres
- Database: workshop

You can verify the database connection by accessing the endpoint:

```bash
curl http://localhost:3000/db-status
```

To stop the database:

```bash
docker-compose down
```

To stop the database and remove all data:

```bash
docker-compose down -v
```

### Running the application

To run the application:
```bash
go run cmd/main.go
```

The server will start on http://localhost:3000. You can access the Hello World endpoint by visiting this URL in your browser or using curl:

```bash
curl http://localhost:3000
```

### Running tests

To run all tests:
```bash
go test ./... -v
```

To run only the main endpoint test:
```bash
go test ./cmd -v
```

The test verifies that the GET / endpoint returns "Hello World" with status code 200.

