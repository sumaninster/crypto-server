package main

import (
    "context"
    "encoding/json"
    "github.com/redis/go-redis/v9"
    "time"
)

/*
Function to cache the currency price data in-memory using the Redis library.
----------------------------------------------------------------------------
In this function, we are using the json package to serialize the CurrencyPrice struct into JSON format. 
We are then using the Set() method of the Redis client to store the serialized JSON data with a specified expiration time.
*/
func cacheCurrencyPrice(symbol string, currencyPrice CurrencyPrice, duration time.Duration, rdb *redis.Client) error {
    ctx := context.Background()
    currencyPriceJSON, err := json.Marshal(currencyPrice)
    if err != nil {
        return err
    }
    err = rdb.Set(ctx, symbol, currencyPriceJSON, duration).Err()
    if err != nil {
        return err
    }
    return nil
}

/*
Function to retrieve the cached currency price data from Redis.
---------------------------------------------------------------
In this function, we are using the Get() method of the Redis client to retrieve the cached currency price data for the specified symbol. 
We are then using the json package to deserialize the JSON data into a CurrencyPrice struct.
*/
func getCachedCurrencyPrice(symbol string, rdb *redis.Client) (CurrencyPrice, error) {
    ctx := context.Background()
    currencyPriceJSON, err := rdb.Get(ctx, symbol).Bytes()
    if err != nil {
        return CurrencyPrice{}, err
    }
    currencyPrice := CurrencyPrice{}
    err = json.Unmarshal(currencyPriceJSON, &currencyPrice)
    if err != nil {
        return CurrencyPrice{}, err
    }
    return currencyPrice, nil
}
