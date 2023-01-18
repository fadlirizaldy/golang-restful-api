# Golang Restful API

This is golang rest API to do CRUD data

Tools :

- Golang (Echo Framework)
- MySQL
- Open API Swagger

---

# Spec

To see the documentation with Swagger, please visit https://editor.swagger.io/ and put code in "spec.json" to it

---

## Get all casts

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

## Get detail cast

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

## Create Cast

Request :

- Method : POST
- Endpoint : `/api/v1/casts`
- Header :
  - Content-Type: application/json
  - Accept: application/json
- Body :

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

## Update Cast

Request :

- Method : PATCH
- Endpoint : `/api/v1/casts/{castID}`
- Header :
  - Content-Type: application/json
  - Accept: application/json
- Body :

```json
{
  "name": "string",
  "birth_place": "string",
  "birthday": "datetime",
  "rating": "integer"
}
```

## Delete Cast

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

## Get all movies

Request :

- Method : GET
- Endpoint : `/api/v1/movies`
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
      "title": "string"
    }
  ]
}
```

## Get detail movie

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

## Create Movie

Request :

- Method : POST
- Endpoint : `/api/v1/movies`
- Header :
  - Content-Type: application/json
  - Accept: application/json
- Body :

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

## Update Movie

Request :

- Method : PATCH
- Endpoint : `/api/v1/movies/{movieID}`
- Header :
  - Content-Type: application/json
  - Accept: application/json
- Body :

```json
{
  "title": "string",
  "language": "string",
  "status": "integer",
  "rating": "integer"
}
```

## Delete Cast

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
