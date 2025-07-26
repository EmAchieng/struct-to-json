# Structs to JSON: Go REST API
[![Go Version](https://img.shields.io/github/go-mod/go-version/EmAchieng/struct-to-json)](https://golang.org/doc/go1.24)
[![Last Commit](https://img.shields.io/github/last-commit/EmAchieng/struct-to-json)](https://github.com/EmAchieng/struct-to-json/commits)
[![Issues](https://img.shields.io/github/issues/EmAchieng/struct-to-json)](https://github.com/EmAchieng/struct-to-json/issues)
[![Pull Requests](https://img.shields.io/github/issues-pr/EmAchieng/struct-to-json)](https://github.com/EmAchieng/struct-to-json/pulls)
[![Code Style: gofmt](https://img.shields.io/badge/code%20style-gofmt-brightgreen.svg)](https://golang.org/doc/go1.21#go-fmt)
[![License](https://img.shields.io/github/license/EmAchieng/struct-to-json)](./LICENSE)
[![Contributors](https://img.shields.io/github/contributors/EmAchieng/struct-to-json)](https://github.com/EmAchieng/struct-to-json/graphs/contributors)

Welcome! This project is a practical demonstration of a minimal Go REST API for user management. It showcases best practices in:

- JSON handling with Go structs  
- Clean and modular code organization  
- Input validation  
- PostgreSQL integration  
- Standardized error responses  

This demo has been created for a talk at the **First Online Go Language Conference for the South African Community**.

---

## Struct-to-JSON Flow

![Struct-to-JSON Flow Diagram](Struct-to-JSON.png)

This diagram illustrates the HTTP handler flow in Go:

1. **HTTP Request** – A request is received by the server (e.g. `POST /users`).
2. **Router** – The request is routed using a multiplexer.
3. **Handler Function** – The request is handled by a function.
4. **Decode JSON** – The handler decodes the request body into a Go struct using `json.NewDecoder(r.Body).Decode(&input)`.
5. **Create or Prepare Struct** – The handler may directly use the input or prepare a new response struct.
6. **Returns struct/data** – A struct is prepared for response.
7. **Encode JSON Response** – The struct is encoded to JSON using `json.NewEncoder(w).Encode(user)`.
8. **HTTP JSON Response Sent** – JSON is sent back to the client as the response.

📌  _The above can be adapted and extended by service or business logic.—._


---

## Features

- CRUD for users (Create, Read, Update, Delete)
- PostgreSQL storage 
- Clean project structure, easy to extend
- Input validation for safer data
- Consistent, clear error handling
---

## Getting Started

**Requirements:**  
- Go (1.21 +)
- PostgreSQL (local or remote)

**Get Started:** 
```sh
git clone https://github.com/EmAchieng/struct-to-json.git
```
### 1. Prepare Your Database

Create a PostgreSQL database and user for this app.  

1. Create your database and user (in psql or DBeaver, etc):
    ```sql
    CREATE DATABASE structs_demo;
    CREATE USER demo_user WITH PASSWORD '$DB_PASSWORD';  -- Password stored securely elsewhere
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

Example for your terminal:
```sh
export DATABASE_URL="postgres://demo_user:$DB_PASSWORD@localhost:5432/structs_demo?sslmode=disable"
```
If your database is external, use the external host or IP:

```
DATABASE_URL=postgres://demo_user:yourpassword@<EXTERNAL_HOST>:5432/structs_demo?sslmode=disable
```

Replace `<EXTERNAL_HOST>` with your external database server’s address (e.g., `db.example.com` or an IP).

> Make sure your firewall and database settings allow remote connections.

### 3a. Build and Run Locally

Install dependencies and start the server:
```sh
go mod tidy
go run cmd/server/main.go
```
The server will listen on [http://localhost:8080](http://localhost:8080).

### 3b. Build and Run With Docker

Install dependencies and start the server:
```sh
docker build -t structs-to-json .
```

```sh
docker run --rm -p 8080:8080 \
  -e DATABASE_URL="postgres://demo_user:${DB_PASSWORD}@host.docker.internal:5432/structs_demo?sslmode=disable" \
  structs-to-json
```
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
-d '{"username":"gopher","email":"gopher@example.com","active":true}' \
http://localhost:8080/users
```

**List:**
```sh
curl http://localhost:8080/users
```

- Production Tip: Always use HTTPS for API requests in production.
---

## Adapting and Extending

- Want to use a different database? Implement the `UserStore` interface for your backend.
- Add authentication, logging, or more resource types as needed.
- Integrate with a frontend or deploy as a microservice.

- Maintenance Tip: Regularly update dependencies to receive security patches.
- Extend API endpoints to support pagination, filtering, or advanced queries.
---

## License

MIT License — see [LICENSE](LICENSE) for details.

---