# CryptoTracker – Backend API

A RESTful backend service for tracking cryptocurrency prices using the CoinMarketCap Price Ticker API. Built with Go, Gin, and PostgreSQL, the application provides a structured backend architecture for retrieving and managing cryptocurrency market data, with scheduled background tasks for automated operations.

## 🚀 Features

- **Real-Time Cryptocurrency Data:** Fetch cryptocurrency price information through the CoinMarketCap API.
- **RESTful API:** Exposes HTTP endpoints using the Gin web framework.
- **PostgreSQL Integration:** Persistent data storage using PostgreSQL.
- **Scheduled Background Jobs:** Automated tasks powered by `go-cron`.
- **Modular Architecture:** Separation of controllers, models, repositories, routes, and utilities.
- **API Testing:** Endpoint testing and validation using Postman.

## 🛠️ Technology Stack

| Component | Technology |
|---|---|
| Programming Language | Go (Golang) |
| Web Framework | Gin |
| Database | PostgreSQL |
| External API | CoinMarketCap Price Ticker API |
| Background Scheduler | Go Cron |
| API Testing | Postman |

## 📁 Project Structure

```text
CryptoTracker/
│
├── routes/
│   └── routes.go
│
├── src/
│   ├── controllers/
│   ├── models/
│   └── repository/
│
├── utils/
│   ├── coinmarketcap/
│   └── database/
│
├── main.go
├── go.mod
├── go.sum
├── .env.example
└── README.md
```

## ⚙️ Prerequisites

Before running the application, ensure you have the following installed:

- [Go](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [CoinMarketCap API Key](https://coinmarketcap.com/api/)
- [Postman](https://www.postman.com/) (for API testing)

## 🔧 Installation & Setup

### 1. Clone the Repository

```bash
git clone <repository-url>
cd CryptoTracker
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Configure Environment Variables

Create a `.env` file in the project root directory and configure the required credentials.

```env
COINMARKETCAP_API_KEY=your_coinmarketcap_api_key

DB_HOST=localhost
DB_PORT=5432
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_NAME=your_database_name
DB_SSLMODE=disable
```

**Important:** Never commit your `.env` file or expose API keys and database credentials in version control.

### 4. Configure PostgreSQL

Create a PostgreSQL database for the application.

```sql
CREATE DATABASE cryptotracker;
```

Ensure that the database credentials configured in your `.env` file match your PostgreSQL configuration.

### 5. Run the Application

Start the backend server using:

```bash
go run main.go
```

The application will initialize its required services and start the API server.

## 🧪 API Testing

You can test the available API endpoints using Postman.

1. Start the backend application.
2. Open Postman.
3. Configure the appropriate HTTP method and API endpoint.
4. Send requests and inspect the responses.

Refer to the API routes defined in `routes/routes.go` for the available endpoints.

## 🔐 Security & Configuration

For secure application deployment, follow these practices:

- Store sensitive credentials in environment variables.
- Never expose CoinMarketCap API keys in client-side applications.
- Keep `.env` files out of version control.
- Validate incoming API requests.
- Handle external API failures and database errors gracefully.
- Configure appropriate database connection settings for production environments.

## 🏗️ Architecture

The application follows a modular, layered architecture:

| Layer | Responsibility |
|---|---|
| Routes | Defines API endpoints and request routing. |
| Controllers | Handles HTTP requests and responses. |
| Models | Defines application data structures. |
| Repository | Manages database operations. |
| CoinMarketCap Utility | Handles communication with the external cryptocurrency API. |
| Database Utility | Manages PostgreSQL connectivity and initialization. |

## 📌 Future Improvements

Potential enhancements for the project include:

- JWT-based authentication and authorization.
- Docker and Docker Compose support.
- Swagger / OpenAPI API documentation.
- Structured logging and centralized error handling.
- Unit and integration testing.
- Database migrations.
- Rate limiting and request validation.
- CI/CD pipeline integration.

## 📄 License

This project is available for educational and development purposes. Add an appropriate open-source license if you intend to distribute it publicly.
