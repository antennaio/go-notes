# go-notes

A sample API implemented on top of [go-chi](https://github.com/go-chi/chi) and [go-pg](https://github.com/go-pg/pg).

## Features

- modular code organization
- examples of standard CRUD operations
- environment dependent configuration (`.env` file)
- authentication using JSON Web Tokens
- request validation
- query and request logs
- route/struct binding using middleware
- autogenerated slugs using BeforeInsert/BeforeUpdate hooks
- soft deletes
- migrations
- modd - recompiles and runs the `api` package in response to filesystem changes

## Running migrations

`go run migrations/*.go`

## Starting the server

`go run api/*.go` (or `modd`)

## Sample requests

**Unauthorized request**

`curl http://localhost:8080/v1/note/`

```
{
  "status": "Unauthorized",
  "error": "Invalid token."
}
```

**Register a user account**

`curl --data '{"email":"test@test.com","password":"12345678","first_name":"Caroline","last_name":"Dennis"}' http://localhost:8080/auth/register`

```
{
  "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAdGVzdC5jb20iLCJpZCI6N30.01ClhXk_k2E2ozMTeix-oRYhsTsHL2KtGYh3NUBDH4c"
}
```

**Create a note**

`curl --data '{"Title":"This is a note","Content":"Content..."}' --header 'Authorization: BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAdGVzdC5jb20iLCJpZCI6N30.01ClhXk_k2E2ozMTeix-oRYhsTsHL2KtGYh3NUBDH4c' http://localhost:8080/v1/note/`

```
{
  "id": 1,
  "slug": "this-is-a-note",
  "title": "This is a note",
  "content": "Content...",
  "updated_at": "2019-09-03T09:59:48.516568+02:00"
}
```

**Update a note**

`curl -X PUT --data '{"Title":"This is an updated note","Content":"Content..."}' --header 'Authorization: BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAdGVzdC5jb20iLCJpZCI6N30.01ClhXk_k2E2ozMTeix-oRYhsTsHL2KtGYh3NUBDH4c' http://localhost:8080/v1/note/1`

```
{
  "id": 1,
  "slug": "this-is-an-updated-note",
  "title": "This is an updated note",
  "content": "Content...",
  "updated_at": "2019-09-03T10:01:12.625071+02:00"
}
```

**List notes**

`curl --header 'Authorization: BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAdGVzdC5jb20iLCJpZCI6N30.01ClhXk_k2E2ozMTeix-oRYhsTsHL2KtGYh3NUBDH4c' http://localhost:8080/v1/note/`

```
[
  {
    "id": 1,
    "slug": "this-is-an-updated-note",
    "title": "This is an updated note",
    "content": "Content...",
    "updated_at": "2019-09-03T10:01:12.625071+02:00"
  }
]
```

**Delete a note**

`curl -X DELETE --header 'Authorization: BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAdGVzdC5jb20iLCJpZCI6N30.01ClhXk_k2E2ozMTeix-oRYhsTsHL2KtGYh3NUBDH4c' http://localhost:8080/v1/note/1`

```
{
  "status": "Success"
}
```
