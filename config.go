package main

const cacheDuration = 30

var supportedCurrencySymbolsMap = map[string]bool{
    "BTCUSDT": true,
    "ETHBTC":  true,
}

func getSupportedCurrencySymbols() []string {
    var keys []string
    for k, v := range supportedCurrencySymbolsMap {
        if v == true {
            keys = append(keys, k)
        }
    }
    return keys
}