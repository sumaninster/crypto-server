package src

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
    "fmt"
)

type HttpController struct{}

/*
Function to handle the /currency/{symbol} endpoint.
---------------------------------------------------
This function handles the endpoint for fetching the real-time crypto price of a given currency symbol. 
It first retrieves the symbol parameter from the request context and checks if it is empty. 
It then establishes a connection to Redis and attempts to retrieve the cached currency price data for the given symbol. 
If there is a cache hit, it uses the cached currency price data to return the response. 
Otherwise, it calls the fetchCurrencyPrice() function to retrieve the currency price data from the API, caches the data using the cacheCurrencyPrice() function, and returns the response. 
If there is an error at any point during the process, it returns an error response with an appropriate message.
*/
// handleCurrencySymbol godoc
//	@Summary		Get the real-time crypto price for a currency symbol
//	@Description	Get the real-time crypto price for a currency symbol
//	@Tags			Currency
//	@Accept			json
//	@Produce		json
//	@Param			symbol	path		string			true	"Currency symbol (e.g. BTCUSDT, ETHBTC)"
//	@Success		200		{object}	CurrencyPrice	"CurrencyPrice object"
//	@Failure		400		{object}	HTTPBadRequestError
//	@Failure		404		{object}	HTTPFileNotFoundError
//	@Failure		500		{object}	HTTPInternalServerError
//	@Router			/currency/{symbol} [get]
func (h *HttpController) HandleCurrencySymbol(c *gin.Context) {
    symbol := c.Param("symbol")
    if symbol == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "symbol is required"})
        return
    }
    // Check if currency price is cached in Redis
    cachedCurrencyPrice, err := getCachedCurrencyPrice(symbol)
    if err != nil {
        fmt.Println(err)
    }
    var currencyPrice CurrencyPrice
    if cachedCurrencyPrice.Timestamp != "" {
        currencyPrice = cachedCurrencyPrice
    } else {
        // Fetch currency price from API and cache it
        currencyPrice, err = fetchCurrencyPrice(symbol)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        err = cacheCurrencyPrice(symbol, currencyPrice, cacheDuration*time.Second)
        if err != nil {
            fmt.Println(err)
        }
    }
    c.JSON(http.StatusOK, currencyPrice)
}

/*
Function to handle the /currency/all endpoint.
----------------------------------------------
The function first checks if the currency prices are cached in Redis using the getSupportedCurrencySymbols() function which returns an array of supported currency symbols. 
If all the currency prices are cached, it returns them as a map of currency symbol and their respective currency price information.
If not all currency prices are cached, it fetches the currency prices from an external API using the fetchAllCurrencyPrices() function, which takes an array of currency symbols as an argument. 
The function then caches the fetched currency prices in Redis using the cacheCurrencyPrice() function.
Finally, the function returns the currency prices as a map of currency symbol and their respective currency price information.
*/
// handleAllCurrencySymbols godoc
//	@Summary		Get the real-time crypto prices for all supported currencies
//	@Description	Get the real-time crypto prices for all supported currencies
//	@Tags			Currency
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]CurrencyPrice	"Map of currency symbols and their real-time prices"
//	@Failure		400	{object}	HTTPBadRequestError
//	@Failure		404	{object}	HTTPFileNotFoundError
//	@Failure		500	{object}	HTTPInternalServerError
//	@Router			/currency/all [get]
func (h *HttpController) HandleAllCurrencySymbols(c *gin.Context) {
    symbols := getSupportedCurrencySymbols()
    cachedCurrencyPrices := make(map[string]CurrencyPrice)
    for _, key := range symbols {
        // Check if currency prices are cached in Redis
        cachedCurrencyPrice, err := getCachedCurrencyPrice(key)
        if err != nil {
            fmt.Println(err)
        }
        if cachedCurrencyPrice.Timestamp != "" {
            cachedCurrencyPrices[key] = cachedCurrencyPrice
        }
    }

    // If all currency prices are cached, return them
    if len(cachedCurrencyPrices) == len(symbols) {
        c.JSON(http.StatusOK, cachedCurrencyPrices)
        return
    }

    // Fetch currency prices from API and cache them
    currencyPricesMap, err := fetchAllCurrencyPrices(symbols)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    for symbol, currencyPrice := range currencyPricesMap {
        cacheCurrencyPrice(symbol, currencyPrice, cacheDuration*time.Second)
    }
    c.JSON(http.StatusOK, currencyPricesMap)
}
