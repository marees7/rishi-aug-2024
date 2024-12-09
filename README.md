# rishi-aug-2024

# blog post


REST API for blog using golang, echo framework & GORM ORM.


## Overview
- There are 2 different roles, User and admin
- A user can create posts, add comments, edit posts & comments and delete post & comments (User can only update/delete their own posts or comments)
- A admin manages the total posts & total users and can do every operations user can do (Admin can update/delete any posts or comments)


## Features
- User authentication and authorization using JSON Web Tokens (JWT)
- CRUD operations for blog posts
- Pagination and sorting of blog posts
- Error handling and response formatting
- Input validation and data sanitization
- Database integration using PostgreSQL


## Requirements
- Golang 
- Postgres


## Run Locally

Clone the project

```bash
  git clone https://github.com/marees7/rishi-aug-2024.git
```

Go to the project directory
go to the cmd folder and main.go file.
change the credentials of postgres db in the internals.

```bash
  go run main.go
```


## API Endpoints

The following endpoints are available in the API:

## AUTH API

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |
| POST |	/signup	| Register a new user |
| POST |	/login	| Log in and obtain JWT |

## USER API

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |
| GET  |	/v1/users/	| Get all users |
| GET  |	/v1/users/:username	| Get a specific user |
| POST |	/v1/users/posts	| Create a new blog post |
| GET  |	/v1/users/posts	| Get all or specific blog post |
| PUT  |	/v1/users/posts/:post_id	| Update a specific blog post |
| DELETE |	/v1/users/posts/:post_id	| Delete a specific blog post |
| POST |	/v1/users/comments/:post_id	| Create a new comment |
| PUT  |	/v1/users/comments/:post_id	| Update a specific blog post comment |
| DELETE |	/v1/users/comments/:post_id	| Delete a specific blog post comment |

## ADMIN API

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |
| GET  |	/v1/admin/	| Get all users |
| GET  |	/v1/admin/:username	| Get a specific user |
| POST |	/v1/admin/categories	| Create a new category |
| PUT  |	/v1/admin/categories/:category_id	| Update a specific blog post category |
| DELETE |	/v1/admin/categories/:category_id	| Delete a specific blog post category |


## Database Schema

The application uses a PostgreSQL database with the following schema:

```sql
CREATE TABLE IF NOT EXISTS users
(
    user_id BIGINT SERIAL PRIMARY KEY,
    email TEXT NULL,
    username TEXT NULL,
    password TEXT NULL,
    role TEXT ,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
);

CREATE TABLE IF NOT EXISTS posts (
    post_id BIGINT SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    description TEXT,
    user_id BIGINT,
    category_id BIGINT,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
);

CREATE TABLE IF NOT EXISTS comments (
    comment_id BIGINT SERIAL PRIMARY KEY,
    content TEXT NOT NULL ::text,
    user_id BIGINT,
    post_id BIGINT,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
);

CREATE TABLE IF NOT EXISTS categories (
    category_id BIGINT SERIAL PRIMARY KEY,
    category_name TEXT NOT NULL,
    description TEXT,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
);
```

## Dependencies

The project utilizes the following third-party libraries:

- `labstack/echo/v4`: HTTP web framework
- `joho/godotenv`: Environment variable loading
- `golang-jwt/jwt/v5`: JWT implementation
- `gorm.io/gorm`: PostgreSQL ORM
- `gorm.io/driver/postgres`: PostgreSQL extensions for gorm

Make sure to run go mod download as mentioned in the installation steps to fetch these dependencies.

