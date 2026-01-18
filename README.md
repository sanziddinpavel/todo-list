# Todo List API

A simple REST API for managing todos, built with Go using only the standard library.

## Features

- Create, read, update, delete todos
- User registration and login
- JWT authentication
- CORS support

## Project Structure

```
├── cmd/            # Application entry point
├── config/         # Environment configuration
├── database/       # In-memory data storage
├── rest/
│   ├── handler/    # HTTP handlers (todo, user)
│   └── middleware/ # Auth, CORS, logging
└── util/           # JWT and response helpers
```

## Setup

1. Create a `.env` file:

```
VERSION=1.0
SERVICE_NAME=todo-api
HTTP_PORT=8080
JWT_SECRET_KEY=your-secret-key
```

2. Run the server:

```bash
go run main.go
```

## API Endpoints

### Todos (requires auth)
- `GET /todos` - List all todos
- `GET /todo/{id}` - Get a todo
- `POST /todo` - Create a todo
- `PUT /todo/{id}` - Update a todo
- `DELETE /todo/{id}` - Delete a todo

### Users
- `POST /user` - Register
- `POST /login` - Login (returns JWT)
