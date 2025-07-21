# Structs to JSON: Go REST API Example

Welcome! This is a practical Go REST API for user management, demonstrating best practices in Go for JSON handling, code organization, validation, PostgreSQL integration, and error responses. Feel free to use this as a learning project, or as an example of idiomatic Go.

---

## Features

- Full CRUD for users (Create, Read, Update, Delete)
- PostgreSQL storage (no data loss on restart)
- Clean project structure, easy to extend
- Input validation for safer data
- Consistent, clear error handling
- Health check endpoint for monitoring
- Designed for production-readiness (timeouts, graceful shutdown)

---

## Getting Started

**Requirements:**  
- Go (1.21 +)
- PostgreSQL (local or remote)

### 1. Prepare Your Database

Create a PostgreSQL database and user for this app.  
**Tip:** Use a strong password and don't share it in public code!

1. Create your database and user (in psql or DBeaver, etc):
    ```sql
    CREATE DATABASE structs_demo;
    CREATE USER demo_user WITH PASSWORD 'YOUR_STRONG_PASSWORD';
    GRANT ALL PRIVILEGES ON DATABASE structs_demo TO demo_user;
    ```

2. Run the migration to set up the `users` table:
    ```sql
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(64) NOT NULL UNIQUE,
        email VARCHAR(128) NOT NULL UNIQUE,
        active BOOLEAN NOT NULL DEFAULT TRUE,
        created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
    );
    ```

### 2. Configure Environment

Set the `DATABASE_URL` environment variable for your connection.  
**Do not put real passwords in public repos or files.**

Example for your terminal:
```sh
export DATABASE_URL="postgres://demo_user:YOUR_STRONG_PASSWORD@localhost:5432/structs_demo?sslmode=disable"
```

### 3. Build and Run

Install dependencies and start the server:
```sh
go mod tidy
go run cmd/server/main.go
```
The server will listen on [http://localhost:8080](http://localhost:8080).

---

## API Overview

| Method | Path           | Description           |
|--------|----------------|----------------------|
| GET    | `/users`       | List all users       |
| POST   | `/users`       | Create a user        |
| GET    | `/users/{id}`  | Get user by ID       |
| PUT    | `/users/{id}`  | Update user by ID    |
| DELETE | `/users/{id}`  | Delete user by ID    |
| GET    | `/healthz`     | Health check         |

### Example: Creating and Listing Users

**Create:**
```sh
curl -X POST -H "Content-Type: application/json" \
-d '{"username":"alice","email":"alice@example.com","active":true}' \
http://localhost:8080/users
```

**List:**
```sh
curl http://localhost:8080/users
```

---

## Adapting and Extending

- Want to use a different database? Implement the `UserStore` interface for your backend.
- Add authentication, logging, or more resource types as needed.
- Integrate with a frontend or deploy as a microservice.

---

## Architecture Diagram

You can see `diagram.png` 

---

## License

MIT License — see [LICENSE](LICENSE) for details.

---