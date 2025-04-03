# PowerBase

PowerBase is a high-performance API that provides superhero data with efficient caching and database management. It utilizes Go, Redis, and MySQL to deliver fast and reliable responses.

## Features
- 🚀 **Superhero Data API**: Retrieve details about various superheroes.
- ⚡ **High Performance**: Uses Redis caching to speed up responses.
- 🗄 **Database-Driven**: Stores and fetches data using MySQL.
- 📡 **Pagination & Filtering**: Supports limit-based pagination for optimized queries.
- 🔧 **Easy Configuration**: Load environment variables from `.env`.

## Installation

### Prerequisites
Ensure you have the following installed:
- Go 1.20+
- MySQL or MariaDB
- Redis

### Setup
1. Clone the repository:
   ```sh
   git clone https://github.com/AdityaP183/power-base.git
   cd power-base
   ```
2. Copy the example environment file:
   ```sh
   cp .env.example .env
   ```
   Edit `.env` to configure your database and Redis settings.
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Run the application:
   ```sh
   go run main.go
   ```

## API Endpoints

## Troubleshooting
- **Database Not Connecting?** Ensure MySQL is running and `.env` is correctly set.
- **Redis Connection Refused?** Start Redis using `redis-server` or `docker run -d -p 6379:6379 redis`.
- **No Data in Response?** Run database migrations and insert test data.

## Contributing
Pull requests are welcome! Please follow the coding standards and ensure tests pass before submitting.

## License
This project is licensed under the MIT License.

