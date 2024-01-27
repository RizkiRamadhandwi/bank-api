# Bank Aplication

## Overview

This document serves as a guide for Backend Developers working on the development of an API to facilitate interactions between merchants and banks. The API's primary functionalities include user authentication, payment processing, and activity logging. The development will be done using the Go programming language, employing the Gin-Gonic framework for building the API endpoints, JWT for security implementation, and JSON files for simulating customer, merchant, and transaction history data.

## List of Contents

- [Overview](#Overview)
- [Technologies-Used](#Technologies-Used)
- [Prerequisites](#Prerequisites)
- [Instalation](#Instalation)
- [Running-the-Application](#Running-the-Application)
- [API-Spec](#API-Spec)
  - [Login-API](#Login-API)
  - [Transaction-API](#Transaction-API)
    - [Create-Transaction](#Create-Transaction)
    - [List-Transaction](#List-Transaction)
  - [Logout-API](#Logout-API)
- [Security](#Security)

## Technologies-Used

- Programming Language: Golang
- Web Framework: Gin-Gonic
- Authentication: JWT (JSON Web Token)
- Data Storage: JSON files for simulating customer, merchant, and transaction history data

## Prerequisites

Before running the Bank application, make sure you have fulfilled the following prerequisites:

- Go (Golang) is installed on your system.
- An active internet connection is required to download Go dependencies.

## Instalation

This application does not require any special installation. Make sure you have installed Go (programming language) on your computer. To run the application:

1. Open a terminal or command prompt on your computer.
2. Clone repository from GitHub using the following command: `git clone https://github.com/RizkiRamadhandwi/bank-api.git`
3. Navigate to the directory where you want to save the project code by entering `cd bank-api`.
4. Run the application in the terminal with the command `go run .`.

## Running-the-Application

Once the application is running, you can access it through a web browser or use it through an API client such as Postman or cURL. This application provides APIs for Transactions.



## API-Spec

### Login-API

Request :

- Method : `POST`
- Endpoint : `/api/v1/auth/login`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
  "username": "string",
  "password": "string"
}
```

### Transaction-API

##### Create-Transaction

Request :

- Method : POST
- Endpoint : `/api/v1/transaction`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token
- Body :

```json
{
    "merchant_id": "string",
    "amount": int
}
```

Response :

- Status : 201 Created
- Body :

```json
{
  "status": {
    "code": 201,
    "message": "Created"
  },
  "data": {
    "id": int,
    "customer": {
        "id": "string",
        "name": "string"
    },
    "merchant": {
        "id": "string",
        "name": "string"
    },
    "amount": int,
    "createdAt": "2000-01-01T12:00:00Z", (curent time)
  }
}
```

#### List-Transaction 

Request :

- Method : GET
- Endpoint : `/api/v1/transaction`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Query Param :
  - page : int `optional`
  - size : int `optional`
- Authorization : Bearer Token

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 200,
        "message": "Ok"
    },
    "data": [
        {
          "id": int,
          "customer": {
              "id": "string",
              "name": "string"
          },
          "merchant": {
              "id": "string",
              "name": "string"
          },
          "amount": int,
          "createdAt": "2000-01-01T12:00:00Z", (curent time)
        }
    ],
    "paging": {
        "page": 1,          (default value)
        "rowsPerPage": 10,   (default value)
        "totalRows": int,
        "totalPages": int
    }
}

```

### Logout-API

Request :

- Method : POST
- Endpoint : `/api/v1/auth/logout`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token


Response :

- Status : 204 No Content



## Security

Input Validation: All user inputs will be strictly validated to prevent injection attacks and other security vulnerabilities.