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
| GET  |	/v1/admin/users/	| Get all users |
| GET  |	/v1/admin/users/:username	| Get a specific user |
| PUT  |	/v1/users	| Update the logged in user details |
| DELETE |	/v1/users	| Delete the logged in user details |


## POST API

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |
| POST |	/v1/users/posts	| Create a new blog post |
| GET  |	/v1/users/posts	| Get all or specific blog post |
| PUT  |	/v1/users/posts/:post_id	| Update a specific blog post |
| DELETE |	/v1/users/posts/:post_id	| Delete a specific blog post |


## COMMENT API

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |
| POST |	/v1/users/comments/:post_id	| Create a new comment |
| GET  |	/v1/users/comments/:post_id	| Get comments of the specific blog post |
| PUT  |	/v1/users/comments/:comment_id	| Update a specific blog post comment |
| DELETE |	/v1/users/comments/:comment_id	| Delete a specific blog post comment |


## REPLY API

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |
| POST |	v1/users/reply/:comment_id	| Create a new comment |
| PUT  |	v1/users/reply/:reply_id	| Update a specific blog post comment |
| DELETE |	v1/users/reply:reply_id	| Delete a specific blog post comment |



## CATEGORY API

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |
| POST |	/v1/admin/categories	| Create a new category |
| GET  |	/v1/users/categories	| Get all available category |
| PUT  |	/v1/admin/categories/:category_id	| Update a specific blog post category |
| DELETE |	/v1/admin/categories/:category_id	| Delete a specific blog post category |


## Database Schema

The application uses PostgreSQL database with the following schema:

```sql
CREATE TABLE IF NOT EXISTS users
(
    user_id UUID PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role TEXT,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
);

CREATE TABLE IF NOT EXISTS posts (
    post_id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    description TEXT,
    user_id UUID FOREIGN KEY,
    category_id UUID FOREIGN KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
);

CREATE TABLE IF NOT EXISTS comments (
    comment_id UUID PRIMARY KEY,
    content TEXT NOT NULL,
    user_id UUID FOREIGN KEY,
    post_id UUID FOREIGN KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
);

CREATE TABLE IF NOT EXISTS public.replies
(
    reply_id UUID PRIMARY KEY,
    content TEXT NOT NULL,
    user_id UUID FOREIGN KEY,
    comment_id UUID FOREIGN KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
);

CREATE TABLE IF NOT EXISTS categories (
    category_id UUID PRIMARY KEY,
    category_name TEXT UNIQUE NOT NULL,
    description TEXT,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
);
```

## Sample API Requests and Responses

##### POST /signup

sample request:

```json
{
    "email":"rsi28c@gmail.com",
    "username":"rishi.k",
    "name":"rishi",
    "password":"password",
    "role":"admin"
}
```

sample response:

```json
{
    "message": "User created successfully"
}
```

##### POST /login

sample request:

```json
{
    "email":"rsi28c@gmail.com",
    "password":"password"
}
```

sample response:

```json
{
    "message": "Logged in successfully"
}
```


##### PUT v1/user

sample request:

```json
{
    "username":"rishi.k",
    "name":"rishi",
    "password":"newpassword",
    "role":"admin"
}
```

sample response:

```json
{
    "message": "user details updated successfully",
    "data": "rsi28c@gmail.com"
}
```

##### DELETE v1/user

this will delete user if logged in and response back the deleted email

sample response:
```json
{
    "message": "user deleted successfully",
    "data": "varusai21@gmail.com"
}
```



##### POST v1/users/post

sample request:

```json
{
    "title": "My first blog post",
    "content": "This is my first blog post i'm posting here",
    "description": "This is about my first blog",
    "category_id": "4fcdc14f-9545-4236-a88c-ec7c3c60ca4e"
}
```

sample response:

