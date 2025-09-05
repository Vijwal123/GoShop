  

---

## 🚀 Features

- **User Authentication** – Register & login with hashed passwords and JWT-based authentication  
- **Role-Based Authorization** – Roles: `user` and `admin` with protected routes  
- **Product Management** – Add, view products (only admins can create)  
- **Shopping Cart** – Add, view, and remove items from the cart  
- **Order Processing** – Checkout system with order history  
- **Middleware** – JWT middleware for securing API endpoints  
- **PostgreSQL** – Reliable relational database for persistence  
- **Modular Codebase** – Separated into `models`, `handlers`, `middleware`, and `database`  

---

## 🛠️ Tech Stack

| Component      | Technology   |
|----------------|--------------|
| Language       | Go (Golang)  |
| Web Framework  | Fiber        |
| ORM            | GORM         |
| Database       | PostgreSQL   |
| Auth           | JWT          |
| Password Hash  | bcrypt       |

---

## 📁 Project Structure

goshop/
├── go.mod
├── main.go
├── .env.example
├── database/
│ └── connect.go
├── models/
│ ├── user.go
│ ├── product.go
│ ├── cart.go
│ └── order.go
├── handlers/
│ ├── user_handler.go
│ ├── auth_handler.go
│ ├── product_handler.go
│ ├── cart_handler.go
│ └── order_handler.go
└── middleware/
└── auth.go 

