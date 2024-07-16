# CryptoTracker Backend API #
# ------------------------- #

# Step: 1

# Add your key from Coin Market Cap Price Ticker api to the apiKey variable in the main.go file

# Step: 2

# Add your Postgres SQL password and database name to the dsn string in database.go file

# Step: 3

# Run the Project: go run main.go

# --------------------------------- #
# Database Used: Postgres SQL
# Other Technologies: gin, go cron, Price Ticker API of coin Market Cap
# --------------------------------- #

# Structure
        ├──routes
        |  ├──routes.go
        ├──src/
        |  ├──controllers
        |  ├──models
        |  ├──repository
        ├──utils/
        |  ├──coinmarketcap
        |  ├──database
        └──main.go