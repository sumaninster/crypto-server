# Crypto Server

## Libraries Used

The gin library is used for handling HTTP requests and responses.
The redis library is used for caching currency price data in-memory. 
The resty library is used for making HTTP requests to the HitBTC API.

The use of external libraries such as gin, redis, and resty id needed to handle HTTP requests and responses, cache data in-memory for improved performance, and make HTTP requests to external APIs. These libraries are widely used and well-tested in production environments, and using them saves time and effort compared to building these functionalities from scratch.

## Run the service
go run *.go

## Configure supported currencies
To manage supported currencies, update the variable "supportedCurrencySymbolsMap" in the file "config.go"

## Swagger API Docs

http://localhost:8080/swagger/index.html


## GET End Points

### GET /currency/{symbol}

http://localhost:8080/api/v1/currency/BTCUSDT

http://localhost:8080/api/v1/currency/ETHBTC

### GET /currency/all

http://localhost:8080/api/v1/currency/all