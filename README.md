# 📞 Phonebook Application with Go

This project is a simple phonebook application built using the Go programming language. It leverages the powerful [Fiber](https://gofiber.io/) web framework for handling HTTP requests. Data is stored in a MySQL database, and interaction with the database is managed via the popular [GORM](https://gorm.io/) ORM. Environment variables are handled using the `godotenv` package.

## 🚀 Features

- Create, Read, Update, and Delete contacts
- Secure connection to MySQL database
- Clean and organized code structure with Fiber and GORM
- Configuration management via `.env` file

## 🔧 Technologies Used

| Technology    | Description                                   |
|---------------|-----------------------------------------------|
| Go            | Main programming language                     |
| Fiber         | Lightweight and fast HTTP framework           |
| MySQL         | Database for storing contact information      |
| GORM          | ORM for database interactions                 |
| godotenv      | Manages environment variables via `.env` file |


## 🛠 How to Run

1. Create a `.env` file with your database configuration:


2. Run the project using the following command:

```bash
go run main.go


|Method	|Endpoint	|Description                     |
|-------|-----------|---------------------------|
| GET	    | /users	    | Retrieve all contacts      |
| Get       | /users/:id    | get user by id             |
| POST	    | /users	    | Add a new contact          |
| PUT	    | /users/:id	| Update contact by ID       |
| DELETE    | /users/:id	| Delete contact by ID       |