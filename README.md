# 🚵 Ojeg - Backend API

**Ojeg** is a backend system for a ride-hailing service built with **Golang**, using a clean architecture pattern. It supports multi-database (MySQL, PostgreSQL, SQLite), JWT-based authentication, and modular service architecture.

---

## 🚀 Features

* User Registration & Login (JWT-based)
* Secure password hashing with bcrypt
* User CRUD (Create, Read, Update, Delete)
* Custom error handling with code
* JSON standardized responses
* Database migration system
* Configurable via `.env` and `config.yaml`
* Clean, modular folder structure

---

## 📁 Project Structure

```
ojeg/
├── cmd/                    # Application entry point (main.go)
├── configs/                # App configuration (YAML + .env)
│   ├── config.yaml
│   └── config.go
├── delivery/
│   └── http/
│       ├── handler/        # HTTP handlers (controllers)
│       └── router.go       # HTTP router setup
├── infrastructure/
│   ├── db/                 # DB connection & GORM setup
│   └── jwt/                # JWT implementation
├── internal/
│   ├── domain/             # Domain models and DTOs
│   ├── repository/         # Repository interfaces & implementations
│   ├── service/            # Business logic
│   └── usecase/            # Usecases (contract for services)
├── migrations/             # DB migration files
├── pkg/
│   ├── errors/             # Application error codes
│   └── response/           # Response formatter
├── registry/               # Dependency injection wiring
├── bootstrap/              # DB & config initializer
└── go.mod
```

---

## ⚙️ Configuration

### 1. `.env` File

Create a `.env` file in the project root:

```env
APP_ENV=development
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=toor
DB_NAME=ojeg
DB_SSLMODE=disable
JWT_SECRET=supersecretjwtkey
JWT_EXPIRE_HOURS=72
```

### 2. `config.yaml`

```yaml
default:
  app_name: "Ojeg Default"
  db:
    driver: mysql
    host: localhost
    port: 3306
    user: root
    password: toor
    name: ojeg
    sslmode: disable

development:
  app_name: "Ojeg Dev"
  db:
    driver: mysql
    host: localhost
    port: 3306
    user: dev_user
    password: dev_password
    name: ojeg_dev
    sslmode: disable
```

---

## 📦 Setup & Usage

### Run the Server

```bash
go run cmd/main.go
```

### Run Migrations

```bash
go run cmd/main.go migrate
go run cmd/main.go migrate:rollback
go run cmd/main.go migrate:refresh
```

### Create Migration File

```bash
go run cmd/main.go make:migration CreateUserTable
```

---

## 🧪 Example API Requests

### Register

```
POST /api/v1/auth/register
```

```json
{
  "username": "johndoe",
  "name": "John Doe",
  "email": "john@example.com",
  "password": "strongpassword"
}
```

### Login

```
POST /api/v1/auth/login
```

```json
{
  "username": "johndoe",
  "password": "strongpassword"
}
```

### Get User

```
GET /api/v1/users/{id}
Authorization: Bearer <your_jwt_token>
```

---

## 📌 Error Handling Example

All responses are wrapped in a standard structure:

```json
{
  "status": "error",
  "code": 701,
  "message": "Invalid payload"
}
```

---

## ✅ TODO (Next Features)

* Booking and Driver modules
* Rate limiting / throttle
* Swagger/OpenAPI docs
* Role-based authorization
* Unit & integration testing

---

## 📝 License

MIT License © 2025 - Built with 💙 by \Hoeril Ahyar

```
```
