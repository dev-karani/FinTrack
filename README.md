# FinTrack

FinTrack is a RESTful personal finance API built with Go that enables users to securely manage their financial transactions. It provides user authentication, transaction management, and a scalable backend architecture that serves as the foundation for personal finance applications.

The API is designed around modern backend development practices, including layered architecture, JWT authentication, PostgreSQL, and database migrations, making it easy to extend with additional financial features over time.

## Features

### Authentication

* User registration
* User login
* JWT-based authentication
* Refresh token authentication
* Token revocation (logout)
* Secure password hashing
* User profile updates
* Soft deletion of users

### Transactions

* Create transactions
* Retrieve all user transactions
* Retrieve a single transaction
* Update transactions
* Soft delete transactions

### Backend Features

* RESTful API design
* Versioned API (`/api/v1`)
* PostgreSQL database
* SQL generated with sqlc
* Database migrations using Goose
* UUID primary keys
* Layered Handler → Service → Database architecture
* JSON request and response handling
* JWT authentication and authorization
* Soft delete support

## Project Structure

```text
.
├── cmd/
├── internal/
│   ├── auth/
│   ├── database/
│   ├── httpx/
├── sql/
│   ├── queries/
│   └── schema/
├── Transactions/ 
├── Users/
├── main.go
└── routes.go
```

## Technology Stack

* Go
* PostgreSQL
* sqlc
* Goose
* JWT Authentication
* UUIDs

## API Endpoints

### Users

| Method | Endpoint          | Description                        |
| ------ | ----------------- | ---------------------------------- |
| POST   | `/api/v1/users`   | Register a new user                |
| POST   | `/api/v1/login`   | Authenticate a user                |
| POST   | `/api/v1/refresh` | Generate a new access token        |
| POST   | `/api/v1/revoke`  | Revoke active refresh tokens       |
| PUT    | `/api/v1/users`   | Update the authenticated user      |
| DELETE | `/api/v1/users`   | Soft delete the authenticated user |

### Transactions

| Method | Endpoint                               | Description                     |
| ------ | -------------------------------------- | ------------------------------- |
| POST   | `/api/v1/transactions`                 | Create a transaction            |
| GET    | `/api/v1/transactions`                 | Retrieve all user transactions  |
| GET    | `/api/v1/transactions/{transactionID}` | Retrieve a specific transaction |
| PUT    | `/api/v1/transactions/{transactionID}` | Update a transaction            |
| DELETE | `/api/v1/transactions/{transactionID}` | Soft delete a transaction       |

## Getting Started

Clone the repository:

```bash
git clone https://github.com/dev-karani/FinTrack.git
cd FinTrack
```

Run the database migrations:

```bash
goose postgres "<DATABASE_URL>" up
```

Start the server:

```bash
go run .
```

The API will be available at:

```text
http://localhost:8080
```

## Roadmap

Planned improvements include:

* Pagination
* Filtering and search
* Comprehensive testing
* Structured logging
* Docker and Docker Compose
* CI/CD with GitHub Actions
* Graceful shutdown
* Health monitoring
* Cloud deployment
