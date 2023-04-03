package src

// Struct to represent the currency price data returned by the HitBTC API:
type CurrencyPrice struct {
    Ask             string          `json:"ask"`
    Bid             string          `json:"bid"`
    Last            string          `json:"last"`
    Low             string          `json:"low"`
    High            string          `json:"high"`
    Open            string          `json:"open"`
    Volume          string          `json:"volume"`
    VolumeQuote     string          `json:"volume_quote"`
    Timestamp       string          `json:"timestamp"`
    CurrencyPair    CurrencyPair    `json:"currency_pair"`
}

// Struct to represent the currency pair data returned by the HitBTC API:
type CurrencyPair struct {
    Type                string `json:"type"`
    BaseCurrency        string `json:"base_currency"`
    QuoteCurrency       string `json:"quote_currency"`
    Status              string `json:"status"`
    QuantityIncrement   string `json:"quantity_increment"`
    TickSize            string `json:"tick_size"`
    TakeRate            string `json:"take_rate"`
    MakeRate            string `json:"make_rate"`
    FeeCurrency         string `json:"fee_currency"`
    MarginTrading       bool   `json:"margin_trading"`
    MaxInitialLeverage  string `json:"max_initial_leverage"`
}