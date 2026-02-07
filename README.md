# Todo List API

A robust RESTful API for managing todos and users, built with Go and MySQL. This project demonstrates a clean architecture approach using the standard library `net/http` mux with custom middleware, `sqlx` for database interactions, and automatic migrations.

## Features

-   **User Authentication**: Secure registration and login using JWT (JSON Web Tokens).
-   **Todo Management**: Full CRUD (Create, Read, Update, Delete) operations for todos.
-   **Database Migrations**: Automatic schema management using `rubenv/sql-migrate`.
-   **Middleware**: Custom middleware for Logging, CORS, and Authentication.
-   **Configuration**: Environment-based configuration for easy deployment.

## Tech Stack

-   **Language**: Go (v1.25+)
-   **Database**: MySQL
-   **Routing**: Standard `net/http` ServeMux
-   **ORM/DB Tool**: `jmoiron/sqlx`
-   **Migrations**: `rubenv/sql-migrate`
-   **Auth**: `golang-jwt/jwt` (implied from usage)

## Project Structure

```
├── cmd/            # Application entry point (Serve function)
├── config/         # Configuration loader (Environment variables)
├── infra/          # Infrastructure setup (Database connection, Migrations)
├── migrations/     # SQL migration files
├── repo/           # Data access layer (Repositories)
├── rest/           # HTTP layer
│   ├── handler/    # Request handlers (Todo, User)
│   └── middleware/ # HTTP Middleware (Auth, CORS, Logger)
└── util/           # Utility functions
```

## Getting Started

### Prerequisites

-   Go 1.25 or higher
-   MySQL Database

### Installation & Setup

1.  **Clone the repository**

    ```bash
    git clone <repository-url>
    cd todo-list
    ```

2.  **Database Setup**

    Create a MySQL database (e.g., `todo_db`).

3.  **Environment Configuration**

    Create a `.env` file in the root directory. You can copy the example below:

    ```env
    # App Settings
    VERSION=1.0
    SERVICE_NAME=todo-api
    HTTP_PORT=8080
    JWT_SECRET_KEY=super-secret-key-change-me

    # Database Settings
    DB_HOST=localhost
    DB_PORT=3306
    DB_NAME=todo_db
    DB_USER=root
    DB_PASSWORD=your_password
    ENABLE_SSL_MODE=false
    ```

4.  **Run the Application**

    ```bash
    go run main.go
    ```

    The server will start on port `8080` (or the port defined in `.env`). Database migrations will run automatically on startup.

## API Endpoints

### User

-   `POST /signup` - Register a new user
-   `POST /login` - Login and receive a JWT

### Todos (Protected Routes)

Requires `Authorization: Bearer <token>` header.

-   `POST /todo` - Create a new todo
-   `GET /todo` - Get a list of todos
-   `PUT /todo` - Update a todo
-   `DELETE /todo/{id}` - Delete a todo (Path parameter ID handling implied)

## Development

To modify the database schema, add new `.sql` files in the `migrations/` directory following the `sql-migrate` naming convention.
