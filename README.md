# Todo List API

A RESTful API for managing todos and users, built with **Go** and **PostgreSQL**. The project follows a clean architecture pattern with layered separation — domain models, service interfaces, repository implementations, and HTTP handlers — using Go's standard `net/http` mux with a custom middleware chain, `sqlx` for database access, and automatic schema migrations.

## Features

- **User Authentication** — Registration and login with custom HMAC-SHA256 JWT implementation
- **Todo CRUD** — Create, read (single + paginated list), update, and delete todos
- **Paginated Responses** — List endpoints return structured pagination metadata (page, limit, totalItems, totalPages)
- **Automatic Migrations** — Schema management via `rubenv/sql-migrate`, applied on startup
- **Custom Middleware Chain** — Composable middleware manager supporting global and per-route middlewares
  - **Logger** — Request logging
  - **CORS** — Cross-Origin Resource Sharing headers
  - **Preflight** — OPTIONS preflight handling
  - **JWT Auth** — Token verification on protected routes
- **Environment Configuration** — `.env`-based config with strict validation via `joho/godotenv`

## Tech Stack

| Component       | Technology                                       |
| --------------- | ------------------------------------------------ |
| **Language**    | Go 1.25+                                         |
| **Database**    | PostgreSQL                                       |
| **Routing**     | Standard `net/http` ServeMux (Go 1.22+ patterns) |
| **DB Access**   | `jmoiron/sqlx`                                   |
| **DB Driver**   | `lib/pq`                                         |
| **Migrations**  | `rubenv/sql-migrate`                             |
| **Auth**        | Custom HMAC-SHA256 JWT (no external JWT library)  |
| **Config**      | `joho/godotenv`                                  |

## Project Structure

```
todo-list/
├── main.go               # Entry point — calls cmd.Serve()
├── cmd/
│   └── serve.go          # Bootstraps config, DB, repos, services, handlers, and starts the server
├── config/
│   └── config.go         # Loads and validates environment variables into Config struct
├── domain/
│   ├── todo.go           # Todos domain model
│   └── user.go           # User domain model
├── infra/
│   └── db/
│       ├── connection.go  # PostgreSQL connection setup (sqlx + lib/pq)
│       └── migrate.go     # Database migration runner
├── migrations/
│   ├── 000001-create-users.up.sql
│   ├── 000001-create-users.down.sql
│   ├── 000002-create-todos.up.sql
│   └── 000002-create-todos.down.sql
├── repo/
│   ├── todo.go           # TodoRepo implementation (CRUD + pagination + count)
│   └── user.go           # UserRepo implementation (create + find)
├── todo/
│   ├── port.go           # Todo service & repository interfaces
│   └── service.go        # Todo service implementation
├── user/
│   ├── port.go           # User service & repository interfaces
│   └── service.go        # User service implementation
├── rest/
│   ├── server.go         # HTTP server setup and startup
│   ├── handler/
│   │   ├── todo/         # Todo HTTP handlers (create, get, list, update, delete, hello)
│   │   └── user/         # User HTTP handlers (register, login)
│   └── middleware/
│       ├── manager.go            # Middleware chain manager (Use, With, Wrapmux)
│       ├── middleware.go         # Middlewares struct (holds config)
│       ├── authentication_jwt.go # JWT verification middleware
│       ├── cors_middleware.go    # CORS headers middleware
│       ├── logger.go            # Request logger middleware
│       └── preflight.go         # OPTIONS preflight middleware
├── util/
│   ├── create_jwt.go     # Custom JWT creation (HMAC-SHA256)
│   ├── send_data.go      # JSON response helpers (SendData, SendError)
│   └── send_page.go      # Paginated response helper (SendPage)
├── db-queries/           # Reference SQL queries
├── .env                  # Environment configuration
├── go.mod
└── go.sum
```

## Getting Started

### Prerequisites

- Go 1.25 or higher
- PostgreSQL database

