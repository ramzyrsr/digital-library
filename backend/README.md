# 📚 Digital Library Analytics Dashboard

A Golang (Fiber) & PostgreSQL-based digital library system that tracks book lending patterns and reader engagement. Provides insights through analytics APIs, allowing librarians to manage books and lending records efficiently.

## 🚀 Features
- ✅ Books Management (CRUD, Pagination, Filtering)
- ✅ Lending API (Borrow, Return, Transactions)
- ✅ Analytics API (Lending trends, Most borrowed books, Active users)
- ✅ JWT Authentication (User registration, login, protected routes)
- ✅ Error Handling & Middleware

## 🏗️ Project Structure
```sh
digital-library/
│── cmd/                    
│   └── main.go             # Entry point of the application
│
│── config/                 
│   ├── database.go         # PostgreSQL connection setup
│
│── internal/
│   ├── entity/             
│   │   ├── book.go         # Book entity
│   │   ├── lending.go      # Lending entity
│   │   ├── user.go         # User entity
│   │   ├── category.go     # Category entity
│   │
│   ├── repository/         
│   │   ├── book_repository.go    # Book repository with raw SQL
│   │   ├── lending_repository.go # Lending repository
│   │   ├── user_repository.go    # User repository (for auth)
│   │   ├── analytics_repository.go # Analytics repository
│   │
│   ├── handler/           
│   │   ├── book_handler.go       # Book API handlers
│   │   ├── lending_handler.go    # Lending API handlers
│   │   ├── auth_handler.go       # Authentication handlers
│   │   ├── analytics_handler.go  # Analytics API handlers
│   │
│   ├── middleware/         
│   │   ├── jwt_middleware.go     # JWT Authentication middleware
│   │   ├── response_middleware.go # Global response handling
│
│── .env                    # Environment variables
│── go.mod                  # Go module file
│── go.sum                  # Dependency checksums
│── README.md               # Project setup instructions

```
---

# ⚙️ Setup & Installation
## 1️⃣ Clone the Repository
```sh
git clone https://github.com/yourusername/digital-library.git
cd digital-library
```

## 2️⃣ Create a .env File
```sh
PORT=

DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=

JWT_SECRET=
```

## 3️⃣ Start the API Server
```sh
go run cmd/main.go
```

---

## 👨‍💻 Contributors
- Ramzy Syahrul Ramadhan - [GitHub](https://github.com/ramzyrsr)
- Feel free to contribute by opening issues & pull requests!
---