```json
{
    "message": "post created successfully",
    "data": {
        "post_id": "e335b188-810d-4685-bb47-83066048461e",
        "title": "My first blog post",
        "content": "This is my first blog post i'm posting here",
        "description": "This is about my first blog",
        "user_id": "5e3136c2-895a-40d3-a33c-1773b7ddd504",
        "category_id": "4fcdc14f-9545-4236-a88c-ec7c3c60ca4e",
        "created_at": "2024-12-20T12:01:34.219095+05:30",
        "updated_at": "2024-12-20T12:01:34.219095+05:30"
    }
}
```


##### GET v1/users/post/:post_id

sample response:

```json
{
    "message": "Post retrieved successfully",
    "data": {
        "post_id": "e335b188-810d-4685-bb47-83066048461e",
        "title": "My first blog post",
        "content": "This is my first blog post i'm updating here",
        "description": "This is about my first blog",
        "user_id": "5e3136c2-895a-40d3-a33c-1773b7ddd504",
        "category_id": "4fcdc14f-9545-4236-a88c-ec7c3c60ca4e",
        "comments": [
            {
                "comment_id": "43ec1c2d-50e0-4f59-a7ff-3923a72b084e",
                "content": "this comment is made by the user",
                "user_id": "5e3136c2-895a-40d3-a33c-1773b7ddd504",
                "post_id": "e335b188-810d-4685-bb47-83066048461e",
                "created_at": "2024-12-20T12:21:54.561246+05:30",
                "updated_at": "2024-12-20T12:21:54.561246+05:30"
            },
            {
                "comment_id": "508644df-699f-41dc-8441-9fdfad285814",
                "content": "updated by admin",
                "user_id": "5e3136c2-895a-40d3-a33c-1773b7ddd504",
                "post_id": "e335b188-810d-4685-bb47-83066048461e",
                "created_at": "2024-12-20T12:21:57.350714+05:30",
                "updated_at": "2024-12-20T12:22:14.181375+05:30"
            }
        ],
        "created_at": "2024-12-20T12:01:34.219095+05:30",
        "updated_at": "2024-12-20T12:20:40.620055+05:30"
    }
}
```

##### GET v1/users/post

sample response:

```json
{
    "message": "Posts retrieved successfully",
    "data": [
        {
            "post_id": "e335b188-810d-4685-bb47-83066048461e",
            "title": "My first blog post",
            "content": "This is my first blog post i'm updating here",
            "description": "This is about my first blog",
            "user_id": "5e3136c2-895a-40d3-a33c-1773b7ddd504",
            "category_id": "4fcdc14f-9545-4236-a88c-ec7c3c60ca4e",
            "comments": [
                {
                    "comment_id": "43ec1c2d-50e0-4f59-a7ff-3923a72b084e",
                    "content": "this comment is made by the user",
                    "user_id": "5e3136c2-895a-40d3-a33c-1773b7ddd504",
                    "post_id": "e335b188-810d-4685-bb47-83066048461e",
                    "created_at": "2024-12-20T12:21:54.561246+05:30",
                    "updated_at": "2024-12-20T12:21:54.561246+05:30"
                },
                {
                    "comment_id": "508644df-699f-41dc-8441-9fdfad285814",
                    "content": "updated by admin",
                    "user_id": "5e3136c2-895a-40d3-a33c-1773b7ddd504",
                    "post_id": "e335b188-810d-4685-bb47-83066048461e",
                    "created_at": "2024-12-20T12:21:57.350714+05:30",
                    "updated_at": "2024-12-20T12:22:14.181375+05:30"
                }
            ],
            "created_at": "2024-12-20T12:01:34.219095+05:30",
            "updated_at": "2024-12-20T12:20:40.620055+05:30"
        }
    ],
    "limit": 10,
    "total_records": 8
}
```

##### PUT v1/users/post/:post_id

sample request:

```json
{
    "title":"My first blog post",
    "content":"This is my first blog post i'm updating here (modified)",
    "description":"This is about my first blog"
}
```