### Installation & Setup

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd todo-list
   ```

2. **Create the database**

   ```sql
   CREATE DATABASE todo;
   ```

3. **Configure environment variables**

   Create a `.env` file in the project root:

   ```env
   # App Settings
   VERSION=1.0.0
   SERVICE_NAME=TODO
   HTTP_PORT=2222
   JWT_SECRET_KEY=hi!mynameispavel@#

   # Database Settings
   DB_HOST=localhost
   DB_PORT=5432
   DB_NAME=todo
   DB_USER=postgres
   DB_PASSWORD=your_password
   ENABLE_SSL_MODE=false
   ```

4. **Install dependencies**

   ```bash
   go mod download
   ```

5. **Run the application**

   ```bash
   go run main.go
   ```

   The server will start on the configured port (default `2222`). Database migrations run automatically on startup.

## API Endpoints

### Public Routes

| Method | Endpoint           | Description              |
| ------ | ------------------ | ------------------------ |
| `GET`  | `/hello`           | Health check / greeting  |
| `GET`  | `/todos`           | List todos (paginated)   |
| `POST` | `/users/register`  | Register a new user      |
| `POST` | `/users/login`     | Login and receive a JWT  |

### Protected Routes

Requires `Authorization: Bearer <token>` header.

| Method   | Endpoint        | Description         |
| -------- | --------------- | ------------------- |
| `POST`   | `/todos`        | Create a new todo   |
| `GET`    | `/todos/{id}`   | Get a todo by ID    |
| `PUT`    | `/todos/{id}`   | Update a todo       |
| `DELETE` | `/todos/{id}`   | Delete a todo       |

### Request / Response Examples

**Register a user**
```bash
curl -X POST http://localhost:2222/users/register \
  -H "Content-Type: application/json" \
  -d '{"first_name":"John","last_name":"Doe","email":"john@example.com","password":"secret123"}'
```

**Login**
```bash
curl -X POST http://localhost:2222/users/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"secret123"}'
```

**Create a todo** (authenticated)
```bash
curl -X POST http://localhost:2222/todos \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"text":"Buy groceries","description":"Milk, eggs, bread","is_done":false}'
```

**List todos** (paginated)
```bash
curl http://localhost:2222/todos?page=1&limit=10
```

## Database Schema

### `users`

| Column       | Type                | Constraints                    |
| ------------ | ------------------- | ------------------------------ |
| `id`         | `SERIAL`            | Primary Key                    |
| `first_name` | `VARCHAR(100)`      | NOT NULL                       |
| `last_name`  | `VARCHAR(100)`      | NOT NULL                       |
| `email`      | `VARCHAR(255)`      | UNIQUE, NOT NULL               |
| `password`   | `TEXT`              | NOT NULL                       |
| `is_done`    | `BOOLEAN`           | DEFAULT FALSE                  |
| `created_at` | `TIMESTAMP`         | DEFAULT CURRENT_TIMESTAMP      |
| `updated_at` | `TIMESTAMP`         | DEFAULT CURRENT_TIMESTAMP      |

### `todos`

| Column        | Type           | Constraints                    |
| ------------- | -------------- | ------------------------------ |
| `id`          | `SERIAL`       | Primary Key                    |
| `text`        | `TEXT`          | NOT NULL                       |
| `description` | `TEXT`          | Nullable                       |
| `is_done`     | `BOOLEAN`      | NOT NULL, DEFAULT FALSE        |
| `created_at`  | `TIMESTAMP`    | DEFAULT CURRENT_TIMESTAMP      |
| `updated_at`  | `TIMESTAMP`    | DEFAULT CURRENT_TIMESTAMP      |

## Architecture

```
main.go → cmd.Serve()
              │
              ├── config.GetConfig()        ← loads .env
              ├── db.NewConnection()         ← PostgreSQL via sqlx
              ├── db.MigrateDB()            ← auto-migrate on start
              │
              ├── repo.NewTodoRepo(db)       ← data access layer
              ├── repo.NewUserRepo(db)
              │
              ├── todo.NewService(repo)      ← business logic
              ├── user.NewService(repo)
              │
              ├── todoHandler.NewHandler()   ← HTTP handlers
              ├── userHandler.NewHandler()
              │
              └── rest.NewServer().Start()   ← middleware chain + mux
```

## Development

- **Add new migrations**: Create `.sql` files in `migrations/` following the `sql-migrate` naming convention (`NNNNNN-description.up.sql` / `.down.sql`)
- **Add new routes**: Define handler functions in `rest/handler/`, register them in the corresponding `routes.go`
- **Add middleware**: Implement the `func(http.Handler) http.Handler` signature in `rest/middleware/`, then apply via the Manager's `Use()` (global) or `With()` (per-route)
