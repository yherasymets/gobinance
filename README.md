**Assessment Task Golang**

Next, you need to connect the app to the Binance API using the library found at https://github.com/aiviaio/go-binance.

For the purpose of connecting to the public endpoints, you do not require any API keys.

Retrieve the first five trading pairs (symbols) through the endpoint api/v3/exchangeInfo. You should construct a slice of five symbols such as BTCUSDT, ETHUSDT, and so on.

To obtain the latest price of each symbol, use the endpoint /fapi/v1/ticker/price. Create goroutines and pass each symbol as a parameter to a separate goroutine. Each goroutine should retrieve the last price of its corresponding symbol from the endpoint and transmit the symbol and price as a map to the channel.

The channel will receive data from the goroutines and display it on the screen. The final output should appear as follows:

BTCUSDT 22000
ETHUSDT 1500
...

The goroutines should be terminated appropriately.
