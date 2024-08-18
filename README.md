# Strong Password Checker API

This project implements a REST API for checking the strength of passwords and suggesting steps to make them stronger. It's built with Go, uses Gin as the web framework, PostgreSQL for data storage, and Docker for containerization.

## Prerequisites

Before you begin, ensure you have the following installed on your local machine:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/doc/install) (version 1.16 or later)

## Deploying Locally

To deploy the application on your local machine, follow these steps:

1. Clone the repository:

   ```
   git clone https://github.com/nanpipat/nanpipat_agnos_backend.git
   cd nanpipat_agnos_backend
   ```

2. Build and start the Docker containers:

   ```
   docker-compose up --build
   ```

   This command will build the Docker images and start the containers for the Go application, PostgreSQL database, and Nginx reverse proxy.

3. The application should now be running and accessible at `http://localhost:80`.

4. To stop the application, use:
   ```
   docker-compose down
   ```

## API Usage

Once the application is running, you can use the API to check password strength:

- **Endpoint**: `POST /api/strong_password_steps`
- **Content-Type**: `application/json`
- **Request Body**:
  ```json
  {
    "init_password": "YourPasswordHere"
  }
  ```
- **Response**:
  ```json
  {
    "num_of_steps": 2
  }
  ```

Example using curl:

```
curl -X POST http://localhost/api/strong_password_steps \
     -H "Content-Type: application/json" \
     -d '{"init_password": "aA1"}'
```

## Running Unit Tests

To run the unit tests for this project, follow these steps:

1. Ensure you're in the project root directory.

2. Run the tests using the Go test command:

   ```
   go test ./... -v
   ```

   This command will run all tests in the project and its subdirectories with verbose output.

3. To run tests for a specific package (e.g., the password service), use:

   ```
   go test ./services -v
   ```

4. To generate a coverage report, use:

   ```
   go test ./... -coverprofile=coverage.out
   ```

   You can then view the coverage report in your browser by running:

   ```
   go tool cover -html=coverage.out
   ```

## Troubleshooting

If you encounter any issues:

1. Ensure all required ports (80, 5432) are free on your machine.
2. Check Docker logs for any error messages:
   ```
   docker-compose logs
   ```
3. Ensure your Go version is up to date.

For more detailed information or if you encounter persistent issues, please open an issue in the GitHub repository.
