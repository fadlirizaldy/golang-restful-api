# Golang Restful API

This is golang rest API to do CRUD data

Tools :

- [Golang (Echo Framework)](https://echo.labstack.com/)
- [MySQL](https://dev.mysql.com/doc/)
- [GraphQL (gqlgen)](https://gqlgen.com/)
- [Open API Swagger](https://swagger.io/specification/)

---

## How to Use

This project contains 2 server, API server and GraphQL server for you query data visually.

In the project directory, you can run:

```bash
# install package dependencies
go mod tidy

# Run API server
go run server.go apiserver

# Run GraphQL server
go run server.go gqlserver

# Testing Controller
cd controller
go test -v

# note : don't forget to change .env based on your computer
```

---

## Directory structure

- `/config` - Configuration to connect mysql database.
- `/constants` - Include constants variable.
- `/controller` - Go files used for handle routes function to doing CRUD and testing.
- `/graph` - Contains GraphQL schema, GQL model, and resolver.
- `/helper` - Go files contains function to create some function to hashing password.
- `/middleware` - Middleware to provide some function e.g. logger, trailing slash, and jwt.
- `/model` - Contains model to create tables in the database.
- `/router` - Path to sending some request in this server.
- `/service` - Contains service for GraphQL server.

---

## GraphQL Query

some GraphQL query, you can copy paste to the query field.

```bash
# Query for all movies
query queryAll{
  movies {
    id
    title
    casts {
      birthday
      name
    }
  }
}

# Query For get movie with ID
query queryById{
  movie (id: "1") {
    id
    title
    status
    casts {
      id
      name
      birthday
    }
  }
}

#Query for add new movie
mutation addNewMovie{
  createMovie(input: {
    title: "Arti Bang Messi"
    language: "Jakartans"
    status: "Ongoing"
    rating: 4.7
  }) {
    id
  }
}
```

---

## API Spec

To see the documentation with Swagger, please visit https://editor.swagger.io/ and put code in "spec.json" to it

<i>Note: before try to request, do register and login to get jwt token</i>

Here's the spec :

### Register

Request :

- Method : PUT
- Endpoint : `/register`
- Header :
  - Content-Type: application/json

Body :

```json
{
  "name": "string",
  "email": "string",
  "password": "string"
}
```

Response :

```json
{
  "success": "boolean",
  "message": "string"
}
```

### Login

Request :

- Method : PUT
- Endpoint : `/register`
- Header :
  - Content-Type: application/json

Body :

```json
{
  "email": "string",
  "password": "string"
}
```

Response :

```json
{
  "success": "boolean",
  "user": {
    "id": "integer",
    "name": "string",
    "email": "string",
    "token": "string"
  }
}
```

---

### Get all casts

Request :

- Method : GET
- Endpoint : `/api/v1/casts`
- Header :
  - Accept: application/json

Response :

```json
{
  "Success": "boolean",
  "Message": "string",
  "Data": [
    {
      "id": "integer",
      "name": "string"
    },
    {
      "id": "integer",
      "name": "string"
    }
  ]
}
```

### Get detail cast

Request :

- Method : GET
- Endpoint : `/api/v1/casts/{castID}`
- Header :
  - Accept: application/json

Response :

```json
{
  "id": "integer",
  "name": "string",
  "birth_place": "string",
  "birthday": "datetime",
  "rating": "integer"
}
```

### Create Cast

Request :

- Method : POST
- Endpoint : `/api/v1/casts`
- Header :
  - Content-Type: application/json
  - Accept: application/json

Body :

```json
{
  "name": "string",
  "birth_place": "string",
  "birthday": "datetime",
  "rating": "integer"
}
```

Response :

```json
{
  "Success": "boolean",
  "Message": "string",
  "Data": {
    "id": "integer",
    "name": "string",
    "birth_place": "string",
    "birthday": "datetime",
    "rating": "integer"
  }
}
```

### Update Cast

Request :

- Method : PATCH
- Endpoint : `/api/v1/casts/{castID}`
- Header :
  - Content-Type: application/json
  - Accept: application/json

Body :

```json
{
  "name": "string",
  "birth_place": "string",
  "birthday": "datetime",
  "rating": "integer"
}
```

Response :

```json
{
  "Success": "boolean",
  "Message": "string",
  "Data": {
    "id": "integer",
    "name": "string",
    "birth_place": "string",
    "birthday": "datetime",
    "rating": "integer"
  }
}
```

### Delete Cast

Request :

- Method : DELETE
- Endpoint : `/api/v1/casts/{castID}`
- Header :
  - Accept: application/json

Response :

```json
{
  "messages": "string"
}
```

---

### Get all movies

Request :

- Method : GET
- Endpoint : `/api/v1/movies`
- Header :
  - Accept : application/json

Response :

```json
{
  "Success": "boolean",
  "Message": "string",
  "Data": [
    {
      "id": "integer",
      "title": "string"
    }
  ]
}
```

### Get detail movie

Request :

- Method : GET
- Endpoint : `/api/v1/movies/{movieID}`
- Header :
  - Accept: application/json

Response :

```json
{
  "id": "integer",
  "title": "string",
  "language": "string",
  "status": "string",
  "rating": "number",
  "Casts": [
    {
      "id": "integer",
      "name": "string",
      "birth_place": "string",
      "birthday": "datetime",
      "rating": "integer"
    }
  ]
}
```

### Create Movie

Request :

- Method : POST
- Endpoint : `/api/v1/movies`
- Header :
  - Content-Type: application/json
  - Accept: application/json

Body :

```json
{
  "title": "string",
  "language": "string",
  "status": "string",
  "rating": "number"
}
```

Response :

```json
{
  "Success": "boolean",
  "Message": "string",
  "Data": {
    "id": "integer",
    "title": "string",
    "language": "string",
    "status": "string",
    "rating": "number"
  }
}
```

### Update Movie

Request :

- Method : PATCH
- Endpoint : `/api/v1/movies/{movieID}`
- Header :
  - Content-Type: application/json
  - Accept: application/json

Body :

```json
{
  "title": "string",
  "language": "string",
  "status": "integer",
  "rating": "integer"
}
```

Response :

```json
{
  "Success": "boolean",
  "Message": "string",
  "Data": {
    "id": "integer",
    "title": "string",
    "language": "string",
    "status": "string",
    "rating": "number"
  }
}
```

### Delete Cast

Request :

- Method : DELETE
- Endpoint : `/api/v1/movies/{movieID}`
- Header :
  - Accept: application/json

Response :

```json
{
  "messages": "string"
}
```
