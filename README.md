# Book Dictionary REST API

A RESTful API built with Go, Gin framework, and PostgreSQL for managing a book dictionary with user authentication.

## Project Structure

```
├── cmd/server/           # Application entry point
├── internal/
│   ├── config/          # Configuration management
│   ├── models/          # Data models/entities
│   ├── repository/      # Data access layer
│   │   ├── interfaces/  # Repository interfaces
│   │   └── postgres/    # PostgreSQL implementations
│   ├── service/         # Business logic layer
│   ├── handler/         # HTTP handlers (controllers)
│   ├── middleware/      # HTTP middleware
│   ├── database/        # Database connection & migrations
│   └── utils/           # Utility functions
├── pkg/                 # Public packages
├── migrations/          # SQL migration files
├── docker-compose.yml   # PostgreSQL setup
└── .env                # Environment variables
```

## Features

- **Authentication**: JWT-based authentication
- **Authorization**: Role-based access (Regular User, Admin)
- **Database**: PostgreSQL with migration support
- **Architecture**: Clean architecture with repository pattern

## User Roles

- **Regular User**: Can view books
- **Admin**: Can create, read, update, and delete books

## Getting Started

1. Start PostgreSQL:
   ```bash
   docker-compose up -d
   ```

2. Run the application:
   ```bash
   go run cmd/server/main.go
   ```

## API Endpoints

### Books
- `GET /books` - Get all books (authenticated)
- `GET /books/:id` - Get book by ID (authenticated)
- `POST /books` - Create book (admin only)
- `PUT /books/:id` - Update book (admin only)
- `DELETE /books/:id` - Delete book (admin only)

### Authentication
- `POST /auth/register` - Register new user
- `POST /auth/login` - Login user
- `POST /auth/refresh` - Refresh JWT token
