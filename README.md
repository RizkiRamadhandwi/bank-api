# Bank API

## Overview

This document serves as a guide for Backend Developers working on the development of an API to facilitate interactions between merchants and banks. The API's primary functionalities include user authentication, payment processing, and activity logging. The development will be done using the Go programming language, employing the Gin-Gonic framework for building the API endpoints, JWT for security implementation, and JSON files for simulating customer, merchant, and transaction history data.

## List of Contents

- [Overview](#Overview)
- [Technologies-Used](#Technologies-Used)
- [Prerequisites](#Prerequisites)
- [Repository-Structure](#Repository-Structure)
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

## Repository-Structure

    bank-api/
    │
    ├── config/                 
    │   ├── app_config.go         
    │   └── config.go
    │
    ├── delivery/ 
    │   ├── controller/
    │   │   ├── auth_controller_test.go          
    │   │   ├── auth_controller.go  
    │   │   ├── transaction_controller_test.go                       
    │   │   └── transaction_controller.go 
    │   │        
    │   ├── middleware/         
    │   │   └── auth_middleware.go 
    │   │        
    │   └── server.go  
    │               
    ├── entity/ 
    │   ├── dto/
    │   │   ├── auth_dto.go          
    │   │   ├── transaction_dto.go                        
    │   │   └── user_dto.go
    │   │ 
    │   ├── transaction.go    
    │   └── user.go
    │                       
    ├── logging/               
    │   └── logging.go   
    │             
    ├── mock/
    │   ├── data_mock/   
    │   │   ├── customers.json          
    │   │   ├── merchants.json                        
    │   │   └── transactions.json
    │   │ 
    │   ├── middleware_mock/                          
    │   │   └── auth_middleware_mock.go
    │   │ 
    │   ├── repository_mock/   
    │   │   ├── merchant_repository_mock.go         
    │   │   ├── transaction_repository_mock.go                        
    │   │   └── user_repository_mock.go
    │   │ 
    │   ├── data_mock/                           
    │   │   └── jwt_service_mock.go
    │   │ 
    │   │  
    │   └── usecase_mock/   
    │       ├── merchant_usecase_mock.go         
    │       ├── transaction_usecase_mock.go                        
    │       └── user_usecase_mock.go
    │                 
    ├── repository/
    │   ├── data/
    │   │   ├── customers.json          
    │   │   ├── merchants.json                        
    │   │   └── transactions.json
    │   │ 
    │   ├── merchant_repository_test.go 
    │   ├── merchant_repository_test.go 
    │   ├── transaction_repository_test.go 
    │   ├── transaction_repository_test.go 
    │   ├── user_repository_test.go    
    │   └── user_repository.go
    │                  
    ├── shared/ 
    │   ├── common/                           
    │   │   └── json_response.go
    │   │ 
    │   ├── model/
    │   │   ├── json_model.go          
    │   │   ├── my_custom_claim.go                        
    │   │   └── pagination_model.go
    │   │  
    │   └── service/                         
    │       └── jwt_service.go
    │               
    ├── usecase/ 
    │   ├── merchant_usecase_test.go 
    │   ├── merchant_usecase_test.go 
    │   ├── transaction_usecase_test.go 
    │   ├── transaction_usecase_test.go 
    │   ├── user_usecase_test.go    
    │   └── user_usecase.go
    │                 
    ├── .env                   
    ├── coverage.out
    ├── docker-compose.yml                   
    ├── Dockerfile                   
    ├── go.mod                   
    ├── go.sum                   
    ├── main.go                   
    ├── README.md                   
    └── user_activity.log


## Instalation

This application does not require any special installation. Make sure you have installed Go (programming language) on your computer. To run the application:

1. Open a terminal or command prompt on your computer.
2. Clone the repository from GitHub using the following command: git clone https://github.com/RizkiRamadhandwi/bank-api.git
3. Navigate to the directory where you want to save the project code by typing cd bank-api.
4. Run go mod tidy to install dependencies.
5. Run the application in the terminal with the command go run ..

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