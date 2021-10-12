# Golang Simple blockchain


Simple implementation of the blockchain in Go. Inspiration taken from article https://mycoralhealth.medium.com/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc.
External dependencies used: mux, godotenv and spew 



## Installation and running

* Clone git repo
* Navigate to the project folder
* Resolve all external dependencies with

```
go mod tidy
```
* Run the application with

```
go run .
```

## Usage

Application exposes 2 endpoints. For better experience please use API testing tools e.g Postman

* LIST whole blockchain with GET method on `http://localhost:8282`

* ADD a new block with POST method hitting `http://localhost:8282` and passing a json in the body of the request
  example of the request body:
    ```json
    {
    "Operation": "Sell",
    "Price": 5.5
    }
    ```
