package src

// Cache duration in seconds
const cacheDuration = 3

// Redis connection details
const redisHost = "localhost:32768"
const redisPassword = "redispw"
const redisDB = 0

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