sample response:

```json
{
    "message": "Post updated successfully",
    "data": "e335b188-810d-4685-bb47-83066048461e"
}
```

##### DELETE v1/users/post/:post_id

this will delete post with given post id and response back the deleted post id

sample response:
```json
{
    "message": "Post deleted successfully",
    "data": "8936154e-40a2-4055-ad2d-7f2b83d8d4e4"
}
```


##### POST v1/admin/categories

sample request:

```json
{
    "category_name":"entertainment blogs",
    "description":"entertainment blogs can be stored here"
}
```

sample response:

```json
{
    "message": "Category created successfully",
    "data": {
        "category_id": "8d453de4-6c54-4f3f-966b-ccaa8552fc7a",
        "category_name": "entertainment blogs",
        "description": "entertainment blogs can be stored here",
        "created_at": "2024-12-20T17:34:39.9220777+05:30",
        "updated_at": "2024-12-20T17:34:39.9220777+05:30"
    }
}
```


##### GET v1/users/categories

sample response:

```json
{
    "message": "retrieved categories successfully",
    "data": [
        {
            "category_id": "9aaa11e3-6661-43c6-9a1f-ca1af4878e84",
            "category_name": "some blogs",
            "description": "some blogs can be stored here",
            "created_at": "2024-12-20T12:01:00.54411+05:30",
            "updated_at": "2024-12-20T12:01:00.54411+05:30"
        },
        {
            "category_id": "4fcdc14f-9545-4236-a88c-ec7c3c60ca4e",
            "category_name": "new blogs",
            "description": "new blogs can be stored here",
            "created_at": "2024-12-20T12:01:14.172411+05:30",
            "updated_at": "2024-12-20T12:01:14.172411+05:30"
        },
        {
            "category_id": "8d453de4-6c54-4f3f-966b-ccaa8552fc7a",
            "category_name": "entertainment blogs",
            "description": "entertainment blogs can be stored here",
            "created_at": "2024-12-20T17:34:39.922077+05:30",
            "updated_at": "2024-12-20T17:34:39.922077+05:30"
        }
    ],
    "limit": 10,
    "total_records": 3
}
```


##### PUT v1/users/categories/:category_id

sample request:

```json
{
    "category_name":"common blogs",
    "description":"delete blogs can be stored here"
}
```

sample response:

```json
{
    "message": "Category updated successfully",
    "data": {
        "category_id": "9aaa11e3-6661-43c6-9a1f-ca1af4878e84"
    }
}
```

##### DELETE v1/users/categories/:category_id

this will delete category with given category id and response back the deleted category id

sample response:
```json
{
    "message": "Category deleted successfully",
    "data": {
        "category_id": "9aaa11e3-6661-43c6-9a1f-ca1af4878e84"
    }
}
```

##### POST v1/users/comment/:post_id

sample request:

```json
{
    "content":"this comment is made by me"
}
```

sample response:

```json
{
    "message": "comment added successfully",
    "data": {
        "comment_id": "eb9d1044-f3a7-444b-96b9-622c61966c9c",
        "content": "this comment is made by me",
        "user_id": "26538d56-2638-4dc5-af22-8ce45ef2556b",
        "post_id": "e335b188-810d-4685-bb47-83066048461e",
        "created_at": "2024-12-20T18:40:29.1256198+05:30",
        "updated_at": "2024-12-20T18:40:29.1256198+05:30"
    }
}
```


##### GET v1/users/comment/:post_id

sample response:

```json
{
    "message": "Comments retrieved successfully",
    "data": [
        {
            "comment_id": "43ec1c2d-50e0-4f59-a7ff-3923a72b084e",
            "content": "this comment is made by the user",
            "user_id": "5e3136c2-895a-40d3-a33c-1773b7ddd504",
            "post_id": "e335b188-810d-4685-bb47-83066048461e",
            "created_at": "2024-12-20T12:21:54.561246+05:30",
            "updated_at": "2024-12-20T12:21:54.561246+05:30"
        },
        {
            "comment_id": "508644df-699f-41dc-8441-9fdfad285814",
            "content": "updated by admin",
            "user_id": "5e3136c2-895a-40d3-a33c-1773b7ddd504",
            "post_id": "e335b188-810d-4685-bb47-83066048461e",
            "replies": [
                {
                    "reply_id": "fd1e4fd3-8eac-4d43-a09b-ea42753b4640",
                    "content": "this reply is made by the admin",
                    "user_id": "5e3136c2-895a-40d3-a33c-1773b7ddd504",
                    "comment_id": "508644df-699f-41dc-8441-9fdfad285814",
                    "created_at": "2024-12-20T12:22:32.004315+05:30",
                    "updated_at": "2024-12-20T12:22:32.004315+05:30"
                }
            ],
            "created_at": "2024-12-20T12:21:57.350714+05:30",
            "updated_at": "2024-12-20T12:22:14.181375+05:30"
        },
        {
            "comment_id": "eb9d1044-f3a7-444b-96b9-622c61966c9c",
            "content": "this comment is made by me",
            "user_id": "26538d56-2638-4dc5-af22-8ce45ef2556b",
            "post_id": "e335b188-810d-4685-bb47-83066048461e",
            "created_at": "2024-12-20T18:40:29.125619+05:30",
            "updated_at": "2024-12-20T18:40:29.125619+05:30"
        }
    ],
    "limit": 10,
    "total_records": 3
}
```


##### PUT v1/users/comment/:comment_id

sample request:

```json
{
    "content":"updated by me"
}
```

sample response:

```json
{
    "message": "comment edited successfully",
    "data": {
        "comment_id": "508644df-699f-41dc-8441-9fdfad285814"
    }
}
```

##### DELETE v1/users/comment/:comment_id

this will delete comment with given comment id and response back the deleted comment id

sample response:
```json
{
    "message": "comment deleted successfully",
    "data": "43ec1c2d-50e0-4f59-a7ff-3923a72b084e"
}
```

##### POST v1/users/reply/:comment_id

sample request:

```json
{
    "content":"this reply is made by the other user"
}
```

sample response:

```json
{
    "message": "reply added to the comment successfully",
    "data": {
        "reply_id": "f3d642ff-1d8e-43bd-80d1-852480f77daa",
        "content": "this reply is made by the other user",
        "user_id": "5e3136c2-895a-40d3-a33c-1773b7ddd504",
        "comment_id": "eb9d1044-f3a7-444b-96b9-622c61966c9c",
        "created_at": "2024-12-20T18:45:49.6367464+05:30",
        "updated_at": "2024-12-20T18:45:49.6367464+05:30"
    }
}
```


##### PUT v1/users/reply/:reply_id

sample request:

```json
{
    "content":"updated reply"
}
```

sample response:

```json
{
    "message": "reply edited successfully",
    "data": {
        "reply_id": "f3d642ff-1d8e-43bd-80d1-852480f77daa",
    }
}
```

##### DELETE v1/users/reply/:reply_id

this will delete reply with given reply id and response back the deleted reply id

sample response:
```json
{
    "message": "reply deleted successfully",
    "data": "fd1e4fd3-8eac-4d43-a09b-ea42753b4640"
}
```


## Dependencies

The project utilizes the following third-party libraries:

- `labstack/echo/v4`: HTTP web framework
- `joho/godotenv`: Environment variable loading
- `golang-jwt/jwt/v5`: JWT implementation
- `gorm.io/gorm`: PostgreSQL ORM
- `gorm.io/driver/postgres`: PostgreSQL extensions for gorm
- `google/uuid`: UUID generation
- `go-playground/validator/v10`: Struct field validation

Make sure to run go mod download as mentioned in the installation steps to fetch these dependencies.

