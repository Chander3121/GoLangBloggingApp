# 📚 Go Blog API

A simple and clean **blogging REST API** built using:

- **Go (Golang)**
- **Gin Web Framework**
- **PostgreSQL**
- **GORM ORM**
- **Environment variables (.env)**

This project includes full **CRUD operations** for blog posts.

---

## 🚀 Features

- Create a blog post  
- Get all blog posts  
- Get a single blog post  
- Update a blog post  
- Delete a blog post  
- Auto database migration  
- PostgreSQL connection via `.env`  
- Modular clean folder structure  

---

# 📁 Project Structure

```
go-blog/
│
├── main.go
├── README.md
├── .env
│
├── config/
│   └── db.go
│
├── models/
│   └── blog.go
│
├── controllers/
│   └── blog_controller.go
│
└── routes/
    └── blog_routes.go
```

---

# ⚙️ Setup Instructions

## 1️⃣ Clone the repository
```bash
git clone https://github.com/yourusername/go-blog.git
cd go-blog
```

## 2️⃣ Install Go dependencies
```bash
go mod tidy
```

## 3️⃣ Create PostgreSQL Database
Open PostgreSQL:

```bash
psql -U postgres
```

Create database:

```sql
CREATE DATABASE blogdb;
```

Exit:

```sql
\q
```

---

# 🔐 4️⃣ Create `.env` file

Create a `.env` in the project root:

```
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=blogdb
DB_PORT=5432
DB_SSLMODE=disable
```

---

# ▶️ 5️⃣ Run the Application

```bash
go run main.go
```

You should see:

```
Database connected successfully!
Listening and serving HTTP on :8080
```

---

# 🛣️ API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/blogs` | Create blog |
| GET | `/blogs` | Get all blogs |
| GET | `/blogs/:id` | Get single blog |
| PUT | `/blogs/:id` | Update blog |
| DELETE | `/blogs/:id` | Delete blog |

---

# 📌 Example Requests

### ➤ Create Blog
**POST /blogs**
```json
{
  "title": "My First Blog",
  "content": "This is the content of my blog"
}
```

### ➤ Update Blog
**PUT /blogs/1**
```json
{
  "title": "Updated Title",
  "content": "Updated content"
}
```

---

# 🛠️ Technologies Used

- Go 1.21+
- Gin Web Framework
- GORM ORM
- PostgreSQL
- Godotenv

---

# 🚀 Future Improvements (Optional)

- Add JWT authentication  
- Add users and comments  
- Add pagination  
- Add categories & tags  
- Add soft deletes  
- Add Docker support  
- Add validation layer
