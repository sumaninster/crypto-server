package src

import (
    "github.com/go-resty/resty/v2"
    "encoding/json"
    "strings"
    "errors"
)

/*
Function to fetch the currency price data from the HitBTC API.
--------------------------------------------------------------
The function fetchCurrencyPrice is a function that takes a symbol string parameter and returns a CurrencyPrice struct and an error.
It first checks whether the symbol is supported by checking if it is present in the supportedCurrencySymbolsMap. 
If the symbol is not supported, it returns an error indicating that the currency is not supported.
It then constructs the URL using the symbol parameter and fetches the currency price data using the HitBTC API. 
It unmarshals the response body into a CurrencyPrice struct.
The function also calls the fetchCurrencyPair function to get the CurrencyPair data for the symbol. 
If the fetchCurrencyPair function call returns no error, it sets the CurrencyPair field of the CurrencyPrice struct.
Finally, the function returns the CurrencyPrice struct and any errors that occurred.
*/
func fetchCurrencyPrice(symbol string) (CurrencyPrice, error) {
    if supportedCurrencySymbolsMap[symbol] != true {
        return CurrencyPrice{}, errors.New("Currency not supported")
    }
    url := "https://api.hitbtc.com/api/3/public/ticker/" + symbol
    client := resty.New()
    resp, err := client.R().
        SetHeader("Content-Type", "application/json").
        Get(url)
    if err != nil {
        return CurrencyPrice{}, err
    }
    var currencyPrice CurrencyPrice
    err = json.Unmarshal(resp.Body(), &currencyPrice)
    if err != nil {
        return CurrencyPrice{}, err
    }
    currencyPair, err := fetchCurrencyPair(symbol)
    if err == nil {
        currencyPrice.CurrencyPair = currencyPair
    }
    return currencyPrice, nil
}

/*
Function to fetch the currency price data for all supported symbols from the HitBTC API.
----------------------------------------------------------------------------------------
The function fetchAllCurrencyPrices takes a slice of strings symbols as input, which represents the currency symbols for which the prices need to be fetched. 
It constructs a URL using the given symbols and makes an HTTP GET request to the HitBTC API to fetch the currency prices.
The function then unmarshals the JSON response body into a map of string keys and CurrencyPrice values. 
If an error occurs during unmarshaling or while making the API request, it returns the error.
It then calls the fetchAllCurrencyPair function to get the currency pair details for each symbol in the symbols slice. 
If an error occurs while fetching currency pair details, it ignores the error and continues to the next symbol.
Finally, the function updates the CurrencyPair field of each CurrencyPrice struct with the corresponding currency pair details fetched earlier. 
It returns the updated currencyPricesMap with currency pairs and any error that occurred during execution.
*/
func fetchAllCurrencyPrices(symbols []string) (map[string]CurrencyPrice, error) {
    url := "https://api.hitbtc.com/api/3/public/ticker?symbols=" + strings.Join(symbols, ",")
    client := resty.New()
    resp, err := client.R().
        SetHeader("Content-Type", "application/json").
        Get(url)
    if err != nil {
        return nil, err
    }
    var currencyPricesMap map[string]CurrencyPrice
    err = json.Unmarshal(resp.Body(), &currencyPricesMap)
    if err != nil {
        return nil, err
    }
    currencyPairsMap, err := fetchAllCurrencyPair(symbols)
    if err == nil {
        for k, v := range currencyPricesMap {
            v.CurrencyPair = currencyPairsMap[k]
            currencyPricesMap[k] = v
        }
    }
    return currencyPricesMap, nil
}

/*
Function to fetch the currency pair data from the HitBTC API.
-------------------------------------------------------------
The function fetchCurrencyPair makes an API call to get the details of a currency pair using the symbol as the parameter. 
It first checks whether the given symbol is supported by checking whether the symbol exists in a map of supported currency symbols.
If the symbol is supported, it constructs the API endpoint URL with the symbol parameter and uses the resty library to make an HTTP GET request to the URL. 
It expects a JSON response containing details of the currency pair and tries to unmarshal the response into a CurrencyPair struct.
If there are no errors, it returns the CurrencyPair struct containing the details of the currency pair. 
If the symbol is not supported or there is any other error, it returns an error object with a corresponding error message.
*/
func fetchCurrencyPair(symbol string) (CurrencyPair, error) {
    if supportedCurrencySymbolsMap[symbol] != true {
        return CurrencyPair{}, errors.New("Currency not supported")
    }
    url := "https://api.hitbtc.com/api/3/public/symbol/" + symbol
    client := resty.New()
    resp, err := client.R().
        SetHeader("Content-Type", "application/json").
        Get(url)
    if err != nil {
        return CurrencyPair{}, err
    }
    var currencyPair CurrencyPair
    err = json.Unmarshal(resp.Body(), &currencyPair)
    if err != nil {
        return CurrencyPair{}, err
    }
    return currencyPair, nil
}

/*
Function to fetch the currency pair data for all supported symbols from the HitBTC API.
---------------------------------------------------------------------------------------
The fetchAllCurrencyPair function sends a GET request to the HitBTC API to fetch details about the currency pairs associated with the provided symbols. 
It takes a slice of strings (symbols) as input and returns a map of CurrencyPair objects with the symbol as the key and the associated CurrencyPair object as the value. 
It uses the resty package to send the HTTP request and the json package to unmarshal the response body into the map of CurrencyPair objects. 
If an error occurs during the request or unmarshalling process, the function returns nil and the error.
*/
func fetchAllCurrencyPair(symbols []string) (map[string]CurrencyPair, error) {
    url := "https://api.hitbtc.com/api/3/public/symbol?symbols=" + strings.Join(symbols, ",")
    client := resty.New()
    resp, err := client.R().
        SetHeader("Content-Type", "application/json").
        Get(url)
    if err != nil {
        return nil, err
    }
    var currencyPairsMap map[string]CurrencyPair
    err = json.Unmarshal(resp.Body(), &currencyPairsMap)
    if err != nil {
        return nil, err
    }
    return currencyPairsMap, nil
}
