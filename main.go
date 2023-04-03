package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger" // gin-swagger middleware
    "github.com/swaggo/files" // swagger embed files
    _ "github.com/sumaninster/crypto-server/docs"
    "github.com/sumaninster/crypto-server/src"
)

/*
Function to start the HTTP server and handle the endpoints.
In this function, we are using the gin library to define HTTP routes for each endpoint 
and start the HTTP server on port 8080.
*/
//	@title			Crypto Server API
//	@version		1.0
//	@description	This is a Crypto Server to fetch data from HitBTC API.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    r := gin.Default()
    v1 := r.Group("/api/v1")
    {
        currency := v1.Group("/currency")
        {
            h := src.HttpController{}
            currency.GET("/:symbol", h.HandleCurrencySymbol)
            currency.GET("/all", h.HandleAllCurrencySymbols)
        }
    }
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    err := r.Run(":8080")
    if err != nil {
        fmt.Println(err)
    }
}
