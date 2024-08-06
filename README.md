# Crypto & Currency Exchange CLI

This is a command-line interface (CLI) application written in Go that displays live prices for Bitcoin and Ethereum, as well as currency exchange rates for GBP/USD, GBP/KSH, and USD/KSH.

## Features

- Real-time cryptocurrency prices (Bitcoin and Ethereum)
- Live currency exchange rates (GBP/USD, GBP/KSH, USD/KSH)
- Automatic updates every second
- Asynchronous data fetching to maintain responsiveness

## Prerequisites

- Go 1.13 or higher

## Installation

1. Clone the repository:
git clone https://github.com/QoSz/crypto-price-tracker.git
cd crypto-price-tracker

2. Build the application:
go build

## Usage
Run the application using:
./crypto-cli
The application will display the cryptocurrency prices and exchange rates, updating the screen every second with the latest available data.

## Data Sources

- Cryptocurrency data: CoinGecko API
- Exchange rate data: ExchangeRate-API

## Note on API Usage

This application uses free public APIs. Please be mindful of their usage limits:

- CoinGecko API: 50 calls/minute
- ExchangeRate-API: 1500 requests per month

The application is configured to fetch:
- Cryptocurrency data every 60 seconds
- Exchange rate data every 60 seconds

## License

[MIT](https://choosealicense.com/licenses/mit/)